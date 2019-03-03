package matcher

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/corvus-ch/bilocation/search/internal/grammar"
)

type path struct {
	grammar.BasePathListener
	stack *Stack
}

// NewPath returns path listener which builds up a matcher stack.
func NewPath(stack *Stack) grammar.PathListener {
	l := new(path)
	l.stack = stack

	return l
}

func (l *path) ExitQuery(c *grammar.QueryContext) {
	if l.stack.Depth() <= 0 {
		l.stack.Push(NewNil())
	}
}

func (l *path) ExitNoneExpr(c *grammar.NoneExprContext) {
	l.stack.Drain()
	l.stack.Push(NewNone())
}

func (l *path) ExitAnyExpr(c *grammar.AnyExprContext) {
	l.stack.Drain()
	l.stack.Push(NewAny())
}

func (l *path) ExitAndExpr(c *grammar.AndExprContext) {
	l.stack.Push(NewAnd(l.stack.Drain()))
}

func (l *path) ExitOrExpr(c *grammar.OrExprContext) {
	l.stack.Push(NewOr(l.stack.Drain()))
}

func (l *path) ExitClassifierTagExpr(c *grammar.ClassifierTagExprContext) {
	m := NewClassifier(getClassifierText(c))

	if op := getOp(c); op == "!=" {
		m = NewNot(m)
	}

	l.stack.Push(m)
}

func (l *path) ExitNameTagExpr(c *grammar.NameTagExprContext) {
	var m Matcher

	if classifier := getClassifierText(c); len(classifier) > 0 {
		m = NewNameAndClassifier(getName(c), classifier)
	} else {
		m = NewName(getName(c))
	}

	if op := getOp(c); op == "!" || op == "!=" {
		m = NewNot(m)
	}

	l.stack.Push(m)
}

type nameContext interface {
	GetName() antlr.Token
}

func getName(ctx nameContext) string {
	if t := ctx.GetName(); t != nil {
		return t.GetText()
	}

	return ""
}

type classifierContext interface {
	GetClassifier() antlr.Token
}

func getClassifierText(ctx classifierContext) string {
	if t := ctx.GetClassifier(); t != nil {
		return t.GetText()
	}

	return ""
}

type opContext interface {
	GetOp() antlr.Token
}

func getOp(ctx opContext) string {
	if t := ctx.GetOp(); t != nil {
		return t.GetText()
	}

	return ""
}
