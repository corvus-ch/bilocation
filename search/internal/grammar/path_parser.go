// Code generated from search/internal/grammar/Path.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // Path
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 122,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 5, 2, 24, 10,
	2, 3, 2, 5, 2, 27, 10, 2, 3, 2, 3, 2, 5, 2, 31, 10, 2, 3, 2, 5, 2, 34,
	10, 2, 3, 2, 3, 2, 5, 2, 38, 10, 2, 3, 2, 5, 2, 41, 10, 2, 3, 2, 3, 2,
	5, 2, 45, 10, 2, 3, 2, 5, 2, 48, 10, 2, 3, 2, 3, 2, 5, 2, 52, 10, 2, 3,
	2, 5, 2, 55, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 63, 10, 4,
	13, 4, 14, 4, 64, 5, 4, 67, 10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 72, 10, 5, 13,
	5, 14, 5, 73, 3, 5, 3, 5, 3, 5, 6, 5, 79, 10, 5, 13, 5, 14, 5, 80, 3, 5,
	3, 5, 3, 5, 6, 5, 86, 10, 5, 13, 5, 14, 5, 87, 3, 5, 3, 5, 3, 5, 6, 5,
	93, 10, 5, 13, 5, 14, 5, 94, 5, 5, 97, 10, 5, 3, 6, 3, 6, 3, 6, 6, 6, 102,
	10, 6, 13, 6, 14, 6, 103, 3, 7, 3, 7, 5, 7, 108, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 120, 10, 9, 3, 9, 2,
	2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 3, 3, 2, 5, 6, 2, 141, 2, 54, 3,
	2, 2, 2, 4, 56, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 8, 96, 3, 2, 2, 2, 10, 98,
	3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 119, 3, 2, 2,
	2, 18, 20, 7, 3, 2, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21,
	3, 2, 2, 2, 21, 23, 5, 8, 5, 2, 22, 24, 7, 3, 2, 2, 23, 22, 3, 2, 2, 2,
	23, 24, 3, 2, 2, 2, 24, 55, 3, 2, 2, 2, 25, 27, 7, 3, 2, 2, 26, 25, 3,
	2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 5, 10, 6, 2, 29,
	31, 7, 3, 2, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 55, 3, 2, 2,
	2, 32, 34, 7, 3, 2, 2, 33, 32, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35,
	3, 2, 2, 2, 35, 37, 5, 12, 7, 2, 36, 38, 7, 3, 2, 2, 37, 36, 3, 2, 2, 2,
	37, 38, 3, 2, 2, 2, 38, 55, 3, 2, 2, 2, 39, 41, 7, 3, 2, 2, 40, 39, 3,
	2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 5, 6, 4, 2, 43,
	45, 7, 3, 2, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 55, 3, 2, 2,
	2, 46, 48, 7, 3, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49,
	3, 2, 2, 2, 49, 51, 5, 4, 3, 2, 50, 52, 7, 3, 2, 2, 51, 50, 3, 2, 2, 2,
	51, 52, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 19, 3,
	2, 2, 2, 54, 26, 3, 2, 2, 2, 54, 33, 3, 2, 2, 2, 54, 40, 3, 2, 2, 2, 54,
	47, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 3, 3, 2, 2, 2, 56, 57, 7, 7, 2,
	2, 57, 5, 3, 2, 2, 2, 58, 67, 7, 4, 2, 2, 59, 62, 7, 4, 2, 2, 60, 61, 7,
	3, 2, 2, 61, 63, 7, 4, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64,
	62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 58, 3, 2, 2,
	2, 66, 59, 3, 2, 2, 2, 67, 7, 3, 2, 2, 2, 68, 71, 5, 10, 6, 2, 69, 70,
	7, 3, 2, 2, 70, 72, 5, 10, 6, 2, 71, 69, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2,
	73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 97, 3, 2, 2, 2, 75, 78, 5,
	10, 6, 2, 76, 77, 7, 3, 2, 2, 77, 79, 5, 12, 7, 2, 78, 76, 3, 2, 2, 2,
	79, 80, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 97, 3,
	2, 2, 2, 82, 85, 5, 12, 7, 2, 83, 84, 7, 3, 2, 2, 84, 86, 5, 10, 6, 2,
	85, 83, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 97, 3, 2, 2, 2, 89, 92, 5, 12, 7, 2, 90, 91, 7, 3, 2, 2, 91,
	93, 5, 12, 7, 2, 92, 90, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 92, 3, 2,
	2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 68, 3, 2, 2, 2, 96, 75,
	3, 2, 2, 2, 96, 82, 3, 2, 2, 2, 96, 89, 3, 2, 2, 2, 97, 9, 3, 2, 2, 2,
	98, 101, 5, 12, 7, 2, 99, 100, 7, 9, 2, 2, 100, 102, 5, 12, 7, 2, 101,
	99, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3,
	2, 2, 2, 104, 11, 3, 2, 2, 2, 105, 108, 5, 14, 8, 2, 106, 108, 5, 16, 9,
	2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 13, 3, 2, 2, 2, 109,
	110, 7, 10, 2, 2, 110, 111, 9, 2, 2, 2, 111, 112, 7, 4, 2, 2, 112, 15,
	3, 2, 2, 2, 113, 120, 7, 10, 2, 2, 114, 115, 7, 8, 2, 2, 115, 120, 7, 10,
	2, 2, 116, 117, 7, 10, 2, 2, 117, 118, 9, 2, 2, 2, 118, 120, 7, 10, 2,
	2, 119, 113, 3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 120,
	17, 3, 2, 2, 2, 23, 19, 23, 26, 30, 33, 37, 40, 44, 47, 51, 54, 64, 66,
	73, 80, 87, 94, 96, 103, 107, 119,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'/'", "", "'='", "'!='", "'!*'", "'!'", "','",
}
var symbolicNames = []string{
	"", "AND", "ANY", "EQ", "NEQ", "NONE", "NOT", "OR", "STRING",
}

