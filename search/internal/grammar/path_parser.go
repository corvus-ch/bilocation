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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 109, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 5, 2, 18, 10, 2, 3, 2, 3, 2, 5, 2, 22, 10, 2, 3, 2, 5, 2,
	25, 10, 2, 3, 2, 3, 2, 5, 2, 29, 10, 2, 3, 2, 5, 2, 32, 10, 2, 3, 2, 3,
	2, 5, 2, 36, 10, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 3, 2, 5, 2, 43, 10, 2,
	3, 2, 5, 2, 46, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 52, 10, 3, 13, 3,
	14, 3, 53, 5, 3, 56, 10, 3, 3, 4, 3, 4, 3, 4, 6, 4, 61, 10, 4, 13, 4, 14,
	4, 62, 3, 4, 3, 4, 3, 4, 6, 4, 68, 10, 4, 13, 4, 14, 4, 69, 3, 4, 3, 4,
	3, 4, 6, 4, 75, 10, 4, 13, 4, 14, 4, 76, 3, 4, 3, 4, 3, 4, 6, 4, 82, 10,
	4, 13, 4, 14, 4, 83, 5, 4, 86, 10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 91, 10, 5,
	13, 5, 14, 5, 92, 3, 6, 3, 6, 5, 6, 97, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 107, 10, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10,
	12, 14, 2, 2, 2, 125, 2, 45, 3, 2, 2, 2, 4, 55, 3, 2, 2, 2, 6, 85, 3, 2,
	2, 2, 8, 87, 3, 2, 2, 2, 10, 96, 3, 2, 2, 2, 12, 98, 3, 2, 2, 2, 14, 106,
	3, 2, 2, 2, 16, 18, 7, 3, 2, 2, 17, 16, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2,
	18, 19, 3, 2, 2, 2, 19, 21, 5, 6, 4, 2, 20, 22, 7, 3, 2, 2, 21, 20, 3,
	2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 46, 3, 2, 2, 2, 23, 25, 7, 3, 2, 2, 24,
	23, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 5, 8, 5,
	2, 27, 29, 7, 3, 2, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 46,
	3, 2, 2, 2, 30, 32, 7, 3, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 33, 3, 2, 2, 2, 33, 35, 5, 10, 6, 2, 34, 36, 7, 3, 2, 2, 35, 34, 3,
	2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 46, 3, 2, 2, 2, 37, 39, 7, 3, 2, 2, 38,
	37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 5, 4, 3,
	2, 41, 43, 7, 3, 2, 2, 42, 41, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 46,
	3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 17, 3, 2, 2, 2, 45, 24, 3, 2, 2, 2,
	45, 31, 3, 2, 2, 2, 45, 38, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 3, 3, 2,
	2, 2, 47, 56, 7, 4, 2, 2, 48, 51, 7, 4, 2, 2, 49, 50, 7, 3, 2, 2, 50, 52,
	7, 4, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 47, 3, 2, 2, 2, 55, 48, 3,
	2, 2, 2, 56, 5, 3, 2, 2, 2, 57, 60, 5, 8, 5, 2, 58, 59, 7, 3, 2, 2, 59,
	61, 5, 8, 5, 2, 60, 58, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 60, 3, 2, 2,
	2, 62, 63, 3, 2, 2, 2, 63, 86, 3, 2, 2, 2, 64, 67, 5, 8, 5, 2, 65, 66,
	7, 3, 2, 2, 66, 68, 5, 10, 6, 2, 67, 65, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 86, 3, 2, 2, 2, 71, 74, 5,
	10, 6, 2, 72, 73, 7, 3, 2, 2, 73, 75, 5, 8, 5, 2, 74, 72, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 86, 3, 2, 2,
	2, 78, 81, 5, 10, 6, 2, 79, 80, 7, 3, 2, 2, 80, 82, 5, 10, 6, 2, 81, 79,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2,
	84, 86, 3, 2, 2, 2, 85, 57, 3, 2, 2, 2, 85, 64, 3, 2, 2, 2, 85, 71, 3,
	2, 2, 2, 85, 78, 3, 2, 2, 2, 86, 7, 3, 2, 2, 2, 87, 90, 5, 10, 6, 2, 88,
	89, 7, 6, 2, 2, 89, 91, 5, 10, 6, 2, 90, 88, 3, 2, 2, 2, 91, 92, 3, 2,
	2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 9, 3, 2, 2, 2, 94, 97,
	5, 12, 7, 2, 95, 97, 5, 14, 8, 2, 96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2,
	2, 97, 11, 3, 2, 2, 2, 98, 99, 7, 7, 2, 2, 99, 100, 7, 5, 2, 2, 100, 101,
	7, 4, 2, 2, 101, 13, 3, 2, 2, 2, 102, 107, 7, 7, 2, 2, 103, 104, 7, 7,
	2, 2, 104, 105, 7, 5, 2, 2, 105, 107, 7, 7, 2, 2, 106, 102, 3, 2, 2, 2,
	106, 103, 3, 2, 2, 2, 107, 15, 3, 2, 2, 2, 21, 17, 21, 24, 28, 31, 35,
	38, 42, 45, 53, 55, 62, 69, 76, 83, 85, 92, 96, 106,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'/'", "", "'='", "','",
}
var symbolicNames = []string{
	"", "AND", "ANY", "EQ", "OR", "STRING",
}

