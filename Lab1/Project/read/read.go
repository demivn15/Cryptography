package read

import "fmt"

func ReadInput() string {
	var input string
	fmt.Print("Enter some text: ")
	fmt.Scan(&input)
	return input
}