var ruleNames = []string{
	"query", "noneExpr", "anyExpr", "andExpr", "orExpr", "tag", "classifierTagExpr",
	"nameTagExpr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PathParser struct {
	*antlr.BaseParser
}

func NewPathParser(input antlr.TokenStream) *PathParser {
	this := new(PathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Path.g4"

	return this
}

// PathParser tokens.
const (
	PathParserEOF    = antlr.TokenEOF
	PathParserAND    = 1
	PathParserANY    = 2
	PathParserEQ     = 3
	PathParserNEQ    = 4
	PathParserNONE   = 5
	PathParserNOT    = 6
	PathParserOR     = 7
	PathParserSTRING = 8
)

// PathParser rules.
const (
	PathParserRULE_query             = 0
	PathParserRULE_noneExpr          = 1
	PathParserRULE_anyExpr           = 2
	PathParserRULE_andExpr           = 3
	PathParserRULE_orExpr            = 4
	PathParserRULE_tag               = 5
	PathParserRULE_classifierTagExpr = 6
	PathParserRULE_nameTagExpr       = 7
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) AndExpr() IAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *QueryContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(PathParserAND)
}

func (s *QueryContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(PathParserAND, i)
}

func (s *QueryContext) OrExpr() IOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *QueryContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *QueryContext) AnyExpr() IAnyExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyExprContext)
}

func (s *QueryContext) NoneExpr() INoneExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoneExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoneExprContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *PathParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PathParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(16)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(19)
			p.AndExpr()
		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(20)
				p.Match(PathParserAND)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(23)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(26)
			p.OrExpr()
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(27)
				p.Match(PathParserAND)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(30)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(33)
			p.Tag()
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(34)
				p.Match(PathParserAND)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(37)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(40)
			p.AnyExpr()
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(41)
				p.Match(PathParserAND)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(44)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(47)
			p.NoneExpr()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(48)
				p.Match(PathParserAND)
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)

	}

	return localctx
}

// INoneExprContext is an interface to support dynamic dispatch.
type INoneExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoneExprContext differentiates from other interfaces.
	IsNoneExprContext()
}

type NoneExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoneExprContext() *NoneExprContext {
	var p = new(NoneExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_noneExpr
	return p
}

func (*NoneExprContext) IsNoneExprContext() {}

func NewNoneExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoneExprContext {
	var p = new(NoneExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_noneExpr

	return p
}

func (s *NoneExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NoneExprContext) NONE() antlr.TerminalNode {
	return s.GetToken(PathParserNONE, 0)
}

func (s *NoneExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoneExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoneExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterNoneExpr(s)
	}
}

func (s *NoneExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitNoneExpr(s)
	}
}

func (p *PathParser) NoneExpr() (localctx INoneExprContext) {
	localctx = NewNoneExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PathParserRULE_noneExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(PathParserNONE)
	}

	return localctx
}

// IAnyExprContext is an interface to support dynamic dispatch.
type IAnyExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyExprContext differentiates from other interfaces.
	IsAnyExprContext()
}

type AnyExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyExprContext() *AnyExprContext {
	var p = new(AnyExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_anyExpr
	return p
}

func (*AnyExprContext) IsAnyExprContext() {}

func NewAnyExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyExprContext {
	var p = new(AnyExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_anyExpr

	return p
}

func (s *AnyExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyExprContext) AllANY() []antlr.TerminalNode {
	return s.GetTokens(PathParserANY)
}

func (s *AnyExprContext) ANY(i int) antlr.TerminalNode {
	return s.GetToken(PathParserANY, i)
}

func (s *AnyExprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(PathParserAND)
}

func (s *AnyExprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(PathParserAND, i)
}

func (s *AnyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterAnyExpr(s)
	}
}

func (s *AnyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitAnyExpr(s)
	}
}

func (p *PathParser) AnyExpr() (localctx IAnyExprContext) {
	localctx = NewAnyExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PathParserRULE_anyExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(PathParserANY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(PathParserANY)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(58)
					p.Match(PathParserAND)
				}
				{
					p.SetState(59)
					p.Match(PathParserANY)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_andExpr
	return p
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) AllOrExpr() []IOrExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrExprContext)(nil)).Elem())
	var tst = make([]IOrExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) OrExpr(i int) IOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *AndExprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(PathParserAND)
}

func (s *AndExprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(PathParserAND, i)
}

func (s *AndExprContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *AndExprContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *PathParser) AndExpr() (localctx IAndExprContext) {
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PathParserRULE_andExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.OrExpr()
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(67)
					p.Match(PathParserAND)
				}
				{
					p.SetState(68)
					p.OrExpr()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.OrExpr()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(74)
					p.Match(PathParserAND)
				}
				{
					p.SetState(75)
					p.Tag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Tag()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(81)
					p.Match(PathParserAND)
				}
				{
					p.SetState(82)
					p.OrExpr()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Tag()
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(88)
					p.Match(PathParserAND)
				}
				{
					p.SetState(89)
					p.Tag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_orExpr
	return p
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *OrExprContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *OrExprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(PathParserOR)
}

func (s *OrExprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(PathParserOR, i)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (p *PathParser) OrExpr() (localctx IOrExprContext) {
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PathParserRULE_orExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Tag()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PathParserOR {
		{
			p.SetState(97)
			p.Match(PathParserOR)
		}
		{
			p.SetState(98)
			p.Tag()
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) ClassifierTagExpr() IClassifierTagExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassifierTagExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassifierTagExprContext)
}

func (s *TagContext) NameTagExpr() INameTagExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameTagExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameTagExprContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *PathParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PathParserRULE_tag)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.ClassifierTagExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.NameTagExpr()
		}

	}

	return localctx
}

// IClassifierTagExprContext is an interface to support dynamic dispatch.
type IClassifierTagExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetClassifier returns the classifier token.
	GetClassifier() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetClassifier sets the classifier token.
	SetClassifier(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsClassifierTagExprContext differentiates from other interfaces.
	IsClassifierTagExprContext()
}

type ClassifierTagExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	classifier antlr.Token
	op         antlr.Token
	name       antlr.Token
}

