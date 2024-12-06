package main

import (
	"errors"
	"fmt"
)

// Env maps identifiers to Values
type Env map[string]Value

func (e Env) Lookup(s string) (Value, error) {
	val, ok := e[s]
	if !ok {
		return nil, fmt.Errorf("Unbound identifier: %s", s)
	}
	return val, nil
}

func (e Env) Extend(bindings map[string]Value) Env {
	newEnv := make(Env)
	for k, v := range e {
		newEnv[k] = v
	}
	for k, v := range bindings {
		newEnv[k] = v
	}
	return newEnv
}

// interp evaluates an expression given an environment
func interp(expr ExprC, env Env) (Value, error) {
	switch e := expr.(type) {
	case NumC:
		return NumV{N: e.N}, nil
	case IdC:
		return env.Lookup(e.S)
	case StringC:
		return StringV{S: e.S}, nil
	case IfC:
		testVal, err := interp(e.Test, env)
		if err != nil {
			return nil, err
		}
		boolVal, ok := testVal.(BoolV)
		if !ok {
			return nil, errors.New("If test did not evaluate to a boolean")
		}
		if boolVal.B {
			return interp(e.Then, env)
		} else {
			return interp(e.Else, env)
		}
	case LamC:
		return CloV{Params: e.Params, Body: e.Body, Env: env}, nil
	case AppC:
		funVal, err := interp(e.Fun, env)
		if err != nil {
			return nil, err
		}
		var argVals []Value
		for _, a := range e.Args {
			av, aerr := interp(a, env)
			if aerr != nil {
				return nil, aerr
			}
			argVals = append(argVals, av)
		}
		switch fv := funVal.(type) {
		case CloV:
			if len(fv.Params) != len(argVals) {
				return nil, fmt.Errorf("Function expected %d arguments, got %d", len(fv.Params), len(argVals))
			}
			bindings := make(map[string]Value)
			for i, p := range fv.Params {
				bindings[p] = argVals[i]
			}
			newEnv := fv.Env.Extend(bindings)
			return interp(fv.Body, newEnv)
		case PrimOpV:
			return fv.Impl(fv.Name, argVals)
		default:
			return nil, errors.New("Attempted to call a non-function value")
		}
	default:
		return nil, fmt.Errorf("Unknown expression type")
	}
}
