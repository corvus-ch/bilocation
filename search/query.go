package search

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/corvus-ch/bilocation/search/internal/grammar"
	"github.com/corvus-ch/bilocation/search/internal/matcher"
	"github.com/corvus-ch/bilocation/tag"
)

type Matcher interface {
	Match(tag.Set) bool
}

type queryMatcher struct {
	matcher matcher.Matcher
}

func NewQuery(query string) Matcher {
	input := antlr.NewInputStream(query)
	lexer := grammar.NewPathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := grammar.NewPathParser(stream)
	stack := matcher.NewStack()
	antlr.ParseTreeWalkerDefault.Walk(matcher.NewPath(stack), p.Query())

	return &queryMatcher{matcher: stack.Pop()}
}

func (q *queryMatcher) Match(set tag.Set) bool {
	return q.matcher.Match(set)
}
