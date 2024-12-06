package main

import (
	"fmt"
	"os"
)

// Tests an expression against an expected result
func runTest(expr ExprC, env Env, expected string) {
	val, err := interp(expr, env)
	if err != nil {
		fmt.Printf("Test failed (Expression: %T): error %v\n", expr, err)
		os.Exit(1)
	}
	got := val.Serialize()
	if got != expected {
		fmt.Printf("Test failed.\nExpression: %#v\nExpected: %s\nGot: %s\n", expr, expected, got)
		os.Exit(1)
	} else {
		fmt.Printf("Test passed: %T -> %s\n", expr, got)
	}
}

func main() {
	env := topEnv()

	// Basic tests
	runTest(NumC{42}, env, "42")
	runTest(IdC{"true"}, env, "true")

	// Test addition
	runTest(AppC{Fun: IdC{"+"}, Args: []ExprC{NumC{1}, NumC{2}}}, env, "3")

	// If test
	runTest(IfC{Test: IdC{"true"}, Then: NumC{1}, Else: NumC{0}}, env, "1")
	runTest(IfC{Test: IdC{"false"}, Then: NumC{1}, Else: NumC{0}}, env, "0")

	// Primops
	runTest(AppC{Fun: IdC{"-"}, Args: []ExprC{NumC{5}, NumC{3}}}, env, "2")
	runTest(AppC{Fun: IdC{"*"}, Args: []ExprC{NumC{4}, NumC{2}}}, env, "8")
	runTest(AppC{Fun: IdC{"/"}, Args: []ExprC{NumC{6}, NumC{3}}}, env, "2")
	runTest(AppC{Fun: IdC{"<="}, Args: []ExprC{NumC{3}, NumC{4}}}, env, "true")
	runTest(AppC{Fun: IdC{"<="}, Args: []ExprC{NumC{5}, NumC{5}}}, env, "true")
	runTest(AppC{Fun: IdC{"<="}, Args: []ExprC{NumC{6}, NumC{5}}}, env, "false")

	// equal?
	runTest(AppC{Fun: IdC{"equal?"}, Args: []ExprC{NumC{10}, NumC{10}}}, env, "true")
	runTest(AppC{Fun: IdC{"equal?"}, Args: []ExprC{IdC{"true"}, IdC{"true"}}}, env, "true")
	runTest(AppC{Fun: IdC{"equal?"}, Args: []ExprC{StringC{"hello"}, StringC{"hello"}}}, env, "true")
	runTest(AppC{Fun: IdC{"equal?"}, Args: []ExprC{StringC{"hello"}, StringC{"world"}}}, env, "false")

	// Lambda application: ((lambda (x) x) 5) = 5
	lambdaExpr := LamC{
		Params: []string{"x"},
		Body:   IdC{"x"},
	}
	runTest(AppC{Fun: lambdaExpr, Args: []ExprC{NumC{5}}}, env, "5")

	// Nested lambda test: (((lambda (x) (lambda (y) (+ x y))) 3) 4) = 7
	nestedLambda := AppC{
		Fun: AppC{
			Fun: LamC{Params: []string{"x"},
				Body: LamC{Params: []string{"y"},
					Body: AppC{Fun: IdC{"+"}, Args: []ExprC{IdC{"x"}, IdC{"y"}}}}},
			Args: []ExprC{NumC{3}},
		},
		Args: []ExprC{NumC{4}},
	}
	runTest(nestedLambda, env, "7")

	fmt.Println("All tests passed.")
}
