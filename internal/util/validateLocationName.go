package util

import (
	"errors"
	"fmt"
)

func ValidateLocationName(locationName string, fixedName []string) error {
	for _, name := range fixedName {
		if name == locationName {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("lokasi: %s tidak valid", locationName))
}