var ruleNames = []string{
	"query", "anyExpr", "andExpr", "orExpr", "tag", "classifierTagExpr", "nameTagExpr",
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
	PathParserOR     = 4
	PathParserSTRING = 5
)

// PathParser rules.
const (
	PathParserRULE_query             = 0
	PathParserRULE_anyExpr           = 1
	PathParserRULE_andExpr           = 2
	PathParserRULE_orExpr            = 3
	PathParserRULE_tag               = 4
	PathParserRULE_classifierTagExpr = 5
	PathParserRULE_nameTagExpr       = 6
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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(14)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(17)
			p.AndExpr()
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(18)
				p.Match(PathParserAND)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(21)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(24)
			p.OrExpr()
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(25)
				p.Match(PathParserAND)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(28)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(31)
			p.Tag()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(32)
				p.Match(PathParserAND)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(35)
				p.Match(PathParserAND)
			}

		}
		{
			p.SetState(38)
			p.AnyExpr()
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PathParserAND {
			{
				p.SetState(39)
				p.Match(PathParserAND)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)

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
	p.EnterRule(localctx, 2, PathParserRULE_anyExpr)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(PathParserANY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(PathParserANY)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(47)
					p.Match(PathParserAND)
				}
				{
					p.SetState(48)
					p.Match(PathParserANY)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 4, PathParserRULE_andExpr)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.OrExpr()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(56)
					p.Match(PathParserAND)
				}
				{
					p.SetState(57)
					p.OrExpr()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.OrExpr()
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(63)
					p.Match(PathParserAND)
				}
				{
					p.SetState(64)
					p.Tag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Tag()
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(70)
					p.Match(PathParserAND)
				}
				{
					p.SetState(71)
					p.OrExpr()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Tag()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(77)
					p.Match(PathParserAND)
				}
				{
					p.SetState(78)
					p.Tag()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, PathParserRULE_orExpr)
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
		p.SetState(85)
		p.Tag()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PathParserOR {
		{
			p.SetState(86)
			p.Match(PathParserOR)
		}
		{
			p.SetState(87)
			p.Tag()
		}

		p.SetState(90)
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
	p.EnterRule(localctx, 8, PathParserRULE_tag)

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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.ClassifierTagExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
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

func (s *ClassifierTagExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(PathParserEQ, 0)
}

func (s *ClassifierTagExprContext) ANY() antlr.TerminalNode {
	return s.GetToken(PathParserANY, 0)
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
	p.EnterRule(localctx, 10, PathParserRULE_classifierTagExpr)

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

		var _m = p.Match(PathParserSTRING)

		localctx.(*ClassifierTagExprContext).classifier = _m
	}
	{
		p.SetState(97)

		var _m = p.Match(PathParserEQ)

		localctx.(*ClassifierTagExprContext).op = _m
	}
	{
		p.SetState(98)

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

	// GetClassifier returns the classifier token.
	GetClassifier() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetClassifier sets the classifier token.
	SetClassifier(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsNameTagExprContext differentiates from other interfaces.
	IsNameTagExprContext()
}

type NameTagExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       antlr.Token
	classifier antlr.Token
	op         antlr.Token
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

func (s *NameTagExprContext) GetClassifier() antlr.Token { return s.classifier }

func (s *NameTagExprContext) GetOp() antlr.Token { return s.op }

func (s *NameTagExprContext) SetName(v antlr.Token) { s.name = v }

func (s *NameTagExprContext) SetClassifier(v antlr.Token) { s.classifier = v }

func (s *NameTagExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *NameTagExprContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PathParserSTRING)
}

func (s *NameTagExprContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PathParserSTRING, i)
}

func (s *NameTagExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(PathParserEQ, 0)
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
	p.EnterRule(localctx, 12, PathParserRULE_nameTagExpr)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).name = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).classifier = _m
		}
		{
			p.SetState(102)

			var _m = p.Match(PathParserEQ)

			localctx.(*NameTagExprContext).op = _m
		}
		{
			p.SetState(103)

			var _m = p.Match(PathParserSTRING)

			localctx.(*NameTagExprContext).name = _m
		}

	}

	return localctx
}
