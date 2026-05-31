package main

import (
	"fmt"
	"strings"
)

func main() {

	// Basic String
	name := "Atib"

	fmt.Println(name)

	// Length
	fmt.Println("Length:", len(name))

	// Concatenation
	greeting := "Hello " + name

	fmt.Println(greeting)

	// Strings Are Immutable
	str := "AB"

	// str[0] = 'g' // Compile Error

	fmt.Println(str)

	// Bytes
	fmt.Println("\nBytes:")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c -> %d\n", str[i], str[i])
	}

	for index, r := range str {
		fmt.Printf("index=%d rune=%c codepoint=%d\n",
			index,
			r,
			r)
	}

	// Unicode Example
	chinese := "你好"

	fmt.Println("\nUnicode Example")

	fmt.Println("String:", chinese)
	fmt.Println("len():", len(chinese))

	// Rune Iteration
	fmt.Println("\nRune Iteration")

	for index, r := range chinese {
		fmt.Printf("index=%d rune=%c codepoint=%d\n",
			index,
			r,
			r)
	}

	// strings package
	fmt.Println("\nStrings Package")

	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Contains(name, "ti"))
	fmt.Println(strings.ReplaceAll(name, "Atib", "John"))

	// String Builder examples
	var builder strings.Builder
	builder.WriteString("Hello ")
	builder.WriteString("Go")
	fmt.Println(builder.String())
}
