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

## How to create a module in Go.
`go mod init example.com/hello` and the module path name
This creates a Go mod file that tracks your code's dependencies.

By default: Code is executed sequentially, i.e. processing of tasks one after the other. To prevent this blocking behavior, you use GoRoutines to utilize concurrency.
GoRoutines are concurrent functions that run while the rest of the program is executing.

`go` keyword - starts a new thread in Go.
A goroutine is a lightweight thread managed by the Go runtime.

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

### How to save mixed data types in Go

This can be done using a feature called `Struct` - i.e. Structure in Go.

Custom types are created in Go using the `type` keyword with the name you specify.

```
  type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
	isSignedUpForNewsLetter bool
}
```
This creates a type called "UserData" based on the values specified in the structure. Structures give us flexibility over Maps as they allow the use of mixed data types and also define the fields of our Custom types.

Structures are lightweight classes similar to classes in OOP languages like Java, C++, but they do not support inheritance.


## Synchronizing Goroutines
By default, the main goroutine does NOT wait for other goroutines to finish.
```
  var wg = sync.WaitGroup{}

  wg.Add(n) - n is the number of goroutines.
  wg.Wait()
  wg.Done()
```

`WaitGroup` waits for the launched goroutine to finish.
Package `sync` provides basic synchronization functionality.

- `Add` method: sets the number of goroutines to wait for (It increases the counter by the provided number).
- `Wait` method: Blocks until all the added threads are done executing, i.e. the wait counter is zero.
- `Done` method: It decrements the wait counter by one, i.e. this is called by the goroutine to indicate that it has finished executing.

Comparision between goroutine and threads in other languages:
- Writing concurrent code is more complex.
- More overhead in terms of memory and resources.

Threads vs. Goroutines:
- Creating thread is more expensive, and slow to startup.
- Threads are heavyweight and requires more hardware resources.

- Go uses the concept of `Green thread` which is an abstraction of an actual thread.
- Goroutines are managed by the Go runtime, thus we are only interacting with these high level goroutines. Instead of OS Threads that are managed by the OS Kernel, and are hardware dependent.
- Channels is a built-in function used for communication between goroutines, when they have shared data.

Go code is grouped into packages and packages are grouped into modules.