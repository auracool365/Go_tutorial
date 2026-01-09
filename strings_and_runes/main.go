package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. String
	// Strings in Go are utf-8 encoded. Unlike in other languages where strings are a sequence of characters,
	// in Go, strings are a sequence of bytes (a read-only slice[] of bytes) and are immutable, which means
	// that operations like concatenation create a new string. The character that make up the string are called Runes
	// Since strings are a sequence of bytes, it means that indexing into strings outputs bytes not characters(runes) as in some other languages.
	// The byte can be broken down to binary to fully understand.
	// I still don't really get it, but I'm working on it

	myText := "Hello"
	fmt.Println(myText)
	index := myText[0]
	fmt.Println(index)
	fmt.Printf("%d %c %T \n", index, index, index) // byte (also Unicode number in ASCII), character, type

	text1 := "déal"
	for index, element := range text1 {
		fmt.Println(index, element, string(element)) // index 2 was not printed
	}
	// Unlike utf-32 (32 bits for all characters, wasting space), utf-8 allows variable-length encoding (automatically
	// adjusting to accommodate characters longer than 8 bits). The number of bytes required for a character in the
	// string depends on its Unicode code point:
	//   - 0 to 127 = 1 byte
	//   - 128 to 2,047 = 2 bytes
	//   - 2,048 to 65,535 = 3 bytes
	//   - 65,536 to 1,114,111 = 4 bytes
	//
	// For the binary layout:
	//   - 0xxxxxxx = 1 byte
	//   - 110xxxxx 10xxxxxx = 2 bytes
	//   - 1110xxxx 10xxxxxx 10xxxxxx = 3 bytes
	//   - 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx = 4 bytes
	//
	// So in the "déal" example, the character 'é' has Unicode code point 233, which falls between 128 and 2047,
	// so it uses 2 bytes in UTF-8. The binary representation of 233 is 11101001,
	// For 2-byte UTF-8 encoding:
	// First 5 bits go in 110xxxxx
	// Remaining 6 bits go in 10xxxxxx
	// So 233 → 00011101001 (11 bits total) and in UTF-8 it's encoded as:
	//   - First byte: 11000011
	//   - Second byte: 10101001
	// This explains why Go's range loop sees 'é' as one character, but indexing the string by bytes sees only parts of it.
	//
	// Breakdown:
	// Character      | Unicode(byte values) | Bytes (Binary)     | Byte indexes
	// ----------     | -------              | ------------------ | -------------
	// d              | 100                  | 01100100           | 0
	// é              | 233                  | 11000011 10101001  | 1, 2
	// a              | 97                   | 01100001           | 3
	// l              | 108                  | 01101100           | 4
	//
	// Hence, the range loop prints indices: 0, 1, 3, 4 (skipping index 2 because it's the continuation byte of 'é')

	// Note:
	// When I access text_1[1], I am indexing into the raw bytes of the UTF-8 encoded string.
	// Since the character 'é' is stored using two bytes (11000011 and 10101001),
	// text_1[1] returns just the second byte: 10101001, which is 169 in decimal.
	// This byte alone does not represent a valid character, so printing it with %c may produce gibberish or �.
	// It also shows why direct indexing into strings doesn't work well for multibyte characters.
	a := text1[1]
	fmt.Println(a)
	fmt.Printf("%d %c %T \n", a, a, a) // prints: byte value of the first byte (as int), character, and the type (uint8)
	fmt.Printf("%x\n", []byte(text1))  // see exact bytes in hex
	fmt.Printf("%08b\n", text1[1])     // returns the first byte for é (%08b = 8 bits, %16b = 16 bits)
	fmt.Printf("%08b\n", text1[2])     // returns the second byte for é

	// Summary: Strings in Go are values whose underlying data is a sequence of bytes.
	// Therefore, checking the length of a string returns the number of bytes, not the number of characters (runes).
	fmt.Printf("The length of %v is %d\n", myText, len(myText))
	fmt.Printf("The length of %v is %d\n", text1, len(text1)) // 'é' contributes an extra byte (2 bytes total)

	string1 := []string{"H", "e", "l", "l", "o"}
	fmt.Println(string1)

	var string2 = ""
	for i := range string1 {
		string2 += string1[i]
	}
	fmt.Println(string2)

	fmt.Println(string1[0])

	// string_2[0] = "h" // Error due to the immutability of strings in Go

	// The above copying of string1 into string2 is not efficient as it creates a new copy each time due to string immutability
	// To fix this "strings" can be imported to make use of str.Builder(). string3 is efficient

	strBuilder := strings.Builder{}
	for i := range string1 {
		strBuilder.WriteString(string1[i])
	}

	string3 := strBuilder.String()
	fmt.Println(string3)

	// 2. Runes (char)
	var distinction rune = 'A'
	fmt.Println(distinction)                                         // decimal ascii code
	fmt.Printf("%d %c %T \n", distinction, distinction, distinction) // decimal ascii code, character, data type of the ascii

	// More examples
	var price string = "$48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}

}
