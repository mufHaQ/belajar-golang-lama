package main

import "fmt"

func main() {
	fmt.Println("===Anonymous Function===")

	// ================================================================================
	// Anonymous Function:

	userBlacklist := func(name string) bool {
		if name == "root" {
			return true
		} else {
			return false
		}
	}

	registerUser("Ulhaq", userBlacklist)
	registerUser("root", func(name string) bool { // Mirip seperti callback pada JavaScript
		if name == "root" {
			return true
		} else {
			return false
		}
	})
	// ================================================================================
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
