package main

import (
	// "file/filepath"
	"fmt"
	"io/fs"
	"os"
	// "path"
	"errors"
	"flag"
	"path/filepath"
)

func FileWalker(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", p)
	return nil
}

type Flags struct {
	f, d, sl bool
	ext      string
}

var fl Flags

func init() {
	flag.BoolVar(&fl.f, "f", false, "print file")
	flag.BoolVar(&fl.d, "d", false, "print directory")
	flag.BoolVar(&fl.sl, "sl", false, "print symbolic links")
	flag.StringVar(&fl.ext, "ext", "", "print only files with extension")
}

func main() {
	flag.Parse()
	root, err := checkFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath.WalkDir(root, func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if p == root {
			return nil
		}
		if info.Type()&os.ModeSymlink != 0 && fl.sl {
			symlink, _ := filepath.EvalSymlinks(p)
			if _, err := os.Stat(symlink); err == nil {
				fmt.Printf("%s -> %s\n", p, symlink)
			} else {
				fmt.Printf("%s -> %s\n", p, "[broken]")
			}
		}
		if fl.f && info.Type().IsRegular() {
			if fl.ext != "" && filepath.Ext(p) == ("."+fl.ext) {
				fmt.Printf("%s\n", p)
			} else if fl.ext == "" {
				fmt.Printf("%s\n", p)
			}
		}
		if fl.d && info.IsDir() {
			fmt.Printf("%s\n", p)
		}
		return nil
	})
}

func checkFlags() (string, error) {
	args := flag.Args()
	if len(args) == 0 {
		return "", errors.New("ERROR: missing file path")
	}

	if !fl.f && fl.ext != "" {
		return "", errors.New("ERROR: -ext may only be used with -f")
	}

	if !fl.f && !fl.d && !fl.sl {
		fl.f, fl.d, fl.sl = true, true, true
	}

	return args[0], nil
}
