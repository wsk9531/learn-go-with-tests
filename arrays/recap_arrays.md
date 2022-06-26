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

**SLICES** are used for this reason, they can have any size and do not encode the size of the collection.

In writing the Sum function, I could either create a new function to handle slice inputs, or break the existing function by changing the argument from an array to a slice.

I did the latter and had to update the first test to pass in a slice.
Once compiling, there were two tests - one that tested 5 values, the other that tested 3. Since both were now using slices and could handle a collection of any size, I removed a test. This got rid of the redundancy.

## CLI commands
Demonstrate code coverage with:
- go test -cover
