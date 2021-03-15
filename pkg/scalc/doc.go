// Package scalc provides a set calculator that parses and evaluates given expression.
//
// Grammar of the calculator is the following:
//    expression := "[" operator sets "]"
//    sets := set | set sets
//    set := file | expression
//    operator := "EQ" | "LE" | "GR"
//
// Each file must contain sorted integers, one integer in a line.
//
// Meaning of operators:
//    EQ - returns a set of integers which consists only from values which exists
// in exactly N sets - arguments of operator
//    LE - returns a set of integers which consists only from values which exists
// in less then N sets - arguments of operator
//    GR - returns a set of integers which consists only from values which exists
// in more then N sets - arguments of operator
package scalc
