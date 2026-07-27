package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        first_word := cleanInput(scanner.Text())
		fmt.Println("Your command was:", first_word[0])

        // Check for scanner reading errors
        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", err)
        }
    }
}
