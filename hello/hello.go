package main

import (
	"fmt";
	"rsc.io/quote";
	"log";
	"example.com/greetings";
)

func main()  {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string {"Zhang", "San", "Li", "Si", "Wang", "Wu"}

	// Request greeting messages for the names
	messages, err := greetings.Hellos(names)

	// Request a greeting message.
	message, err := greetings.Hello("Aris")
	if err != nil {
		log.Fatal(err)
	}


	// If no error was returned, print the returned message
	fmt.Println(quote.Go())
	// message := greetings.Hello("Gladys")
	fmt.Println(message)
	
	for i, m := range messages {
		fmt.Println(i, m)
	}

}