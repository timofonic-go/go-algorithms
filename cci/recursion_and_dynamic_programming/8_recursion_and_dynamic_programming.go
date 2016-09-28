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
