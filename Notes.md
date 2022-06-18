# My notes

- create a new go module
```
go mod init <modulepath>
```

- testing
```
go test
go test -v
```


- benchmarking
```
got test -bench="."
```

A struct is just a named collection of fields where you can store data. 

A method is a function with a receiver.

Interfaces allow you to create functions that can be used with different types and create highly-decoupled code whilst still maintaining type safety. 

In Go interface resolution is implicit. If the type you pass in to function matches what interface is asking for, it will compile. 

Table driven tests - useful when you want to build a list of test cases that can be tested in a same manner. 

In Go, when you call a function or method, arguments are copied.
We get the pointer (memory address) of something by placing an & character at the beginning of the symbol.

These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.

In Go, if you want to indicate an error it is idiomatic for your function to return an err for the caller to check and act on.

 Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

errors.New creates a new error with a message of your choosing.

The var keyword allows us to define values global to the package.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.
Never initialize map like this:
```var m map[string]string```
Instead:
```
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```


# bufio
scanner := bufio.NewScanner(filename)
scanner.Scan()
lineFromFile := scanner.Text()

# strings
lineWithoutPrefix := strings.TrimPrefix(line, prefix)

