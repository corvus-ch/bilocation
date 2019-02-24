package matcher

import (
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
	classifier := ""

	if classifierToken := c.GetClassifier(); classifierToken != nil {
		classifier = classifierToken.GetText()
	}

	l.stack.Push(NewClassifier(classifier))
}

func (l *path) ExitNameTagExpr(c *grammar.NameTagExprContext) {
	name := ""
	if nameToken := c.GetName(); nameToken != nil {
		name = nameToken.GetText()
	}

	if classifierToken := c.GetClassifier(); classifierToken != nil {
		l.stack.Push(NewNameAndClassifier(name, classifierToken.GetText()))
	}

	l.stack.Push(NewName(name))
}
