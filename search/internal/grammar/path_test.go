package grammar_test

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bketelsen/logr"
	"github.com/corvus-ch/bilocation/search/internal/grammar"
	"github.com/corvus-ch/logr/buffered"
	"github.com/sebdah/goldie"
)

var queries = map[string][]string{
	"empty": {""},
	"any":   {"*", "**", "*/*", "/*/*", "*/*/", "/*/*/", "**/*", "/**/*", "**/*/", "/**/*/"},

	"a":   {"a", "/a", "/a/", "a/"},
	"a=b": {"a=b", "/a=b", "/a=b/", "a=b/"},

	"and ab":      {"a/b", "/a/b", "/a/b/", "a/b/"},
	"and abc":     {"a/b/c", "/a/b/c", "/a/b/c/", "a/b/c/"},
	"and abcd":    {"a/b/c/d", "/a/b/c/d", "/a/b/c/d/", "a/b/c/d/"},
	"and a=bc":    {"a=b/c", "/a=b/c", "/a=b/c/", "a=b/c/"},
	"and ab=c":    {"a/b=c", "/a/b=c", "/a/b=c/", "a/b=c/"},
	"and abc=d":   {"a/b/c=d", "/a/b/c=d", "/a/b/c=d/", "a/b/c=d/"},
	"and a=bcd":   {"a=b/c/d", "/a=b/c/d", "/a=b/c/d/", "a=b/c/d/"},
	"and ab=ac":   {"a/b=c/d", "/a/b=c/d", "/a/b=c/d/", "a/b=c/d/"},
	"and a or bd": {"a/b,d", "/a/b,d", "/a/b,d/", "a/b,d/"},
	"and a=*b":    {"a=*/b", "/a=*/b", "/a=*/b/", "a=*/b/"},
	"and ab=*":    {"a/b=*", "/a/b=*", "/a/b=*/", "a/b=*/"},
	"and abc=*":   {"a/b/c=*", "/a/b/c=*", "/a/b/c=*/", "a/b/c=*/"},
	"and a=*bc":   {"a=*/b/c", "/a=*/b/c", "/a=*/b/c/", "a=*/b/c/"},
	"and ab=*c":   {"a/b=*/c", "/a/b=*/c", "/a/b=*/c/", "a/b=*/c/"},

	"or ab":       {"a,b", "/a,b", "/a,b/", "a,b/"},
	"or abc":      {"a,b,c", "/a,b,c", "/a,b,c/", "a,b,c/"},
	"or abcd":     {"a,b,c,d", "/a,b,c,d", "/a,b,c,d/", "a,b,c,d/"},
	"or a=bc":     {"a=b,c", "/a=b,c", "/a=b,c/", "a=b,c/"},
	"or ab=c":     {"a,b=c", "/a,b=c", "/a,b=c/", "a,b=c/"},
	"or abc=d":    {"a,b,c=d", "/a,b,c=d", "/a,b,c=d/", "a,b,c=d/"},
	"or a=bcd":    {"a=b,c,d", "/a=b,c,d", "/a=b,c,d/", "a=b,c,d/"},
	"or ab=ac":    {"a,b=c,d", "/a,b=c,d", "/a,b=c,d/", "a,b=c,d/"},
	"or ab and d": {"a,b/d", "/a,b/d", "/a,b/d/", "a,b/d/"},
	"or a=*b":     {"a=*,b", "/a=*,b", "/a=*,b/", "a=*,b/"},
	"or ab=*":     {"a,b=*", "/a,b=*", "/a,b=*/", "a,b=*/"},
	"or abc=*":    {"a,b,c=*", "/a,b,c=*", "/a,b,c=*/", "a,b,c=*/"},
	"or a=*bc":    {"a=*,b,c", "/a=*,b,c", "/a=*,b,c/", "a=*,b,c/"},
	"or a*=cb":    {"a,b=*,c", "/a,b=*,c", "/a,b=*,c/", "a,b=*,c/"},
}

type TestListener struct {
	*grammar.BasePathListener
	*antlr.DefaultErrorListener

	log    logr.Logger
	t      *testing.T
	indent string
}

func (tl *TestListener) EnterQuery(c *grammar.QueryContext) {
	tl.log.Infof("query:")
}

func (tl *TestListener) EnterAnyExpr(c *grammar.AnyExprContext) {
	tl.indent += "\t"
	tl.log.Infof("%sANY", tl.indent)
}

func (tl *TestListener) EnterAndExpr(c *grammar.AndExprContext) {
	tl.indent += "\t"
	tl.log.Infof("%sAND: %s", tl.indent, c.GetText())
}

func (tl *TestListener) EnterOrExpr(c *grammar.OrExprContext) {
	tl.indent += "\t"
	tl.log.Infof("%sOR: %s", tl.indent, c.GetText())
}

func (tl *TestListener) EnterClassifierTagExpr(c *grammar.ClassifierTagExprContext) {
	tl.indent += "\t"
	classifier := ""
	if classifierToken := c.GetClassifier(); classifierToken != nil {
		classifier = classifierToken.GetText()
	}
	tl.log.Infof("%sclassifier tag: classifier(%s)\n", tl.indent, classifier)
}

func (tl *TestListener) EnterNameTagExpr(c *grammar.NameTagExprContext) {
	tl.indent += "\t"
	name := ""
	classifier := ""
	op := ""
	if nameToken := c.GetName(); nameToken != nil {
		name = nameToken.GetText()
	}
	if classifierToken := c.GetClassifier(); classifierToken != nil {
		classifier = classifierToken.GetText()
	}
	if opToken := c.GetOp(); opToken != nil {
		op = opToken.GetText()
	}
	tl.log.Infof("%sname tag: name(%s) classifier(%s) op(%s)\n", tl.indent, name, classifier, op)
}

func (tl *TestListener) ExitAndExpr(c *grammar.AndExprContext) {
	tl.indent = tl.indent[1:]
}

func (tl *TestListener) ExitOrExpr(c *grammar.OrExprContext) {
	tl.indent = tl.indent[1:]
}

func (tl *TestListener) ExitTag(c *grammar.TagContext) {
	tl.indent = tl.indent[1:]
}

func (tl *TestListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	tl.t.Fatalf("line %d:%d %s", line, column, msg)
}

func TestGrammar(t *testing.T) {
	for name, queries := range queries {
		t.Run(name, func(t *testing.T) {
			for _, query := range queries {
				t.Run(query, func(t *testing.T) {
					log := buffered.New(0)
					listener := new(TestListener)
					listener.log = log
					listener.t = t
					input := antlr.NewInputStream(query)
					lexer := grammar.NewPathLexer(input)
					stream := antlr.NewCommonTokenStream(lexer, 0)
					p := grammar.NewPathParser(stream)
					p.AddErrorListener(listener)
					antlr.ParseTreeWalkerDefault.Walk(listener, p.Query())
					goldie.Assert(t, name, log.Buf().Bytes())
				})
			}
		})
	}
}
