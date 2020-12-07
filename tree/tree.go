package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	VERSION = "1.0.0"
)

func walk(path, prefix string, isAll bool) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Tree Walking Errot: %v", err)
		os.Exit(1)
	}
	for i, info := range infos {
		var sa []string
		var sb []string
		if i+1 == len(infos) {
			sa = append(sa, "└")
			sb = append(sb, " ")
		}
		if info.IsDir() {
			fmt.Print(prefix)
			sa = append(sa, "├")
			sb = append(sb, "│")
			fmt.Printf("%s── %s\n", sa[0], info.Name())
			walk(filepath.Join(path, info.Name()),  prefix + sb[0] + "   ", isAll)
		} else if isAll {
			fmt.Print(prefix)
			sa = append(sa, "│")
			fmt.Printf("%s── %s\n", sa[0], info.Name())
		}
	}
}

func Parse(path string, isAll bool) {
	isAbs := filepath.IsAbs(path)
	if !isAbs {
		dir, _ := os.Getwd()
		path = filepath.Join(dir, path)
	}
	_, rn := filepath.Split(path)
	fmt.Println(rn)
	walk(path, "", isAll)
}