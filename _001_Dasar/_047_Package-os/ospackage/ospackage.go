package ospackage

import (
	"os"
)

func GetArguments() []interface{} {
	args := os.Args // Akan kita beri command seperit ini: go run main.go 123 456 789

	var newArgs []interface{}

	for index, value := range args {
		if index == 0 {
			continue
		}

		newArgs = append(newArgs, value)
	}

	return newArgs
}

func GetHostname() (string, error) {
	return os.Hostname()
}
