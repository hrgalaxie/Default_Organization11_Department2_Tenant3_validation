// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650281957076/Rule4.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Rule4

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRule4Listener is a complete listener for a parse tree produced by Rule4Parser.
type BaseRule4Listener struct{}

var _ Rule4Listener = &BaseRule4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRule4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRule4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRule4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRule4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRule4Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRule4Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseRule4Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseRule4Listener) ExitLabel(ctx *LabelContext) {}
