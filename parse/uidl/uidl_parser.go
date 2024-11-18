// Code generated from UIDL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package uidl // UIDL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type UIDLParser struct {
	*antlr.BaseParser
}

var UIDLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func uidlParserInit() {
	staticData := &UIDLParserStaticData
	staticData.LiteralNames = []string{
		"", "'uidl'", "'{'", "'}'", "'('", "')'", "'='", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "Bool", "DoubleQuotedString", "BackQuotedString",
		"Identifier", "Natural", "Float", "Int", "Semicolon", "Comma", "WhiteSpace",
	}
	staticData.RuleNames = []string{
		"uidl", "version", "commands", "command", "commandSeparator", "commandBody",
		"attributes", "attribute", "value", "listValue", "simpleValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 136, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 3, 0, 24, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 30, 8, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 42, 8, 2, 10,
		2, 12, 2, 45, 9, 2, 1, 2, 3, 2, 48, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		54, 8, 3, 1, 3, 1, 3, 3, 3, 58, 8, 3, 1, 3, 3, 3, 61, 8, 3, 1, 4, 1, 4,
		1, 4, 3, 4, 66, 8, 4, 1, 5, 1, 5, 3, 5, 70, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 3, 6, 77, 8, 6, 1, 6, 1, 6, 1, 6, 5, 6, 82, 8, 6, 10, 6, 12, 6,
		85, 9, 6, 1, 6, 3, 6, 88, 8, 6, 3, 6, 90, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		3, 7, 96, 8, 7, 1, 7, 1, 7, 3, 7, 100, 8, 7, 1, 7, 1, 7, 3, 7, 104, 8,
		7, 1, 8, 1, 8, 3, 8, 108, 8, 8, 1, 9, 1, 9, 3, 9, 112, 8, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 117, 8, 9, 10, 9, 12, 9, 120, 9, 9, 1, 9, 3, 9, 123, 8, 9,
		3, 9, 125, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		134, 8, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0,
		1, 2, 0, 13, 13, 15, 15, 149, 0, 23, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4,
		37, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 67, 1, 0, 0,
		0, 12, 74, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 107, 1, 0, 0, 0, 18, 109,
		1, 0, 0, 0, 20, 133, 1, 0, 0, 0, 22, 24, 5, 18, 0, 0, 23, 22, 1, 0, 0,
		0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27,
		3, 8, 4, 0, 27, 29, 3, 4, 2, 0, 28, 30, 5, 18, 0, 0, 29, 28, 1, 0, 0, 0,
		29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 5, 0, 0, 1, 32, 1, 1, 0,
		0, 0, 33, 34, 5, 1, 0, 0, 34, 35, 5, 18, 0, 0, 35, 36, 5, 13, 0, 0, 36,
		3, 1, 0, 0, 0, 37, 43, 3, 6, 3, 0, 38, 39, 3, 8, 4, 0, 39, 40, 3, 6, 3,
		0, 40, 42, 1, 0, 0, 0, 41, 38, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		46, 48, 3, 8, 4, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 5, 1, 0,
		0, 0, 49, 50, 5, 12, 0, 0, 50, 51, 5, 18, 0, 0, 51, 53, 5, 12, 0, 0, 52,
		54, 5, 18, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0,
		0, 0, 55, 57, 3, 12, 6, 0, 56, 58, 5, 18, 0, 0, 57, 56, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 61, 3, 10, 5, 0, 60, 59, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 7, 1, 0, 0, 0, 62, 66, 5, 16, 0, 0, 63, 64,
		5, 18, 0, 0, 64, 66, 4, 4, 0, 1, 65, 62, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0,
		66, 9, 1, 0, 0, 0, 67, 69, 5, 2, 0, 0, 68, 70, 5, 18, 0, 0, 69, 68, 1,
		0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 3, 4, 2, 0, 72,
		73, 5, 3, 0, 0, 73, 11, 1, 0, 0, 0, 74, 76, 5, 4, 0, 0, 75, 77, 5, 18,
		0, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 89, 1, 0, 0, 0, 78, 83,
		3, 14, 7, 0, 79, 80, 5, 17, 0, 0, 80, 82, 3, 14, 7, 0, 81, 79, 1, 0, 0,
		0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 87,
		1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 88, 5, 17, 0, 0, 87, 86, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 78, 1, 0, 0, 0, 89, 90, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 5, 5, 0, 0, 92, 13, 1, 0, 0, 0, 93,
		95, 5, 12, 0, 0, 94, 96, 5, 18, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0,
		0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 5, 6, 0, 0, 98, 100, 5, 18, 0, 0, 99,
		98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 3,
		16, 8, 0, 102, 104, 5, 18, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 15, 1, 0, 0, 0, 105, 108, 3, 20, 10, 0, 106, 108, 3, 18, 9,
		0, 107, 105, 1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 17, 1, 0, 0, 0, 109,
		111, 5, 7, 0, 0, 110, 112, 5, 18, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 124, 1, 0, 0, 0, 113, 118, 3, 20, 10, 0, 114, 115, 5,
		17, 0, 0, 115, 117, 3, 20, 10, 0, 116, 114, 1, 0, 0, 0, 117, 120, 1, 0,
		0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 121, 123, 5, 17, 0, 0, 122, 121, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 113, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 8, 0, 0, 127, 19, 1, 0,
		0, 0, 128, 134, 5, 10, 0, 0, 129, 134, 5, 11, 0, 0, 130, 134, 5, 14, 0,
		0, 131, 134, 7, 0, 0, 0, 132, 134, 5, 9, 0, 0, 133, 128, 1, 0, 0, 0, 133,
		129, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 132,
		1, 0, 0, 0, 134, 21, 1, 0, 0, 0, 22, 23, 29, 43, 47, 53, 57, 60, 65, 69,
		76, 83, 87, 89, 95, 99, 103, 107, 111, 118, 122, 124, 133,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// UIDLParserInit initializes any static state used to implement UIDLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewUIDLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UIDLParserInit() {
	staticData := &UIDLParserStaticData
	staticData.once.Do(uidlParserInit)
}

// NewUIDLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewUIDLParser(input antlr.TokenStream) *UIDLParser {
	UIDLParserInit()
	this := new(UIDLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &UIDLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "UIDL.g4"

	return this
}

// UIDLParser tokens.
const (
	UIDLParserEOF                = antlr.TokenEOF
	UIDLParserT__0               = 1
	UIDLParserT__1               = 2
	UIDLParserT__2               = 3
	UIDLParserT__3               = 4
	UIDLParserT__4               = 5
	UIDLParserT__5               = 6
	UIDLParserT__6               = 7
	UIDLParserT__7               = 8
	UIDLParserBool               = 9
	UIDLParserDoubleQuotedString = 10
	UIDLParserBackQuotedString   = 11
	UIDLParserIdentifier         = 12
	UIDLParserNatural            = 13
	UIDLParserFloat              = 14
	UIDLParserInt                = 15
	UIDLParserSemicolon          = 16
	UIDLParserComma              = 17
	UIDLParserWhiteSpace         = 18
)

// UIDLParser rules.
const (
	UIDLParserRULE_uidl             = 0
	UIDLParserRULE_version          = 1
	UIDLParserRULE_commands         = 2
	UIDLParserRULE_command          = 3
	UIDLParserRULE_commandSeparator = 4
	UIDLParserRULE_commandBody      = 5
	UIDLParserRULE_attributes       = 6
	UIDLParserRULE_attribute        = 7
	UIDLParserRULE_value            = 8
	UIDLParserRULE_listValue        = 9
	UIDLParserRULE_simpleValue      = 10
)

// IUidlContext is an interface to support dynamic dispatch.
type IUidlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Version() IVersionContext
	CommandSeparator() ICommandSeparatorContext
	Commands() ICommandsContext
	EOF() antlr.TerminalNode
	AllWhiteSpace() []antlr.TerminalNode
	WhiteSpace(i int) antlr.TerminalNode

	// IsUidlContext differentiates from other interfaces.
	IsUidlContext()
}

type UidlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUidlContext() *UidlContext {
	var p = new(UidlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_uidl
	return p
}

func InitEmptyUidlContext(p *UidlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_uidl
}

func (*UidlContext) IsUidlContext() {}

func NewUidlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UidlContext {
	var p = new(UidlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_uidl

	return p
}

func (s *UidlContext) GetParser() antlr.Parser { return s.parser }

func (s *UidlContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *UidlContext) CommandSeparator() ICommandSeparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandSeparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandSeparatorContext)
}

func (s *UidlContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *UidlContext) EOF() antlr.TerminalNode {
	return s.GetToken(UIDLParserEOF, 0)
}

func (s *UidlContext) AllWhiteSpace() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserWhiteSpace)
}

func (s *UidlContext) WhiteSpace(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, i)
}

func (s *UidlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UidlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Uidl() (localctx IUidlContext) {
	localctx = NewUidlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UIDLParserRULE_uidl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(22)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(25)
		p.Version()
	}
	{
		p.SetState(26)
		p.CommandSeparator()
	}
	{
		p.SetState(27)
		p.Commands()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(28)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(31)
		p.Match(UIDLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WhiteSpace() antlr.TerminalNode
	Natural() antlr.TerminalNode

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) WhiteSpace() antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, 0)
}

