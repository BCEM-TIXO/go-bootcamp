package main

import (
	// "encoding/json"
	"bufio"
	"fmt"
	// "lo
	"os"
	"os/exec"
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

func main() {
	args, err := GetArgs()
	if err != nil {
		fmt.Println(err)
	}
	if len(os.Args) < 2 {
		fmt.Println("No command(((")
		return
	}
	comand := os.Args[1]
	if len(os.Args) >= 3 {
		args = append(os.Args[2:], args...)
	}
	cmd := exec.Command(comand, args...)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

}
