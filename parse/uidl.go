package parse

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/flowdev/fdialog/parse/uidl"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/x/omap"
	"io"
	"strconv"
)

const UIDLVersion = 1

type ErrorCollector interface {
	CollectError(error)
}
type AntlrErrorListener struct {
	*antlr.DefaultErrorListener
	errs []error
}

func NewAntlrErrorListener() *AntlrErrorListener {
	return &AntlrErrorListener{errs: make([]error, 0, 32)}
}

// SyntaxError stores Go errors with messages of
// the following format:
//
//	line <line>:<column> <message>
func (ael *AntlrErrorListener) SyntaxError(
	_ antlr.Recognizer,
	_ interface{},
	line, column int, msg string,
	_ antlr.RecognitionException,
) {
	ael.CollectError(fmt.Errorf("line %d:%d %s", line, column, msg))
}

func (ael *AntlrErrorListener) CollectError(err error) {
	ael.errs = append(ael.errs, err)
}

func (ael *AntlrErrorListener) CombinedError() error {
	if len(ael.errs) == 0 {
		return nil
	}
	return errors.Join(ael.errs...)
}

func UIDL(input io.Reader, _ string) (ui.CommandsDescr, error) {
	inputStr, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	antlrInput := antlr.NewInputStream(string(inputStr))
	antlrLexer := uidl.NewUIDLLexer(antlrInput)
	antlrStream := antlr.NewCommonTokenStream(antlrLexer, 0)
	antlrParser := uidl.NewUIDLParser(antlrStream)
	ael := NewAntlrErrorListener()
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(ael)

	if err = ael.CombinedError(); err != nil {
		return nil, err
	}
	return convertUIDL(antlrParser.Uidl(), ael), nil
}

func convertUIDL(antlrUIDL uidl.IUidlContext, errColl ErrorCollector) ui.CommandsDescr {
	version := antlrUIDL.Version().Natural()
	errCtx := errorContext(version.GetSymbol())

	intVersion, err := strconv.Atoi(version.GetText())
	if err != nil {
		errColl.CollectError(fmt.Errorf("%s %w", errCtx, err))
		return nil
	}
	if intVersion != UIDLVersion {
		errColl.CollectError(fmt.Errorf("%s expected version %d, got: %d",
			errCtx, UIDLVersion, intVersion))
		return nil
	}
	return convertCommands(antlrUIDL.Commands().AllCommand(), errColl)
}

// convertCommands converts all commands to a map.
func convertCommands(antlrCommands []uidl.ICommandContext, errColl ErrorCollector) ui.CommandsDescr {
	commandMap := omap.New[string, ui.AttributesDescr](len(antlrCommands))

	for _, command := range antlrCommands {
		keyword := command.Identifier(0)
		name := command.Identifier(1)
		if keyword == nil || name == nil { // not really a command (shouldn't be possible, but ...)
			continue
		}
		strName := name.GetText()
		attrMap := convertAttributes(command.Attributes().AllAttribute(), errColl)
		attrMap[ui.AttrKeyword] = keyword.GetText()
		if ok := commandMap.Add(strName, attrMap); !ok {
			errColl.CollectError(
				fmt.Errorf("%s duplicate command name: %q",
					errorContext(name.GetSymbol()), strName),
			)
			continue
		}

		if body := command.CommandBody(); body != nil && body.Commands() != nil {
			attrMap[ui.AttrChildren] = convertCommands(body.Commands().AllCommand(), errColl)
		}
	}

	return commandMap
}

func convertAttributes(attributes []uidl.IAttributeContext, errColl ErrorCollector) ui.AttributesDescr {
	attrMap := make(ui.AttributesDescr, len(attributes)+2) // space for keyword + children

	for _, attribute := range attributes {
		name := attribute.Identifier()
		strName := name.GetText()
		if _, ok := attrMap[strName]; ok {
			errColl.CollectError(
				fmt.Errorf("%s duplicate attribute key: %q",
					errorContext(name.GetSymbol()), strName),
			)
			continue
		}
		attrMap[strName] = convertAttributeValue(attribute.Value(), errColl)
	}

	return attrMap
}

func convertAttributeValue(antlrValue uidl.IValueContext, errColl ErrorCollector) any {
	simpleValue := antlrValue.SimpleValue()
	if simpleValue != nil {
		return convertSimpleValue(simpleValue, errColl)
	}
	listValue := antlrValue.ListValue()
	if listValue != nil {
		return convertListValue(listValue, errColl)
	}
	return nil
}

func convertListValue(antlrList uidl.IListValueContext, errColl ErrorCollector) any {
	antlrValues := antlrList.AllSimpleValue()
	result := make([]any, len(antlrValues))
	for i := 0; i < len(antlrValues); i++ {
		result[i] = convertSimpleValue(antlrValues[i], errColl)
	}
	return result
}

func convertSimpleValue(antlrValue uidl.ISimpleValueContext, errColl ErrorCollector) any {
	doubleQuotedString := antlrValue.DoubleQuotedString()
	backQuotedString := antlrValue.BackQuotedString()
	aFloat := antlrValue.Float()
	natural := antlrValue.Natural()
	aInt := antlrValue.Int()
	aBool := antlrValue.Bool()

	switch {
	case doubleQuotedString != nil:
		s, err := strconv.Unquote(doubleQuotedString.GetText())
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.DoubleQuotedString().GetSymbol()), err))
		}
		return s
	case backQuotedString != nil:
		s, err := strconv.Unquote(backQuotedString.GetText())
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.BackQuotedString().GetSymbol()), err))
		}
		return s
	case aFloat != nil:
		f, err := strconv.ParseFloat(aFloat.GetText(), 64)
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.Float().GetSymbol()), err))
		}
		return f
	case natural != nil:
		n, err := strconv.ParseInt(natural.GetText(), 0, 64)
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.Natural().GetSymbol()), err))
		}
		return n
	case aInt != nil:
		i, err := strconv.ParseInt(aInt.GetText(), 0, 64)
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.Int().GetSymbol()), err))
		}
		return i
	case aBool != nil:
		b, err := strconv.ParseBool(aBool.GetText())
		if err != nil {
			errColl.CollectError(fmt.Errorf("%s %w", errorContext(antlrValue.Bool().GetSymbol()), err))
		}
		return b
	}
	errColl.CollectError(fmt.Errorf("%s unknown value %q", errorContext(antlrValue.GetStart()), antlrValue.GetText()))
	return nil
}

func errorContext(symbol antlr.Token) string {
	return fmt.Sprintf("line %d:%d", symbol.GetLine(), symbol.GetColumn())
}
