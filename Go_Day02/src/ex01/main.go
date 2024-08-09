package main

import (
	// "file/filepath"
	"fmt"
	// "io/fs"
	"bufio"
	"errors"
	"flag"
	"os"
	"strings"
	"sync"
)

func FileWalker(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", p)
	return nil
}

type Flags struct {
	l, m, w bool
}

var fl Flags

func init() {
	flag.BoolVar(&fl.l, "l", false, "print lines count")
	flag.BoolVar(&fl.m, "m", false, "print symbols count")
	flag.BoolVar(&fl.w, "w", false, "print word count")
}

func wc(file string) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if info.IsDir() {
		fmt.Println(file, ": Is a directory")
		return
	}
	file_open, err := os.Open(file)
	if err != nil {
		fmt.Println(err)

		return
	}
	defer file_open.Close()
	var l, w int
	scanner := bufio.NewScanner(file_open)
	for scanner.Scan() {
		l++
		w += len(strings.Fields(scanner.Text()))
	}
	if fl.l {
		fmt.Printf("%d\t%s\n", l, file)
		return
	}
	if fl.w {
		fmt.Printf("%d\t%s\n", w, file)
	}
	if fl.m {
		fmt.Printf("%d\t%s\n", info.Size(), file)
		return
	}
}

func main() {
	flag.Parse()
	files, err := checkFlags()
	var wg sync.WaitGroup
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			wc(file)
			defer wg.Done()
		}(file)
	}
	wg.Wait()
}

func xor(A, B bool) bool {
	return (A || B) && !(A && B)
}

func checkFlags() ([]string, error) {
	const errTxt = "ERROR: only one flag can be used"
	if fl.l && fl.m && fl.w {
		return nil, errors.New(errTxt)
	}

	if (xor(fl.l, fl.m) && fl.w) || (xor(fl.m, fl.w) && fl.l) {
		return nil, errors.New(errTxt)
	}
	args := flag.Args()
	if len(args) == 0 {
		return nil, errors.New("ERROR: missing file path")
	}

	if !fl.l && !fl.m && !fl.w {
		fl.w = true
	}
	return args, nil
}
