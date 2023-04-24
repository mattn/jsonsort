package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/oliveagle/jsonpath"
)

func run() int {
	var jp string

	flag.StringVar(&jp, "p", "$.id", "jsonpath to the value for sort column")
	flag.Parse()

	dec := json.NewDecoder(os.Stdin)

	vv := []any{}
	for {
		var v any
		err := dec.Decode(&v)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			return 1
		}
		vv = append(vv, v)
	}

	sort.Slice(vv, func(i, j int) bool {
		v1, err := jsonpath.JsonPathLookup(vv[i], jp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		v2, err := jsonpath.JsonPathLookup(vv[j], jp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		return fmt.Sprint(v1) < fmt.Sprint(v2)
	})
	enc := json.NewEncoder(os.Stdout)
	for _, v := range vv {
		enc.Encode(v)
	}
	return 0
}

func main() {
	os.Exit(run())
}
