// Code generated from /usr/local/lib/Go/a1371094-0865-4f11-8986-a78ce94a7ff1/1650281172361/Rule2.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Rule2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule2Listener is a complete listener for a parse tree produced by Rule2Parser.
type Rule2Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
