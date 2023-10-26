package main

import (
	"fmt"
)

type fileError struct {
	msg string
}

func (e *fileError) Error() string {
	return fmt.Sprintf("File Error: %s", e.msg)
}

type networkError struct {
	msg string
}

func (e *networkError) Error() string {
	return fmt.Sprintf("Network Error: %s", e.msg)
}

type databaseError struct {
	msg string
}

func (e *databaseError) Error() string {
	return fmt.Sprintf("Database Error: %s", e.msg)
}

func faultyOperation(opType string) error {
	switch opType {
	case "file":
		return &fileError{"Unable to open file"}
	case "network":
		return &networkError{"Connection timed out"}
	case "database":
		return &databaseError{"Failed to query the database"}
	default:
		return nil
	}
}

func main() {
	err := faultyOperation("network")

	if err != nil {
		if e, ok := err.(*fileError); ok {
			fmt.Println("Handling file error:", e.Error())
		} else if e, ok := err.(*networkError); ok {
			fmt.Println("Handling network error:", e.Error())
		} else if e, ok := err.(*databaseError); ok {
			fmt.Println("Handling database error:", e.Error())
		} else {
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		fmt.Println("No errors occurred!")
	}
}
