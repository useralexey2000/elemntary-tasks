package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args) < 3 || len(args) > 4 {
		usage(args[0])
		os.Exit(0)

	}
	// can be separated for reading and writing depending on method called
	f, err := os.OpenFile(args[1], os.O_RDWR, 066)
	defer func() {
		if f != nil {
			f.Close()
		}
		usage(args[0])
	}()
	if err != nil {
		fmt.Printf("cant open file %s, %v\n", args[1], err)
		return
	}

	if len(args) == 3 {

		i, err := CountString(f, args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Number of matches is: ", i)
		return
	}
	if len(args) == 4 {

		if err := ReplaceString(&MyFile{f: f}, args[2], args[3]); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func CountString(in io.Reader, s string) (int, error) {

	bs, err := ioutil.ReadAll(in)
	if err != nil {
		return 0, fmt.Errorf("cant read file %w", err)
	}
	i := strings.Count(string(bs), s)

	return i, nil
}

type MyFile struct {
	f *os.File
}

// overwrite write method for os.file to overwrite file and comply with io.ReadWriter
// no need to test them as they are direct calls to methods in standard library
func (ff *MyFile) Write(b []byte) (int, error) {
	return ff.f.WriteAt(b, 0)
}

func (ff *MyFile) Read(b []byte) (int, error) {
	return ff.f.Read(b)
}

func ReplaceString(rw io.ReadWriter, s1, s2 string) error {

	bs, err := ioutil.ReadAll(rw)

	if err != nil {
		return fmt.Errorf("cant read from source %w", err)
	}

	res := strings.ReplaceAll(string(bs), s1, s2)

	_, err = rw.Write([]byte(res))
	if err != nil {
		return fmt.Errorf("cant read file %w", err)
	}

	return nil

}

func usage(n string) {
	fmt.Printf("usage: %v filename<string> | filename<string> old<string> new<string>\n", n)
}
