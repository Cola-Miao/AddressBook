package main

import (
	"AddressBook/entrance"
	"AddressBook/initialize"
)

func main() {
	initialize.LogInit()
	initialize.DirInit()
	for {
		entrance.Entrance()
	}
}
