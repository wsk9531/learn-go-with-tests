# Arrays Recap

Arrays have a fixed capacity, which is defined when at variable declaration:
Either by: 
- [N]type{value1, value2, ..., valueN}
    - e.g. numbers := [5]]int{1, 2, 3, 4, 5}

or 
- [...]type{value1, value2, ..., valueN}
    - e.g. numbers := [...]int{1, 2, 3, 4, 5}

Printing arrays:
- %v placeholder in formatted strings is the "default" and works well for printing arrays

Access a value:
- Use array[index] syntax
---
## Range and iteration:

- Use **range** keyword to iterate over an array.
- Range returns the index and the value.
- Values can be ignored using the blank identifier -> **_**
    - e.g. ignoring the index value from range in situations you only care about the value.

---

Array sizes are encoded into their type.

[4]int passed to a function expecting [5]int will not compile.

**SLICES** are used for this reason, they have a FIXED size but  and do not encode the size of the collection.

We can create new ones from old ones using append()!

Slices look like []type e.g. []int is a slice of integers.
Can use other types. Can even use arrays/slice type slices 
e.g. [][]string

In writing the Sum function, I could either create a new function to handle slice inputs, or break the existing function by changing the argument from an array to a slice.

I did the latter and had to update the first test to pass in a slice.
Once compiling, there were two tests - one that tested 5 values, the other that tested 3. Since both were now using slices and could handle a collection of any size, I removed a test. This got rid of the redundancy.

We can also create slices using **make** keyword. 
- Two input parameters:
    1. The type (e.g. []int to make a slice of integers)
    2. An integer to specify starting capacity

Indexed like arrays - mySlice[N], value reassignment possible with =

## CLI commands
Demonstrate code coverage with:
- go test -cover

## Variadic functions
good resource: https://gobyexample.com/variadic-functions
- Get called with any number of trailing arguments.
e.g. sum(nums ...int) will take any number of integers. nums' equivalent type is []int, and methods like len(nums), range, will all work
- If args are in a slice already, pass them into a variadic function like so:
    - nums := []int{1, 2, 3, 4}
    - sum(nums...)

## On Slice comparison
Slices can only be compared to nil.
We therefore cannot use equality operators (like got, want).
reflect.DeepEqual() is suggested as an alternative. 
However, it is not type safe!!!

## Red, Green, Blue
RED: Add test for SumAll - function that takes in any number of slices
GREEN: Add code to call Sum() for each slice, add each result to a slice called sums
BLUE: Refactor to worry less about capacity, thanks to the append function, which takes a slice and new value, returns a NEW slice with all the items in it. Refactored TestSumAllTails to remove duplicate equality checks.

## Helpful links from the LGWT chapter:
https://blog.golang.org/go-slices-usage-and-internals

https://go.dev/play/p/bTrRmYfNYCp 
Covers slicing an array and changing how slice effects original array, but copies do not affect the original.
https://play.golang.org/p/Poth8JS28sc
Why it's good to make a copy of a slice after slicing a large slice