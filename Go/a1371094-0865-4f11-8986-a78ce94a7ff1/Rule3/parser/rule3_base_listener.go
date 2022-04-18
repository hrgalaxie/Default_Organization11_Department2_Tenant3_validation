// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650281463057/Rule3.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Rule3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule3Listener is a complete listener for a parse tree produced by Rule3Parser.
type BaseRule3Listener struct{}

var _ Rule3Listener = &BaseRule3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule3Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule3Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseRule3Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseRule3Listener) ExitLabel(ctx *LabelContext) {}
