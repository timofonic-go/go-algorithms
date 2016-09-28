package recursion_and_dynamic_programming

// Find nth fibonacci number
/**
	I'll be using a top-to-bottom approach with memorization from the 8th chapter.

	EXAMPLE
	idx	1  2  3  4  5  6  7
	numbers	0, 1, 2, 3, 5, 8, 13

	Let's say I need 4th number (should return 5):


						fib(5)

				fib(4)					fib(3)

	        fib(3)		   	  fib(2)      		fib(2)		fib(1)
	fib(2)	       fib(1)	    fib(1)     fib(0)	   fib(1)   fib(0)
      fib(1)  fib(0)


*/
func FindFibonacciTopDown(i int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	m := make(map[int]int)

	return fib(i-1, m) + fib(i-2, m)
}

func fib(i int, m map[int]int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	_, ok := m[i]
	if !ok {
		m[i] = fib(i-1, m) + fib(i-2, m)
	}

	return m[i]
}

// 3.1 Triple Step
/**
EXAMPLE

How many steps hopping:
| 1 2 3
1 steps = 3: 1,1,1
2 steps = 0: _
3 steps = 1: 1
Total: 4

| 1 2 3 4
1 steps = 4: 1,1,1,1
2 steps = 2: 2, 2
3 steps = 0: _
Also: 1,3
Total: 7

| 1 2 3 4 5
1 steps = 5: 1,1,1,1,1
2 steps = 2: _
3 steps = 0: _
But also:
1,1,3
1,3,1
3,3,1
2,2,1
2,1,2
1,2,2
2,3
3,2
Total: 13


BRUTE-FORCE
Just loop through the whole array and sum up all routes.
*/

func CountWays(n int) int {
	if n < 0 {
		// Number of hops does not match the length
		return 0
	} else if n == 0 {
		// Number of hops allows to "land" on the last step
		return 1
	} else {
		// Have not reached the final step yet, continue hopping.
		return CountWays(n-1) + CountWays(n-2) + CountWays(n-3)
	}
}

