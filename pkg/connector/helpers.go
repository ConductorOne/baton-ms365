package connector

import "fmt"

func wrapError(err error, message string) error {
	return fmt.Errorf("ms365-connector: %s: %w", message, err)
}
