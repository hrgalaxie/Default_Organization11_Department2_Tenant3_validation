// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650281463057/Rule3.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 44, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 7,
	2, 15, 10, 2, 12, 2, 14, 2, 18, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 6, 5, 39, 10, 5, 13, 5, 14, 5, 40, 3, 5, 3, 5, 2, 2, 6, 3, 3,
	5, 4, 7, 5, 9, 6, 3, 2, 4, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 45, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 19, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2,
	9, 38, 3, 2, 2, 2, 11, 12, 5, 7, 4, 2, 12, 16, 5, 5, 3, 2, 13, 15, 9, 2,
	2, 2, 14, 13, 3, 2, 2, 2, 15, 18, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17,
	3, 2, 2, 2, 17, 4, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 19, 20, 7, 38, 2, 2,
	20, 21, 7, 128, 2, 2, 21, 22, 7, 38, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24, 7,
	34, 2, 2, 24, 25, 7, 35, 2, 2, 25, 26, 7, 38, 2, 2, 26, 27, 7, 128, 2,
	2, 27, 28, 7, 38, 2, 2, 28, 29, 7, 37, 2, 2, 29, 30, 7, 39, 2, 2, 30, 31,
	7, 37, 2, 2, 31, 32, 7, 40, 2, 2, 32, 33, 7, 38, 2, 2, 33, 34, 7, 40, 2,
	2, 34, 35, 7, 35, 2, 2, 35, 36, 7, 34, 2, 2, 36, 8, 3, 2, 2, 2, 37, 39,
	9, 3, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 8, 5, 2, 2, 43, 10, 3,
	2, 2, 2, 5, 2, 16, 40, 3, 8, 2, 2,
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
	"", "", "'$~$'", "' !$~$#%#&$&! '",
}

var lexerSymbolicNames = []string{
	"", "LABEL", "OWNKEY", "SPLITKEY", "WS",
}

var lexerRuleNames = []string{
	"LABEL", "OWNKEY", "SPLITKEY", "WS",
}

type Rule3Lexer struct {
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

func NewRule3Lexer(input antlr.CharStream) *Rule3Lexer {

	l := new(Rule3Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rule3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Rule3Lexer tokens.
const (
	Rule3LexerLABEL    = 1
	Rule3LexerOWNKEY   = 2
	Rule3LexerSPLITKEY = 3
	Rule3LexerWS       = 4
)
