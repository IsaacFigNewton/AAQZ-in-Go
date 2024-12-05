package main

import (
	"aaqz/env"
	"fmt"
)

/*
New types to add:

(println s) → boolean
  s : str

(read-num) → real

(read-str) → str

(seq a b ...) → any
  a : any
  b : any*

(++ a b ...) → string
  a : any
  b : any*

test case:
	{seq
	{println "What is your favorite integer between 6 and 7?"}
	{bind [your-number = {read-num}]
		{println {++ "Interesting, you picked " your-number ". Bold choice!"}}}}
*/

/*
Every function must have:
	A commented header line that expresses the result of the function in terms of its inputs, written in English.
		Be as precise as you can within the space of a line or two.
	A type declaration (possibly inline), specifying the input and output types.
	Test cases. A function without test cases is incomplete. Write the test cases first, please.
*/

func main() {
	/*
		focus on developing:
			a representation of an AST
			the interp function
			the parsing function (if you have time left)
	*/

	/*
		‹expr› ::=
			| ‹num›		NumC
			| ‹id›                           IdC
			| ‹string›                       StrC
			| { if ‹expr› ‹expr› ‹expr› }    IfC()		desugar into AppC during interp
			| { bind ‹clause›* ‹expr› }      AppC(LamC())
			| { (‹id›*) => ‹expr› }          LamC
			| { ‹expr› ‹expr›* }             ?

		<clause> ::=
			{ <id> <expr> }
	*/

	fmt.Println(env.TopEnv)
}
