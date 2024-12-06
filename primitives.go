package main

import (
	"errors"
	"fmt"
)

// Function type for primops
type PrimOpImpl func(name string, args []Value) (Value, error)

// Validates and extracts numeric arguments for primops
func numOpCheck(args []Value, opName string, expected int) ([]float64, error) {
	if len(args) != expected {
		return nil, fmt.Errorf("%s expects %d arguments, got %d", opName, expected, len(args))
	}
	nums := make([]float64, expected)
	for i, a := range args {
		nv, ok := a.(NumV)
		if !ok {
			return nil, fmt.Errorf("%s expects numeric arguments", opName)
		}
		nums[i] = nv.N
	}
	return nums, nil
}

// Addition (+)
func plusPrimOp(name string, args []Value) (Value, error) {
	nums, err := numOpCheck(args, "+", 2)
	if err != nil {
		return nil, err
	}
	return NumV{N: nums[0] + nums[1]}, nil
}

// Subtraction (-)
func minusPrimOp(name string, args []Value) (Value, error) {
	nums, err := numOpCheck(args, "-", 2)
	if err != nil {
		return nil, err
	}
	return NumV{N: nums[0] - nums[1]}, nil
}

// Multiplication (*)
func timesPrimOp(name string, args []Value) (Value, error) {
	nums, err := numOpCheck(args, "*", 2)
	if err != nil {
		return nil, err
	}
	return NumV{N: nums[0] * nums[1]}, nil
}

// Division (/)
func divPrimOp(name string, args []Value) (Value, error) {
	nums, err := numOpCheck(args, "/", 2)
	if err != nil {
		return nil, err
	}
	if nums[1] == 0 {
		return nil, errors.New("division by zero")
	}
	return NumV{N: nums[0] / nums[1]}, nil
}

// <= primop
func leqPrimOp(name string, args []Value) (Value, error) {
	nums, err := numOpCheck(args, "<=", 2)
	if err != nil {
		return nil, err
	}
	return BoolV{B: nums[0] <= nums[1]}, nil
}

// equal? primop
func equalPrimOp(name string, args []Value) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("equal? expects 2 arguments, got %d", len(args))
	}
	switch a1 := args[0].(type) {
	case NumV:
		if a2, ok := args[1].(NumV); ok {
			return BoolV{B: a1.N == a2.N}, nil
		}
	case BoolV:
		if a2, ok := args[1].(BoolV); ok {
			return BoolV{B: a1.B == a2.B}, nil
		}
	case StringV:
		if a2, ok := args[1].(StringV); ok {
			return BoolV{B: a1.S == a2.S}, nil
		}
	}
	return BoolV{B: false}, nil
}

// Define the top-level environment
func topEnv() Env {
	return Env{
		"+":      PrimOpV{"+", plusPrimOp},
		"-":      PrimOpV{"-", minusPrimOp},
		"*":      PrimOpV{"*", timesPrimOp},
		"/":      PrimOpV{"/", divPrimOp},
		"<=":     PrimOpV{"<=", leqPrimOp},
		"equal?": PrimOpV{"equal?", equalPrimOp},
		"true":   BoolV{true},
		"false":  BoolV{false},
	}
}
