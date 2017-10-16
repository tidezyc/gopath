package main

import (
	"fmt"
	"os"
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
	paths := strings.Split(os.Getenv("GOPATH"), ":")
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
