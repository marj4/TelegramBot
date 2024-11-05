package e

import (
	"fmt"
)

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	} else {
		return fmt.Errorf("%s:%w", msg, err)
	}
}
