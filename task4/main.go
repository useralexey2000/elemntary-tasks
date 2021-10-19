package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args) < 3 {
		usage(args[0])
		os.Exit(0)

	} else if len(args) == 3 {

		i, err := countString(args[1], args[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Number of matches is: ", i)
		os.Exit(0)

	} else if len(args) == 4 {
		replaceString(args[1], args[2], args[3])
		os.Exit(0)
	}
}

func countString(f string, s string) (int, error) {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return 0, fmt.Errorf("Cant read file %v", err)
	}
	i := strings.Count(string(bs), s)

	return i, nil
}

func replaceString(f, s1, s2 string) error {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		return fmt.Errorf("Cant read file %v", err)
	}
	s := strings.ReplaceAll(string(bs), s1, s2)
	if err = ioutil.WriteFile(f, []byte(s), 066); err != nil {
		return fmt.Errorf("Cant read file %v", err)
	}
	return nil
}

func usage(n string) {
	fmt.Printf("usage: %v %v\n", n, "filename<string> | filename<string> old<string> new<string>")
}
