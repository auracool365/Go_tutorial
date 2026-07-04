package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Multi-error container
// Store multiple concurrent errors
type MultiError struct {
	Errors []error
}

// Return formatted multi-error message
func (m *MultiError) Error() string {

	var builder strings.Builder

	builder.WriteString("multiple errors occurred:\n")

	for i, err := range m.Errors {

		fmt.Fprintf(&builder, "%d. %v\n", i+1, err)
	}

	return builder.String()
}

// Process single item
// Simulate processing an item
func processItem(item string) error {

	// Simulate workload
	time.Sleep(500 * time.Millisecond)

	// Simulate invalid data
	if strings.Contains(item, "bad") {
		return errors.New("invalid item detected")
	}

	// Simulate empty item
	if item == "" {
		return errors.New("empty item")
	}

	return nil
}

// Concurrent processing
// Process multiple items concurrently
func processItems(items []string) error {

	// Error collection channel
	errs := make(chan error, len(items))

	var wg sync.WaitGroup

	// Start worker goroutines
	for _, item := range items {

		wg.Add(1)

		go func(item string) {

			defer wg.Done()

			err := processItem(item)

			if err != nil {

				errs <- fmt.Errorf(
					"processing '%s' failed: %w",
					item,
					err,
				)
			}

		}(item)
	}

	// Wait for all goroutines
	go func() {
		wg.Wait()
		close(errs)
	}()

	// Collect errors
	var errList []error

	for err := range errs {
		errList = append(errList, err)
	}

	// Return grouped errors
	if len(errList) > 0 {
		return &MultiError{
			Errors: errList,
		}
	}

	return nil
}

// usage
func output(items []string) {
	err := processItems(items)

	if err != nil {

		fmt.Println("Concurrent processing errors:")
		fmt.Println(err)

		var multiErr *MultiError

		if errors.As(err, &multiErr) {

			fmt.Println("Detailed error breakdown:")
			for i, e := range multiErr.Errors {
				fmt.Printf("%d. %v\n", i+1, e)
			}
		}

		return
	}

	fmt.Println("\nAll items processed successfully")
}

func main() {

	items1 := []string{ // Used for the errors
		"apple",
		"banana",
		"bad-item",
		"",
		"orange",
		"bad-data",
	}
	items2 := []string{ // Success
		"apple",
		"banana",
		"orange",
	}
	output(items1)
	output(items2)
}
