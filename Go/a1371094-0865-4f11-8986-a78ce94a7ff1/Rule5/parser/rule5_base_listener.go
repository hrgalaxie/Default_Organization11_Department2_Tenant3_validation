// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650282387322/Rule5.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Rule5

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule5Listener is a complete listener for a parse tree produced by Rule5Parser.
type BaseRule5Listener struct{}

var _ Rule5Listener = &BaseRule5Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule5Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule5Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule5Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule5Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule5Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule5Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseRule5Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseRule5Listener) ExitLabel(ctx *LabelContext) {}
