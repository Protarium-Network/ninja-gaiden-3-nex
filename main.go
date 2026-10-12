package main

import (
	"sync"

	"github.com/Protarium-Network/ninja-gaiden-3-nex/nex"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		nex.StartAuthenticationServer()
	}()
	go func() {
		defer wg.Done()
		nex.StartSecureServer()
	}()

	wg.Wait()
}
