package main

import (
	"scancomp/mocking"
	"scancomp/scanning"
)

func main() {
	go scanning.Scan()
	mocking.MockServerRunOn80PortPLS()
}
