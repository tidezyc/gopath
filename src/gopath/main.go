package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	gopath := ""
	if err == nil && wd != "" {
		for wd != "/" {
			info, err := os.Stat(wd + "/src")
			if err == nil && info.IsDir() {
				gopath = wd
				break
			}
			wd = filepath.Dir(wd)
		}
	}
	oldPath := os.Getenv("GOPATH")
	if oldPath == "" {
		env, err := exec.Command("go", "env").Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(bytes.NewReader(env))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "GOPATH=") {
				oldPath = strings.Trim(line[7:], "\"")
				break
			}
		}
	}
	var paths []string
	if oldPath != "" {
		if strings.Contains(oldPath, ":") {
			paths = strings.Split(oldPath, ":")
		} else {
			paths = append(paths, oldPath)
		}
	}
	if gopath != "" {
		exist := false
		for _, v := range paths {
			if v == gopath {
				exist = true
				break
			}
		}
		if !exist {
			paths = append(paths, gopath)
		}
	}
	fmt.Println(strings.Join(paths, ":"))
}