func (s *VersionContext) Natural() antlr.TerminalNode {
	return s.GetToken(UIDLParserNatural, 0)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UIDLParserRULE_version)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(UIDLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(UIDLParserWhiteSpace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(UIDLParserNatural)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext
	AllCommandSeparator() []ICommandSeparatorContext
	CommandSeparator(i int) ICommandSeparatorContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) AllCommandSeparator() []ICommandSeparatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandSeparatorContext); ok {
			len++
		}
	}

	tst := make([]ICommandSeparatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandSeparatorContext); ok {
			tst[i] = t.(ICommandSeparatorContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) CommandSeparator(i int) ICommandSeparatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandSeparatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandSeparatorContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Commands() (localctx ICommandsContext) {
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UIDLParserRULE_commands)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Command()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(38)
				p.CommandSeparator()
			}
			{
				p.SetState(39)
				p.Command()
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(46)
			p.CommandSeparator()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllWhiteSpace() []antlr.TerminalNode
	WhiteSpace(i int) antlr.TerminalNode
	Attributes() IAttributesContext
	CommandBody() ICommandBodyContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserIdentifier)
}

func (s *CommandContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserIdentifier, i)
}

func (s *CommandContext) AllWhiteSpace() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserWhiteSpace)
}

func (s *CommandContext) WhiteSpace(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, i)
}

