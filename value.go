package main

import (
	"fmt"
	"strconv"
)

// Value is the interface for all runtime values
type Value interface {
	isValue()
	Serialize() string
}

type NumV struct {
	N float64
}

func (NumV) isValue() {}
func (v NumV) Serialize() string {
	return fmt.Sprintf("%g", v.N)
}

type BoolV struct {
	B bool
}

func (BoolV) isValue() {}
func (v BoolV) Serialize() string {
	if v.B {
		return "true"
	}
	return "false"
}

type StringV struct {
	S string
}

func (StringV) isValue() {}
func (v StringV) Serialize() string {
	return strconv.Quote(v.S)
}

type CloV struct {
	Params []string
	Body   ExprC
	Env    Env
}

func (CloV) isValue() {}
func (v CloV) Serialize() string {
	return "#<procedure>"
}

type PrimOpV struct {
	Name string
	Impl PrimOpImpl
}

func (PrimOpV) isValue() {}
func (v PrimOpV) Serialize() string {
	return "#<primop>"
}

type VoidV struct{}

func (VoidV) isValue() {}
func (v VoidV) Serialize() string {
	return "#<void>"
}
