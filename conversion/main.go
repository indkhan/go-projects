package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to our pizza app damn it feels smooth")
	fmt.Println("Please rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	println("thanks for rating us ", input)

}
