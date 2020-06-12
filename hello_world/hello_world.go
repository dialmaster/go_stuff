package main

import "fmt"

func main() {
    a := 15.0;
    fmt.Printf("1: a is %f\n", a);
    add1 := func(inpNum float64) float64 {
        fmt.Printf("add1 called!\n")
        a := 14.0
        fmt.Printf("2: a is %f\n", a);

        return inpNum + 1
    }

    applyToVariable := func(fun func(float64) float64, myVar float64) float64 {
        fmt.Printf("applyToVariable called!\n")
        return fun(myVar)
    }

    b := applyToVariable(add1, a)
    fmt.Printf("3: a is %f\n", a)
    fmt.Printf("hello, world\n")

    fmt.Printf("Value is: %f", b)
}
