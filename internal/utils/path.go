package utils

import (
	"fmt"
	"os"
)

func CreatePath(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	return nil
}
