package main

import "fmt"

func main() {
    var a, b, c int = 10, 20, 30
    fmt.Println("a =", a, "\nb =", b, "\nc =", c)

    var d, e, f, g int // create empty variables for results

    d = a - b * c
    fmt.Println("a - b * c =", d)

    e = (a - b) * c
    fmt.Println("a - (b * c) =", e)

    f = a % b * c
    fmt.Println("a % b * c =", f)

    g = a % (b * c)
    fmt.Println("a % (b * c) =", g)
}

/*
Expected output:

a = 10
b = 20
c = 30
a - b * c = -590
a - (b * c) = -300
a % b * c = 300
a % (b * c) = 10

*/
