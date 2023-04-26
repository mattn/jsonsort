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

func fatalIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func run() error {
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
			return err
		}
		vv = append(vv, v)
	}

	sort.Slice(vv, func(i, j int) bool {
		v1, err := jsonpath.JsonPathLookup(vv[i], jp)
		fatalIf(err)
		v2, err := jsonpath.JsonPathLookup(vv[j], jp)
		fatalIf(err)
		return fmt.Sprint(v1) < fmt.Sprint(v2)
	})
	enc := json.NewEncoder(os.Stdout)
	for _, v := range vv {
		enc.Encode(v)
	}
	return nil
}

func main() {
	fatalIf(run())
}
