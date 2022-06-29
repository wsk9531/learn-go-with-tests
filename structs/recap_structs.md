## Structs, Methods, Interfaces

### Struct

- a named collection of fields that store data
- used to create types
- access fields of a struct with . e.g. myStruct.field

- Passing in structs to a function conveys intent better than raw values

### Methods

- a function with a receiver.
- "method declaration binds an identifier (method name) to a method, and associates the method with the receiver's base type"

In short, they are similar to functions but get invoked on an instance of a type.

e.g. If I have a Circle type, I can write a method to calculate its area and because it associates with the type, an instantiation
like c Cicle can invoke area like c.Area()

Adding methods so you can add functionality to your data types and so you can implement interfaces

Declaration syntax:
`func (receiverName ReceiverType) MethodName(args).`

- Go convention that the receiverName is the first letter of the Receiver type (e.g. c Circle)

### Interfaces

https://go.dev/ref/spec#Interface_types

- Interfaces define a type set. A variable of interface type can store a value of any type that is in the type set of the interface

#### Implementing and interface (from Go Docs)

A type T implements an interface I if

- T is not an interface and is an element of the type set of I; or
- T is an interface and the type set of T is a subset of the type set of I.
  A value of type T implements an interface if T implements the interface.

What does this mean?
In its most basic form an interface specifies a (possibly empty) list of methods. The type set defined by such an interface is the set of types which implement all of those methods, and the corresponding method set consists exactly of the methods specified by the interface.

It gets way deeper and I should return to the linked page in a couple of days to review, giving some time to play around with them. It's been a while since entering the world of BASIC :(

#### Result:

- We created a helper function that asks to invoked with an entry of type Shape interface. If the thing we try to call the function with doesn't implement the interface (i.e. c
  ontain Shape's method) the code won't compile.

**Interface resolution is implicit**
Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

- Value is: seperation of concerns! Helper method can only care about if the input item is a Shape.

### Table Driven Tests

- Builds a list of test cases
- We declared a slice of structs with fields that matched Shape, named the variables, etc.

- In a for loop over the range of the slice, we invoke a test (t.Run()) for each entry that specified got and want. Super easy to add a new field

```go
        {shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
        {shape: Circle{Radius: 10}, want: 314.1592653589793},
```

### Other

- %g prints out a more precise decimal number than %f, which helps to REALLY make sure math calcs are the same. Familiar computer math precision topic.
- See: https://pkg.go.dev/fmt

- %#v prints a struct with values in it's field, v helpful for debugging and testing
