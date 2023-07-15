Links:
### Go official documentation
- https://pkg.go.dev/strconv


Go was created and designed to run on multiple cores and it allows the maximum utilization of growing system computing powers, complex infrastructures, and multi-core processors by enabling massive multithreading, concurrency, and performance under-pressure. Which is why most of the modern infrastructures in cloud computing are built using Go, e.g. Kubernetes, Docker etc.

A pointer is a variable that points to the memory location of another variable that references the actual value.
Example:
`availableTickets := "Timothy"` - this holds the memory location of the actual value "Timothy".
`&availableTickets` - this holds the hash value of the memory location of the `availableTickets` variable.

There are two types of lists in Go: `array` which has a fixed size and `slices` has a dynamic size.

In Go, `:=` is a shorthand notation for declaring and initializing variables. It is called the short variable declaration syntax.
Example: `name := "John`;
Without type inference: `name string = "John"`
Here the type of name is inferred from the assigned value.

Blank identifiers are used to ignore a variable you don't want to use.
`_` - is used to specify blank identifiers.

In Go, you have to make unused variables explicit.

GoRoutines are concurrent functions that run while the rest of the program is executing.

```
package main
import (
  "fmt"
  "time"
)
func hello() {
  fmt.Println("Hello world goroutine")
}
func main() {
  go hello()
  time.Sleep(1 * time.Second)
  fmt.Println("main function")
}
```

the `main` function fires off a go-routine using the `go` keyword.

Package level variables are variables that are global in a package. It is accessible to every function in the package and are defined at the top level of the package outside all functions.
They can only be declared using `var or const` keywords.

There are three (3) levels of scope:
- Local
  - Declarations within functions.
  - Declarations within blocks of code (e.g. for, if-else).
- Package
  - Declarations outside all functions and can be used everywhere within the same package.
- Global
  - Global variables which are declarations outside all functions & consist of uppercase first letter, and can be used everywhere across all packages.

Variables and functions defined outside any function can be accessed in all other files within the same package.

To explicitly export a function in Go, you capitalize the first character in the function name.

### Creating a MAP data structure in Go
This is done using the `make` and `map` keywords, and ensuring that the keys and values are of same type as you have explicitly defined.
```
  var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
```

### Creating an empty slice of strings
`var mySlice = []string{}`

### Creating an empty slice of maps
`var mySliceOfMaps = make([]map[string]string)`
When creating a slice with `make` keyword, you have to explicitly declare the size of the slice.