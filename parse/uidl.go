package parse

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/flowdev/fdialog/parse/uidl"
	"github.com/flowdev/fdialog/ui"
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

func ParseUIDL(input io.Reader, _ string) (map[string]map[string]any, error) {
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

	return convertUIDL(antlrParser.Uidl(), ael), ael.CombinedError()
}

func convertUIDL(antlrUIDL uidl.IUidlContext, errColl ErrorCollector) map[string]map[string]any {
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
func convertCommands(antlrCommands []uidl.ICommandContext, errColl ErrorCollector) map[string]map[string]any {
	commandMap := make(map[string]map[string]any, len(antlrCommands))

	for _, command := range antlrCommands {
		keyword := command.Identifier(0)
		name := command.Identifier(1)
		strName := name.GetText()
		if _, ok := commandMap[strName]; ok {
			errColl.CollectError(
				fmt.Errorf("%s duplicate command name: %q",
					errorContext(name.GetSymbol()), strName),
			)
			continue
		}
		attrMap := convertAttributes(command.Attributes().AllAttribute(), errColl)
		attrMap[ui.KeyKeyword] = keyword.GetText()
		commandMap[strName] = attrMap

		if command.CommandBody() != nil {
			attrMap[ui.KeyChildren] = convertCommands(command.CommandBody().Commands().AllCommand(), errColl)
		}
	}

	return commandMap
}

func convertAttributes(attributes []uidl.IAttributeContext, errColl ErrorCollector) map[string]any {
	attrMap := make(map[string]any, len(attributes)+2) // space for keyword + children

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

/*
func parseUIDL(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		parseVersion,
		parseCommands,
		gparselib.NewParseEOFPlugin(nil),
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		versionResult := pd.SubResults[0]
		version := versionResult.Value.(uint64)
		if version != 1 {
			pd.AddError(
				versionResult.Pos,
				fmt.Sprintf("expected version %d, got: %d", UIDLVersion, version),
				nil,
			)
		}
		pd.Result.Value = pd.SubResults[1].Value
		pd.CleanFeedback(true)
		return pd, ctx
	})
}

func parseCommand(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseIdentPlugin(textSemantic, "", ""),
		parseSpaceComment,
		gparselib.NewParseIdentPlugin(textSemantic, "_", "_"),
		parseSpaceComment,
		parseAttributes,
		gparselib.NewParseOptionalPlugin(parseCommandBody, nil),
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		keyword := pd.SubResults[0].Value.(string)
		command := commandValue{
			name:         pd.SubResults[2].Value.(string),
			attributeMap: pd.SubResults[4].Value.(map[string]any),
		}
		achildren := pd.SubResults[5].Value
		if achildren != nil {
			children := achildren.(map[string]map[string]any)
			command.attributeMap[KeyChildren] = children
		}
		command.attributeMap[KeyKeyword] = keyword

		pd.CleanFeedback(false)
		pd.Result.Value = command
		return pd, ctx
	})
}

func parseCommands(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseMulti0(pd, ctx, parseCommand,
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			// convert all commands to a map
			commandMap := make(map[string]map[string]any, len(pd.SubResults))

			for _, subResult := range pd.SubResults {
				command := subResult.Value.(commandValue)
				if _, ok := commandMap[command.name]; ok {
					pd.AddError(subResult.Pos,
						fmt.Sprintf("duplicate command name: %q", command.name),
						nil,
					)
					continue
				}
				commandMap[command.name] = command.attributeMap
			}

			pd.Result.Value = commandMap
			return pd, ctx
		})
}

func parseCommandBody(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseLiteralPlugin(nil, "{"),
		parseSpaceComment,
		parseCommands,
		gparselib.NewParseLiteralPlugin(nil, "}"),
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		pd.Result.Value = pd.SubResults[2].Value
		return pd, ctx
	})
}

func parseAttributes(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pParenOpen := gparselib.NewParseLiteralPlugin(nil, "(")
	pParenClose := gparselib.NewParseLiteralPlugin(nil, ")")

	pNoAttributes := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		pParenOpen,
		parseSpaceComment,
		pParenClose,
		parseSpaceComment,
	}, nil)

	pCommaAttribute := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		gparselib.NewParseLiteralPlugin(nil, ","),
		parseSpaceComment,
		parseAttribute,
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		pd.Result.Value = pd.SubResults[2].Value
		return pd, ctx
	})
	pManyAttributes := gparselib.NewParseMulti0Plugin(pCommaAttribute,
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			attrMap := make(map[string]any, len(pd.SubResults)+1) // there will be one more attribute later

			for _, subResult := range pd.SubResults { // copy over for the type
				attr := subResult.Value.(attributeValue)
				if _, ok := attrMap[attr.key]; ok {
					pd.AddError(subResult.Pos,
						fmt.Sprintf("duplicate attribute key: %q", attr.key),
						nil,
					)
					continue
				}
				attrMap[attr.key] = attr.value
			}

			pd.Result.Value = attrMap
			return pd, ctx
		})
	pAllAttributes := gparselib.NewParseAllPlugin([]gparselib.SubparserOp{
		pParenOpen,
		parseSpaceComment,
		parseAttribute,
		pManyAttributes,
		pParenClose,
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		firstAttr := pd.SubResults[2].Value.(attributeValue)
		attrMap := pd.SubResults[3].Value.(map[string]any)
		if _, ok := attrMap[firstAttr.key]; ok {
			pd.AddError(pd.SubResults[2].Pos,
				fmt.Sprintf("duplicate attribute key: %q", firstAttr.key),
				nil,
			)
		} else {
			attrMap[firstAttr.key] = firstAttr.value
		}

		pd.Result.Value = attrMap
		return pd, ctx
	})

	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{
		pNoAttributes,
		pAllAttributes,
	}, nil)
}

func parseAttribute(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{
		gparselib.NewParseIdentPlugin(textSemantic, "_", "_"),
		parseSpaceComment,
		gparselib.NewParseLiteralPlugin(nil, "="),
		parseSpaceComment,
		parseValue,
		parseSpaceComment,
	}, func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
		attr := attributeValue{
			key:   pd.SubResults[0].Value.(string),
			value: pd.SubResults[4].Value,
		}
		if checkAttributeKey(attr.key, pd) {
			pd.Result.Value = attr
		}
		return pd, ctx
	})
}

func checkAttributeKey(attributeKey string, pd *gparselib.ParseData) bool {
	reservedAttributeKeys := map[string]bool{
		KeyKeyword:  true,
		KeyChildren: true,
		KeyName:     true,
	}
	if reservedAttributeKeys[attributeKey] {
		pd.AddError(
			pd.Result.Pos,
			fmt.Sprintf("reserved key %q can't be used as an attribute name", attributeKey),
			nil,
		)
		return false
	}
	return true
}

func parseValue(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{
		parseNormalString, parseBacktickString, parseBoolValue, parseFloatValue, parseIntValue,
	}, nil)
}

func parseNormalString(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pDQuote := gparselib.NewParseLiteralPlugin(nil, `"`)

	afterBackslash := false
	pInnerString := gparselib.NewParseGoodRunesPlugin(textSemantic,
		func(r rune) bool {
			if afterBackslash {
				afterBackslash = false
				return true
			}
			if r == '\\' {
				afterBackslash = true
				return true
			}
			return r != '"'
		})

	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pDQuote, pInnerString, pDQuote},
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			var err error
			pd.Result.Value, err = strconv.Unquote(`"` + pd.SubResults[1].Value.(string) + `"`)
			if err != nil {
				pd.AddError(pd.Result.Pos, "invalid normal string value", err)
			}
			return pd, ctx
		})
}

func parseBacktickString(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pBacktick := gparselib.NewParseLiteralPlugin(nil, "`")
	pNoBacktick := gparselib.NewParseGoodRunesPlugin(textSemantic,
		func(r rune) bool {
			return r != '`'
		})

	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pBacktick, pNoBacktick, pBacktick},
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value = pd.SubResults[1].Value
			var err error
			pd.Result.Value, err = strconv.Unquote("`" + pd.SubResults[1].Value.(string) + "`")
			if err != nil {
				pd.AddError(pd.Result.Pos, "invalid backtick string value", err)
			}
			return pd, ctx
		})
}

func parseFloatValue(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pPlus := gparselib.NewParseLiteralPlugin(nil, "+")
	pMinus := gparselib.NewParseLiteralPlugin(nil, "-")
	pSign := gparselib.NewParseAnyPlugin([]gparselib.SubparserOp{pPlus, pMinus}, nil)
	pOptSign := gparselib.NewParseOptionalPlugin(pSign, nil)
	pNum, err := gparselib.NewParseNaturalPlugin(nil, 10)
	if err != nil {
		panic(err) // can only happen if cfgRadix is invalid
	}
	pDot := gparselib.NewParseLiteralPlugin(nil, ".")
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pOptSign, pNum, pDot, pNum},
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value, err = strconv.ParseFloat(pd.Result.Text, 64)
			if err != nil {
				pd.AddError(pd.Result.Pos, err.Error(), nil)
			} else {
				pd.CleanFeedback(true)
			}
			return pd, ctx
		})
}

func parseIntValue(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pPlus := gparselib.NewParseLiteralPlugin(nil, "+")
	pMinus := gparselib.NewParseLiteralPlugin(nil, "-")
	pSign := gparselib.NewParseAnyPlugin([]gparselib.SubparserOp{pPlus, pMinus}, nil)
	pOptSign := gparselib.NewParseOptionalPlugin(pSign, nil)
	pNum, err := gparselib.NewParseNaturalPlugin(nil, 10)
	if err != nil {
		panic(err) // can only happen if cfgRadix is invalid
	}
	return gparselib.ParseAll(pd, ctx, []gparselib.SubparserOp{pOptSign, pNum},
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value, err = strconv.ParseInt(pd.Result.Text, 10, 64) // TODO: handle overflow?
			if err != nil {
				pd.AddError(pd.Result.Pos, err.Error(), nil)
			} else {
				pd.CleanFeedback(true)
			}
			return pd, ctx
		})
}

func parseBoolValue(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pTrue := gparselib.NewParseLiteralPlugin(
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value = true
			return pd, ctx
		}, "true")
	pFalse := gparselib.NewParseLiteralPlugin(
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value = false
			return pd, ctx
		}, "false")

	return gparselib.ParseAny(pd, ctx, []gparselib.SubparserOp{pTrue, pFalse}, nil)
}

// parseSpaceComment parses any amount of space (including newline) and line
// (`#` ... <NL>) comments.
//   - Semantic result: The parsed text plus a signal whether a newline was
//     parsed (spaceCommentSemValue).
func parseSpaceComment(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pSpc := gparselib.NewParseSpacePlugin(textSemantic, true)
	pLnCmnt, err := gparselib.NewParseLineCommentPlugin(textSemantic, `#`)
	if err != nil {
		panic(err) // can only be a programming error!
	}
	pSpaceOrComment := gparselib.NewParseAnyPlugin(
		[]gparselib.SubparserOp{pSpc, pLnCmnt}, textSemantic,
	)
	return gparselib.ParseMulti0(pd, ctx, pSpaceOrComment,
		func(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
			pd.Result.Value = pd.Result.Text
			pd.CleanFeedback(true)
			return pd, ctx
		},
	)
}

//
// ----------------------------------------------------------------------------
// Semantics:
//

type commandValue struct {
	name         string
	attributeMap map[string]any
}

type attributeValue struct {
	key   string
	value any
}

// textSemantic stores the successfully parsed text as semantic value.
func textSemantic(pd *gparselib.ParseData, ctx any) (*gparselib.ParseData, any) {
	pd.Result.Value = pd.Result.Text
	return pd, ctx
}
*/
