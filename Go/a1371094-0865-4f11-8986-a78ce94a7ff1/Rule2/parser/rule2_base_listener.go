// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650281172361/Rule2.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type BaseRule2Listener struct{}

var _ Rule2Listener = &BaseRule2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseRule2Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseRule2Listener) ExitLabel(ctx *LabelContext) {}
