package main

//Import the fmt package for prettier stdout
import (
	"fmt"
)

func main() {
	//Set a loop
	for i := 1; i <= 100; i++ {
		if i%6 == 0 { // Our 2 & 3 division check (because 2 * 3 = 6)
			printer(i, "is divisible by two and three")
		} else if i%3 == 0 { // Our 3 division check
			printer(i, "is divisible by three")
		} else if i%2 == 0 { // Our Even check
			printer(i, "even")
		} else { // Our Odds are left over
			printer(i, "odd")
		}
	}
}

func printer(number int, mod string) {
	fmt.Print("The number '", number, "' is ", mod, ".\n")
}
