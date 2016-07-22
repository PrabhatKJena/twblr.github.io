package control_structures

import (
    "fmt"
)

func fizzBuzz(input int32) string {
	if input % 3 == 0 && input % 5 == 0 {
		return "FizzBuzz"
	} else if input % 5 == 0 {
		return "Buzz";
	} else if input % 3 == 0 {
		return "Fizz";
	}
	return fmt.Sprintf("%v",input);
}