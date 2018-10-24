package main

//Import the fmt package for prettier stdout
import (
	"fmt"
)

//Counter Struct: Setup a struct to hold our state
type Counter struct {
	total int
}

func main() {
	//Create a Sync Channel for concurrent state management
	sync := make(chan *Counter)
	//Set a loop to and create a couple threads
	for i := 1; i <= 100; i++ {
		//create threads that will do the alternating counting
		go worker(1, sync)
		go worker(0, sync)
	}
	// Start the Sync between the two threads with our Struct as the State Manager between the threads
	sync <- new(Counter)
	// Empty the Sync and End Program
	<-sync
}

func worker(thread int, sync chan *Counter) {
	for {
		//Read in the current state value
		counter := <-sync
		//Check to see if we have hit our limit
		if counter.total <= 99 {
			if counter.total%2 == thread { //Our Even/Odd Detector
				//Increment the Counter
				counter.total++
				//Dump to STDOUT
				fmt.Print("Thread ", thread, ": The Number is '", counter.total, "'\n")
			}
		}
		//Update the current state
		sync <- counter
	}
}
