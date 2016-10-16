package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me an adjective, please. ")
	firstAdjective, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the adjective has occured.")
	}
	fmt.Println(firstAdjective)
	fmt.Print("Give me a noun, please. ")
	firstNoun, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the noun has occured.")
	}
	fmt.Println(firstNoun)
	fmt.Print("Give me a name, please.")
	firstName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the name has occured.")
	}
	fmt.Println(firstName)
	fmt.Print("Give me a verb, please.")
	firstVerb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the verb has occured.")
	}
	fmt.Println(firstVerb)
	fmt.Print("Give me the past tense of the previous verb, please.")
	pastTenseVerb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the past tense of the previous verb has occured.")
	}
	fmt.Println(pastTenseVerb)
	fmt.Print("Give me a second noun, please.")
	secondNoun, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the second noun has occured.")
	}
	fmt.Println(secondNoun)
	fmt.Print("Give me a body part, please.")
	bodyPart, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the body part has occured.")
	}
	fmt.Print(bodyPart)
	fmt.Print("Give me a second past tense verb, please.")
	secondVerb, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error with the second verb has occured.")
	}
	fmt.Print(secondVerb)
	fmt.Printf("Once upon a time, there was a very %s %s named %s.\n %s really liked to %s. %s %s all day long.\n' She %s so much that she leveled up and became a 1337 %s.\n Then, her %s fell off from so much use, and she %s.\n", firstAdjective, firstNoun, firstName, firstName, firstVerb, firstName, pastTenseVerb, pastTenseVerb, secondNoun, bodyPart, secondVerb)
}

/*
Once upon a time, there was a very {{adjective}} {{noun}}
 named {{name}}.
{{Previous name}} really liked to {{verb}}.
{{Previous name}} {{past tense verb}} all day long.
She {{past tense verb}} so much that she leveled up and became a 1337 {{noun}}.
Then, her {{body part}} fell off from so much {{noun}}, and she {{verb}}.

The End.
*/