func (s *CommandContext) Attributes() IAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *CommandContext) CommandBody() ICommandBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandBodyContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UIDLParserRULE_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(UIDLParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(UIDLParserWhiteSpace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(UIDLParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(52)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(55)
		p.Attributes()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(56)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserT__1 {
		{
			p.SetState(59)
			p.CommandBody()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandSeparatorContext is an interface to support dynamic dispatch.
type ICommandSeparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWs returns the ws token.
	GetWs() antlr.Token

	// SetWs sets the ws token.
	SetWs(antlr.Token)

	// Getter signatures
	Semicolon() antlr.TerminalNode
	WhiteSpace() antlr.TerminalNode

	// IsCommandSeparatorContext differentiates from other interfaces.
	IsCommandSeparatorContext()
}

type CommandSeparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ws     antlr.Token
}

func NewEmptyCommandSeparatorContext() *CommandSeparatorContext {
	var p = new(CommandSeparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commandSeparator
	return p
}

func InitEmptyCommandSeparatorContext(p *CommandSeparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commandSeparator
}

func (*CommandSeparatorContext) IsCommandSeparatorContext() {}

func NewCommandSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandSeparatorContext {
	var p = new(CommandSeparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_commandSeparator

	return p
}

func (s *CommandSeparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandSeparatorContext) GetWs() antlr.Token { return s.ws }

func (s *CommandSeparatorContext) SetWs(v antlr.Token) { s.ws = v }

func (s *CommandSeparatorContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(UIDLParserSemicolon, 0)
}

func (s *CommandSeparatorContext) WhiteSpace() antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, 0)
}

func (s *CommandSeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandSeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) CommandSeparator() (localctx ICommandSeparatorContext) {
	localctx = NewCommandSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UIDLParserRULE_commandSeparator)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case UIDLParserSemicolon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(UIDLParserSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UIDLParserWhiteSpace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)

			var _m = p.Match(UIDLParserWhiteSpace)

			localctx.(*CommandSeparatorContext).ws = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)

		if !(strings.ContainsRune((func() string {
			if localctx.(*CommandSeparatorContext).GetWs() == nil {
				return ""
			} else {
				return localctx.(*CommandSeparatorContext).GetWs().GetText()
			}
		}()), '\n')) {
			p.SetError(antlr.NewFailedPredicateException(p, "strings.ContainsRune($ws.text, '\\n')", "expected semicolon or new line"))
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandBodyContext is an interface to support dynamic dispatch.
type ICommandBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Commands() ICommandsContext
	WhiteSpace() antlr.TerminalNode

	// IsCommandBodyContext differentiates from other interfaces.
	IsCommandBodyContext()
}

type CommandBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandBodyContext() *CommandBodyContext {
	var p = new(CommandBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commandBody
	return p
}

func InitEmptyCommandBodyContext(p *CommandBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_commandBody
}

func (*CommandBodyContext) IsCommandBodyContext() {}

func NewCommandBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandBodyContext {
	var p = new(CommandBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_commandBody

	return p
}

func (s *CommandBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandBodyContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *CommandBodyContext) WhiteSpace() antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, 0)
}

func (s *CommandBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) CommandBody() (localctx ICommandBodyContext) {
	localctx = NewCommandBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UIDLParserRULE_commandBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(UIDLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(68)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(71)
		p.Commands()
	}
	{
		p.SetState(72)
		p.Match(UIDLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributesContext is an interface to support dynamic dispatch.
type IAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WhiteSpace() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsAttributesContext differentiates from other interfaces.
	IsAttributesContext()
}

type AttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesContext() *AttributesContext {
	var p = new(AttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_attributes
	return p
}

func InitEmptyAttributesContext(p *AttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_attributes
}

func (*AttributesContext) IsAttributesContext() {}

func NewAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesContext {
	var p = new(AttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_attributes

	return p
}

func (s *AttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesContext) WhiteSpace() antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, 0)
}

func (s *AttributesContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *AttributesContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *AttributesContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserComma)
}

func (s *AttributesContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserComma, i)
}

func (s *AttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Attributes() (localctx IAttributesContext) {
	localctx = NewAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, UIDLParserRULE_attributes)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(UIDLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(75)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserIdentifier {
		{
			p.SetState(78)
			p.Attribute()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(79)
					p.Match(UIDLParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.Attribute()
				}

			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == UIDLParserComma {
			{
				p.SetState(86)
				p.Match(UIDLParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(91)
		p.Match(UIDLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Value() IValueContext
	AllWhiteSpace() []antlr.TerminalNode
	WhiteSpace(i int) antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(UIDLParserIdentifier, 0)
}

func (s *AttributeContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AttributeContext) AllWhiteSpace() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserWhiteSpace)
}

func (s *AttributeContext) WhiteSpace(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, i)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UIDLParserRULE_attribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(UIDLParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(94)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(97)
		p.Match(UIDLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(98)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(101)
		p.Value()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(102)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleValue() ISimpleValueContext
	ListValue() IListValueContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) SimpleValue() ISimpleValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, UIDLParserRULE_value)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case UIDLParserBool, UIDLParserDoubleQuotedString, UIDLParserBackQuotedString, UIDLParserNatural, UIDLParserFloat, UIDLParserInt:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.SimpleValue()
		}

	case UIDLParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.ListValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WhiteSpace() antlr.TerminalNode
	AllSimpleValue() []ISimpleValueContext
	SimpleValue(i int) ISimpleValueContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_listValue
	return p
}

func InitEmptyListValueContext(p *ListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_listValue
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) WhiteSpace() antlr.TerminalNode {
	return s.GetToken(UIDLParserWhiteSpace, 0)
}

func (s *ListValueContext) AllSimpleValue() []ISimpleValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleValueContext); ok {
			len++
		}
	}

	tst := make([]ISimpleValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleValueContext); ok {
			tst[i] = t.(ISimpleValueContext)
			i++
		}
	}

	return tst
}

func (s *ListValueContext) SimpleValue(i int) ISimpleValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleValueContext)
}

