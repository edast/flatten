package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type jsonVal struct {
	keys []string
	value string
}

func main() {
	flag.Parse()
	args := flag.Args()

	var input []io.Reader
	if len(args) == 0 {
		// read from stdin if no args are specified
		input = append(input, os.Stdin)
		err := parseInput(os.Stdin)
		if err != nil {
			fmt.Errorf("failed parsing input: %w", err)
			os.Exit(1)
		}
		return
	}

	for _, filename := range args {
		r, err := os.Open(filename)
		if err != nil {
			fmt.Errorf("failed opening file %s: %w", filename, err)
			os.Exit(1)
		}
		err = parseInput(r)
		if err != nil {
			fmt.Errorf("failed parsing input: %w", err)
			os.Exit(1)
		}
	}
}

func parseInput(r io.Reader) error {
	var buf []byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buf = append(buf, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	var e interface{}
	if err := json.Unmarshal(buf, &e); err != nil {
		return err
	}

	list := json2dot([]string{}, e)
	for _, v := range list {
		fmt.Printf("%s = %s\n", strings.Join(v.keys, "."), v.value)
	}
	return nil
}

func json2dot(keys []string, d interface{}) []jsonVal {
	var result []jsonVal

	m, ok := d.(map[string]interface{})
	if !ok {
		return append(result, jsonVal{
			keys: keys,
			value: fmt.Sprintf("%v", d),
		})
	}
	for k, v := range m {
		ks := append(keys, k)
		switch v.(type) {
		case map[string]interface{}:
			result = append(result, json2dot(ks, v)...)
		case []interface{}:
			for _, val := range v.([]interface{}) {
				result = append(result, json2dot(ks, val)...)
			}
		default:
			tmpKeys := make([]string, len(ks))
			// make value copy of array not to be overwritten with last value in the loop
			copy(tmpKeys, ks)

			result = append(result, jsonVal{keys: tmpKeys, value: fmt.Sprintf("%v", v)})
		}
	}
	return result
}