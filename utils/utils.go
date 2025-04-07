package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJSON[T any](filePath string, target *[]T) error {
	// lets read the json file data into a slice of bytes
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read data from : %s", filePath)
	}

	// lets unmarshel the content into our target
	if err := json.Unmarshal(content, &target); err != nil {
		return fmt.Errorf("failed to Unmarshal data from : %s,%w", filePath, err)
	}
	return nil
}
