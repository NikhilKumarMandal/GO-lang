package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Our pizza app"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza bewteen 1 to 5:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// Add 1 to the rating and display the result
		fmt.Println("Add 1 to your rating:", numRating+1)
	}
	
}