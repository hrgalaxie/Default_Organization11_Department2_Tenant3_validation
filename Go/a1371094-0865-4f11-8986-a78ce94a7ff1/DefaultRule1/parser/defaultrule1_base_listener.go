// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650280692057/DefaultRule1.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // DefaultRule1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDefaultRule1Listener is a complete listener for a parse tree produced by DefaultRule1Parser.
type BaseDefaultRule1Listener struct{}

var _ DefaultRule1Listener = &BaseDefaultRule1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDefaultRule1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDefaultRule1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDefaultRule1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDefaultRule1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDefaultRule1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDefaultRule1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDefaultRule1Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDefaultRule1Listener) ExitLabel(ctx *LabelContext) {}
