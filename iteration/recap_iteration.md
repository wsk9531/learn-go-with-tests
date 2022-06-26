# Iteration Recap
## More TDD Practice

- Red, Green, Blue (Refactor) was in my head the entire time.

- Practice exercise to change the test so the caller can specify how many repetitions meant adding an extra integer parameter.

    - RED: This threw a "have(string, number), want(string)" test failure at first.
    - GREEN: Changed the Repeat function to accommodate this change
    - BLUE: Pretty simple function, refactored away use of the repeatCount constant with change to take as input parameter
    - I did have a look into the strings package and found a repeat function that would be useful to not reinvent the wheel.

## Learned Go's only looping construct - for loops!
    
- C-style loop
    - init      e.g. i := 0
    - condition e.g. i <= 5
    - post/increment e.g. i++
- Result: 
    - for i := 0; i <= 5; i++ {}
    - no parentheses needed, braces always required

- Other variants
    - While loop style:
    - for sum < 1000 {}
    
    - Infinite loop:
    - for {}

- Use of keyword 'continue' is possible to cleanly move to next iteration.

- Good resource: https://gobyexample.com/for

## Learned how to write benchmarks 
    
- Live in _test.go files.
- pass in variable b *testing.B 
    - NOT testing.t - b for benchmark, t for tests
- Runs b.N times and gives a time per operation performed 
    - the Repeat function from this example took ~120 nanoseconds

## Reinforced Example knowledge
    
- Need to print the value from a function call
- This gets compared to //Output: xxxxx comment. 
Comment must be preset for it to run in test suite

## CLI commands
verbose testing flag
- go test -v 

run benchmark flag
- go test -bench=.



