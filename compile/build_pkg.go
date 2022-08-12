package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// const hello = `
// package main
// import "fmt"
// var a int = 1
// var b int = 2

// func main() {
// 	if a+b >0 {fmt.Println("hello")}
// }
// `

const hello2 = `
package main

import "fmt"
	
func add(a int, b int) int {
	return a + b
}

func main() {
	a := 1
	b := 2
	c := a + b
	d := add(c, 3)
	fmt.Println(d)
}
`

const hello = `
package main

var a int
var b int

func main() {
    a = 4
    b = 7
    c := a + b
    _ = c
}
`

const add = `
package main

func add(x int, y int) int{
	var z int = 2 * y
	return x + y + z
}

var a int
func main(){
	a = 45
	_ = add(a, 13)
}
`

func main() {
	// Replace interface{} with any for this test.
	// ssa.SetNormalizeAnyForTesting(true)
	// defer ssa.SetNormalizeAnyForTesting(false)
	// Parse the source files.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
	if err != nil {
		fmt.Print(err) // parse error
		return
	}
	files := []*ast.File{f}

	// Create the type-checker's package.
	pkg := types.NewPackage("hello", "")

	// Type-check the package, load dependencies.
	// Create and build the SSA program.
	hello, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Print(err) // type error in some package
		return
	}

	// Print out the package.
	hello.WriteTo(os.Stdout)

	// Print out the package-level functions.
	filename := "./1.btc"
	var btcfile *os.File
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		btcfile, _ = os.Create(filename)
	} else {
		btcfile, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	}

	hello.Func("init").WriteTo(os.Stdout)
	hello.Func("main").WriteTo(os.Stdout)

	hello.WriteTo(btcfile)
	hello.Func("init").WriteTo(btcfile)
	hello.Func("main").WriteTo(btcfile)

	// hello.Func("add").WriteTo(os.Stdout)
}
