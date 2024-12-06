package main

// ExprC is the interface for all AST nodes
type ExprC interface {
	isExprC()
}

type NumC struct {
	N float64
}

func (NumC) isExprC() {}

type IdC struct {
	S string
}

func (IdC) isExprC() {}

type LamC struct {
	Params []string
	Body   ExprC
}

func (LamC) isExprC() {}

type AppC struct {
	Fun  ExprC
	Args []ExprC
}

func (AppC) isExprC() {}

type IfC struct {
	Test ExprC
	Then ExprC
	Else ExprC
}

func (IfC) isExprC() {}

type StringC struct {
	S string
}

func (StringC) isExprC() {}
