package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Give me an %s ->", "adjective")
	scanner.Scan()
	firstAdjective := scanner.Text()
	fmt.Println(firstAdjective)

	fmt.Printf("Give me an %s ->", "noun")
	scanner.Scan()
	firstNoun := scanner.Text()
	fmt.Println(firstNoun)

	fmt.Printf("Give me an %s ->", "name")
	scanner.Scan()
	firstName := scanner.Text()
	fmt.Println(firstName)

	fmt.Printf("Give me an %s ->", "verb")
	scanner.Scan()
	firstVerb := scanner.Text()
	fmt.Println(firstVerb)

	fmt.Printf("Give me the %s ->", "past tense of previous verb")
	scanner.Scan()
	pastTenseVerb := scanner.Text()
	fmt.Println(pastTenseVerb)

	fmt.Printf("Give me an %s ->", "noun")
	scanner.Scan()
	secondNoun := scanner.Text()
	fmt.Println(secondNoun)

	fmt.Printf("Give me an %s ->", "body part")
	scanner.Scan()
	bodyPart := scanner.Text()
	fmt.Println(bodyPart)

	fmt.Printf("Give me a %s ->", "past tense verb")
	scanner.Scan()
	pastTenseVerb2 := scanner.Text()
	fmt.Println(pastTenseVerb2)

	/*
	   Once upon a time, there was a very {{adjective}} {{noun}}
	    named {{name}}.
	   {{Previous name}} really liked to {{verb}}.
	   {{Previous name}} {{past tense verb}} all day long.
	   She {{past tense verb}} so much that she leveled up and became a 1337 {{noun}}.
	   Then, her {{body part}} fell off from so much {{noun}}, and she {{verb}}.

	   The End.
	*/
}
