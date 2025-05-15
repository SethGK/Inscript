package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type TraceListener struct {
	*antlr.BaseParseTreeListener
}

func NewTraceListener() *TraceListener {
	return &TraceListener{
		BaseParseTreeListener: &antlr.BaseParseTreeListener{},
	}
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("enter %s\n", ctx.GetText())
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("exit %s\n", ctx.GetText())
}
