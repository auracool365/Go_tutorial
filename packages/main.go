package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

/* Go Packages are a way to organize and reuse code in Go. A package is a collection of related Go source files that are compiled together.
Each package has a unique name, and the files within a package can access each other's functions, variables, and types. Packages can be then
imported into other packages to use their functionality.
*/

/* The main package is a special package in Go that serves as the entry point for the execution of a Go program. The main package must have
a function called main(), which is the starting point of the program. When running a Go program, the Go runtime looks for the main package
and executes the main() function within it.*/

/* The main package is typically used for creating executable programs, while other packages are used for creating libraries that can be
imported and used by other programs. The main package can also import other packages to use their functionality.*/


func main() {
	// 1. fmt package: used for formatted I/O operations.
	fmt.Println("Hello, World!")
	fmt.Printf("The value of pi is approximately %.2f\n", 3.14159)
	fmt.Println("The sum of 5 and 3 is:", 5+3)
	fmt.Println()

	// 2. math package: provides basic constants and mathematical functions.
	fmt.Println("The square root of 16 is:", math.Sqrt(16))
	fmt.Println("The value of pi is:", math.Pi)
	fmt.Println("The sine of 30 degrees is:", math.Sin(30*math.Pi/180))
	fmt.Println("The cosine of 60 degrees is:", math.Cos(60*math.Pi/180))
	fmt.Println("The tangent of 45 degrees is:", math.Tan(45*math.Pi/180))
	fmt.Println("The natural logarithm of 2 is:", math.Log(2))
	fmt.Println("The exponential of 1 is:", math.Exp(1))
	fmt.Println("The absolute value of -5 is:", math.Abs(-5))
	fmt.Println("The maximum of 10 and 20 is:", math.Max(10, 20))
	fmt.Println("The minimum of 10 and 20 is:", math.Min(10, 20))
	fmt.Println()

	// 3. strings package: provides functions for manipulating strings.
	str := "Hello, World!"
	fmt.Println("The length of the string is:", len(str))
	fmt.Println("The string in uppercase is:", strings.ToUpper(str))
	fmt.Println("The string in lowercase is:", strings.ToLower(str))
	fmt.Println()

	// 4. time package: provides functionality for measuring and displaying time.
	timeVar := time.Now()
	fmt.Println("The current time is:", timeVar.Format("2006-01-02 15:04:05"))
	fmt.Println("The current year is:", timeVar.Year())
	fmt.Println("The current month is:", timeVar.Month())
	fmt.Println("The current day is:", timeVar.Day())
	fmt.Println("The current hour is:", timeVar.Hour())
	fmt.Println("The current minute is:", timeVar.Minute())
	fmt.Println("The current second is:", timeVar.Second())
	fmt.Println()

	// 5. net/http package: build a tiny web server and make a request to it.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Hello from the HTTP server")
	})
	go func() {
		_ = http.ListenAndServe(":8080", mux)
	}()

	var resp *http.Response
	var err error
	for range 10 {
		resp, err = http.Get("http://127.0.0.1:8080/")
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("HTTP response:", string(body))
	} else {
		fmt.Println("HTTP example error:", err)
	}
	fmt.Println()

	// 6. io package: read and write data streams.
	reader := strings.NewReader("Learn Go")
	data, _ := io.ReadAll(reader)
	fmt.Println("IO example:", string(data))
	fmt.Println()

	// 7. os package: interact with the filesystem.
	if err := os.WriteFile("hello.txt", []byte("Hello from os in packages folder"), 0o644); err != nil {
		fmt.Println("Write file error:", err)
	} else {
		fileData, readErr := os.ReadFile("hello.txt")
		if readErr == nil {
			fmt.Println("File contents:", string(fileData))
		}
	}
	fmt.Println()

	// 8. encoding/json package: encode and decode JSON.
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "Ada", Age: 36}
	fmt.Println("Original struct:", person)

	encoded, _ := json.Marshal(person)
	fmt.Println("JSON version:", string(encoded))

	var decoded Person
	_ = json.Unmarshal(encoded, &decoded)
	fmt.Println("Decoded JSON:", decoded.Name, decoded.Age)
	fmt.Println()

	// 9. database/sql package: create a generic DB handle.
	db := &sql.DB{}
	fmt.Println("Database handle created:", db != nil)

	// 10. context package: manage cancellation and deadlines.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	fmt.Println("Context example created:", ctx != nil)

	// 11. crypto/rand package: generate cryptographically secure random numbers.
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)
	fmt.Println("Random bytes:", string(randomBytes))

	// 12. math/rand package: generate pseudo-random numbers.
	fmt.Println("Random integer between 0 and 99:", rand.Intn(100))
	fmt.Println("Random float between 0.0 and 1.0:", rand.Float64())
}
