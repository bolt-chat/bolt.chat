package main

import (
	"encoding/hex"
	"fmt"
	"github.com/boltchat/protocol/v2/encoder"
	"github.com/boltchat/protocol/v2/events"
)

// This is very poorly written, but it's just
// for debugging purposes. ;)
func printResult(res []byte) {
	fmt.Printf("bin: ")
	for _, b := range res {
		fmt.Printf("%08b ", b)
	}

	fmt.Println()

	fmt.Printf("dec: ")
	for _, b := range res {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	fmt.Printf("hex: %s", hex.EncodeToString(res))

	fmt.Println()
}

func main() {
	d := encoder.NewEncoder(nil)
	res := d.Encode(&events.Event{
		Header: &events.Header{
			Version:   1,
			EventType: events.JoinEvent,
			HasCRC:    false,
		},
	})

	printResult(res)
}
