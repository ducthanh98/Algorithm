package main

import (
	"fmt"
	"strings"
)

func canReach(s string, minJump int, maxJump int) bool {
	values := strings.Split(s, "")
	length := len(values)
	possibleJumps := make([]int, 1)
	position := 0
	for i := 1; i < length; i++ {

		if position < len(possibleJumps)-1 && i-possibleJumps[position] > maxJump {
			position++
		}

		if values[i] == "1" || i-possibleJumps[position] < minJump {
			continue
		}

		if i-possibleJumps[len(possibleJumps)-1] > maxJump {
			return false
		}
		possibleJumps = append(possibleJumps, i)
	}

	return possibleJumps[len(possibleJumps)-1] == length-1

}

func main() {
	fmt.Println(canReach("01101110", 2, 3))
}