func NewEmptyClassifierTagExprContext() *ClassifierTagExprContext {
	var p = new(ClassifierTagExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_classifierTagExpr
	return p
}

func (*ClassifierTagExprContext) IsClassifierTagExprContext() {}

func NewClassifierTagExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassifierTagExprContext {
	var p = new(ClassifierTagExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_classifierTagExpr

	return p
}

func (s *ClassifierTagExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassifierTagExprContext) GetClassifier() antlr.Token { return s.classifier }

func (s *ClassifierTagExprContext) GetOp() antlr.Token { return s.op }

func (s *ClassifierTagExprContext) GetName() antlr.Token { return s.name }

func (s *ClassifierTagExprContext) SetClassifier(v antlr.Token) { s.classifier = v }

func (s *ClassifierTagExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ClassifierTagExprContext) SetName(v antlr.Token) { s.name = v }

func (s *ClassifierTagExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(PathParserSTRING, 0)
}

func (s *ClassifierTagExprContext) ANY() antlr.TerminalNode {
	return s.GetToken(PathParserANY, 0)
}

func (s *ClassifierTagExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(PathParserEQ, 0)
}

func (s *ClassifierTagExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PathParserNEQ, 0)
}

func (s *ClassifierTagExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassifierTagExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassifierTagExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterClassifierTagExpr(s)
	}
}

func (s *ClassifierTagExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitClassifierTagExpr(s)
	}
}

func (p *PathParser) ClassifierTagExpr() (localctx IClassifierTagExprContext) {
	localctx = NewClassifierTagExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PathParserRULE_classifierTagExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)

		var _m = p.Match(PathParserSTRING)

		localctx.(*ClassifierTagExprContext).classifier = _m
	}
	{
		p.SetState(108)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ClassifierTagExprContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == PathParserEQ || _la == PathParserNEQ) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ClassifierTagExprContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(109)

		var _m = p.Match(PathParserANY)

		localctx.(*ClassifierTagExprContext).name = _m
	}

	return localctx
}

// INameTagExprContext is an interface to support dynamic dispatch.
type INameTagExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetClassifier returns the classifier token.
	GetClassifier() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetClassifier sets the classifier token.
	SetClassifier(antlr.Token)

	// IsNameTagExprContext differentiates from other interfaces.
	IsNameTagExprContext()
}

type NameTagExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       antlr.Token
	op         antlr.Token
	classifier antlr.Token
}

func NewEmptyNameTagExprContext() *NameTagExprContext {
	var p = new(NameTagExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PathParserRULE_nameTagExpr
	return p
}

func (*NameTagExprContext) IsNameTagExprContext() {}

func NewNameTagExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameTagExprContext {
	var p = new(NameTagExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PathParserRULE_nameTagExpr

	return p
}

func (s *NameTagExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NameTagExprContext) GetName() antlr.Token { return s.name }

func (s *NameTagExprContext) GetOp() antlr.Token { return s.op }

func (s *NameTagExprContext) GetClassifier() antlr.Token { return s.classifier }

func (s *NameTagExprContext) SetName(v antlr.Token) { s.name = v }

func (s *NameTagExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *NameTagExprContext) SetClassifier(v antlr.Token) { s.classifier = v }

func (s *NameTagExprContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PathParserSTRING)
}

func (s *NameTagExprContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PathParserSTRING, i)
}

func (s *NameTagExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(PathParserNOT, 0)
}

func (s *NameTagExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(PathParserEQ, 0)
}

func (s *NameTagExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PathParserNEQ, 0)
}

func (s *NameTagExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameTagExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameTagExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.EnterNameTagExpr(s)
	}
}

func (s *NameTagExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PathListener); ok {
		listenerT.ExitNameTagExpr(s)
	}
}

func (p *PathParser) NameTagExpr() (localctx INameTagExprContext) {
	localctx = NewNameTagExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PathParserRULE_nameTagExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).name = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)

			var _m = p.Match(PathParserNOT)

			localctx.(*NameTagExprContext).op = _m
		}
		{
			p.SetState(113)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).name = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).classifier = _m
		}
		{
			p.SetState(115)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*NameTagExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PathParserEQ || _la == PathParserNEQ) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*NameTagExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(116)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).name = _m
		}

	}

	return localctx
}
