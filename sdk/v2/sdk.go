package sdk

import "fmt"

func Create(debug bool) string {
	return fmt.Sprintf("sdk 2.0.0 (debug: %v)", debug)
}
