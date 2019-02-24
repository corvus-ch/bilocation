// Code generated from search/internal/grammar/Path.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 29, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 5, 3, 19, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6,
	6, 26, 10, 6, 13, 6, 14, 6, 27, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	3, 2, 3, 5, 2, 46, 46, 49, 49, 63, 63, 2, 30, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13,
	3, 2, 2, 2, 5, 18, 3, 2, 2, 2, 7, 20, 3, 2, 2, 2, 9, 22, 3, 2, 2, 2, 11,
	25, 3, 2, 2, 2, 13, 14, 7, 49, 2, 2, 14, 4, 3, 2, 2, 2, 15, 19, 7, 44,
	2, 2, 16, 17, 7, 44, 2, 2, 17, 19, 7, 44, 2, 2, 18, 15, 3, 2, 2, 2, 18,
	16, 3, 2, 2, 2, 19, 6, 3, 2, 2, 2, 20, 21, 7, 63, 2, 2, 21, 8, 3, 2, 2,
	2, 22, 23, 7, 46, 2, 2, 23, 10, 3, 2, 2, 2, 24, 26, 10, 2, 2, 2, 25, 24,
	3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2,
	28, 12, 3, 2, 2, 2, 5, 2, 18, 27, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'/'", "", "'='", "','",
}

var lexerSymbolicNames = []string{
	"", "AND", "ANY", "EQ", "OR", "STRING",
}

var lexerRuleNames = []string{
	"AND", "ANY", "EQ", "OR", "STRING",
}

type PathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPathLexer(input antlr.CharStream) *PathLexer {

	l := new(PathLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Path.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PathLexer tokens.
const (
	PathLexerAND    = 1
	PathLexerANY    = 2
	PathLexerEQ     = 3
	PathLexerOR     = 4
	PathLexerSTRING = 5
)
