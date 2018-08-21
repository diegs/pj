// Binary pj reads JSON from stdin, round-trips it through the decoder, and pretty-prints it to stdout.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func roundtripJSON(in, out *os.File) error {
	raw := make(map[string]interface{})
	d := json.NewDecoder(bufio.NewReader(in))
	d.UseNumber()
	if err := d.Decode(&raw); err != nil {
		return fmt.Errorf("read json: %v", err)
	}

	w := bufio.NewWriter(out)
	defer w.Flush()
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	if err := e.Encode(&raw); err != nil {
		return fmt.Errorf("write json: %v", err)
	}

	return nil
}

func main() {
	if err := roundtripJSON(os.Stdin, os.Stdout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
