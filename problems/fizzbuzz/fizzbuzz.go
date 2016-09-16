package main

func main() {

	for i := 1; i <= 100; i++ {

		out := ""

		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}

		//
		if out != "" {
			print(out)
		} else {
			print(i)
		}

		println(" ")

	}

}