func (s *ListValueContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(UIDLParserComma)
}

func (s *ListValueContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(UIDLParserComma, i)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) ListValue() (localctx IListValueContext) {
	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, UIDLParserRULE_listValue)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(UIDLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == UIDLParserWhiteSpace {
		{
			p.SetState(110)
			p.Match(UIDLParserWhiteSpace)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60928) != 0 {
		{
			p.SetState(113)
			p.SimpleValue()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(114)
					p.Match(UIDLParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(115)
					p.SimpleValue()
				}

			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == UIDLParserComma {
			{
				p.SetState(121)
				p.Match(UIDLParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(126)
		p.Match(UIDLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleValueContext is an interface to support dynamic dispatch.
type ISimpleValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DoubleQuotedString() antlr.TerminalNode
	BackQuotedString() antlr.TerminalNode
	Float() antlr.TerminalNode
	Natural() antlr.TerminalNode
	Int() antlr.TerminalNode
	Bool() antlr.TerminalNode

	// IsSimpleValueContext differentiates from other interfaces.
	IsSimpleValueContext()
}

type SimpleValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleValueContext() *SimpleValueContext {
	var p = new(SimpleValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_simpleValue
	return p
}

func InitEmptySimpleValueContext(p *SimpleValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = UIDLParserRULE_simpleValue
}

func (*SimpleValueContext) IsSimpleValueContext() {}

func NewSimpleValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleValueContext {
	var p = new(SimpleValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = UIDLParserRULE_simpleValue

	return p
}

func (s *SimpleValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleValueContext) DoubleQuotedString() antlr.TerminalNode {
	return s.GetToken(UIDLParserDoubleQuotedString, 0)
}

func (s *SimpleValueContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(UIDLParserBackQuotedString, 0)
}

func (s *SimpleValueContext) Float() antlr.TerminalNode {
	return s.GetToken(UIDLParserFloat, 0)
}

func (s *SimpleValueContext) Natural() antlr.TerminalNode {
	return s.GetToken(UIDLParserNatural, 0)
}

func (s *SimpleValueContext) Int() antlr.TerminalNode {
	return s.GetToken(UIDLParserInt, 0)
}

func (s *SimpleValueContext) Bool() antlr.TerminalNode {
	return s.GetToken(UIDLParserBool, 0)
}

func (s *SimpleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *UIDLParser) SimpleValue() (localctx ISimpleValueContext) {
	localctx = NewSimpleValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, UIDLParserRULE_simpleValue)
	var _la int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case UIDLParserDoubleQuotedString:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(UIDLParserDoubleQuotedString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UIDLParserBackQuotedString:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(UIDLParserBackQuotedString)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UIDLParserFloat:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(UIDLParserFloat)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case UIDLParserNatural, UIDLParserInt:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			_la = p.GetTokenStream().LA(1)

			if !(_la == UIDLParserNatural || _la == UIDLParserInt) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case UIDLParserBool:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.Match(UIDLParserBool)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *UIDLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *CommandSeparatorContext = nil
		if localctx != nil {
			t = localctx.(*CommandSeparatorContext)
		}
		return p.CommandSeparator_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *UIDLParser) CommandSeparator_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return strings.ContainsRune((func() string {
			if localctx.(*CommandSeparatorContext).GetWs() == nil {
				return ""
			} else {
				return localctx.(*CommandSeparatorContext).GetWs().GetText()
			}
		}()), '\n')

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
