// Code generated from search/internal/grammar/Path.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar // Path
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePathListener is a complete listener for a parse tree produced by PathParser.
type BasePathListener struct{}

var _ PathListener = &BasePathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BasePathListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasePathListener) ExitQuery(ctx *QueryContext) {}

// EnterNoneExpr is called when production noneExpr is entered.
func (s *BasePathListener) EnterNoneExpr(ctx *NoneExprContext) {}

// ExitNoneExpr is called when production noneExpr is exited.
func (s *BasePathListener) ExitNoneExpr(ctx *NoneExprContext) {}

// EnterAnyExpr is called when production anyExpr is entered.
func (s *BasePathListener) EnterAnyExpr(ctx *AnyExprContext) {}

// ExitAnyExpr is called when production anyExpr is exited.
func (s *BasePathListener) ExitAnyExpr(ctx *AnyExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BasePathListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BasePathListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BasePathListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BasePathListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterTag is called when production tag is entered.
func (s *BasePathListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasePathListener) ExitTag(ctx *TagContext) {}

// EnterClassifierTagExpr is called when production classifierTagExpr is entered.
func (s *BasePathListener) EnterClassifierTagExpr(ctx *ClassifierTagExprContext) {}

// ExitClassifierTagExpr is called when production classifierTagExpr is exited.
func (s *BasePathListener) ExitClassifierTagExpr(ctx *ClassifierTagExprContext) {}

// EnterNameTagExpr is called when production nameTagExpr is entered.
func (s *BasePathListener) EnterNameTagExpr(ctx *NameTagExprContext) {}

// ExitNameTagExpr is called when production nameTagExpr is exited.
func (s *BasePathListener) ExitNameTagExpr(ctx *NameTagExprContext) {}
