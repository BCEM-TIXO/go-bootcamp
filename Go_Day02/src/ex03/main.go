package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

func GetArgs() ([]string, error) {
	fileScanner := bufio.NewScanner(os.Stdin)
	var res []string
	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		res = append(res, fileScanner.Text())
	}
	return res, nil
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func tar(file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf(file, ": Is a directory")
	}
	filePath := ""
	if fl.a != "." {
		filePath = fl.a
	} else {
		filePath = filepath.Dir(file)
	}
	file_name := fileNameWithoutExt(info.Name())
	mtime := info.ModTime().Unix()
	file_name = filePath + "/" + file_name + "_" + strconv.FormatInt(mtime, 10) + ".tar.gz"

	cmd := exec.Command("tar", "-czf", file_name, file)
	return cmd.Run()
}

type Flags struct {
	a string
}

var fl Flags

func init() {
	flag.StringVar(&fl.a, "a", ".", "out put dir")
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
			err := tar(file)
			if err != nil {
				fmt.Println(err)
			}
			defer wg.Done()
		}(file)
	}
	wg.Wait()

}

func checkFlags() ([]string, error) {
	args := flag.Args()
	if len(args) == 0 {
		return nil, errors.New("ERROR: missing file path")
	}

	info, err := os.Stat(fl.a)
	if os.IsNotExist(err) || !info.IsDir() {
		return nil, fmt.Errorf("ERROR: \"%s\" dir is not exist or it is a file", fl.a)
	}
	return args, nil
}
