package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var mainJson string

func main() {
	pwd, _ := os.Getwd()
	arr := make([]string, 0, 10)
	flag.StringVar(&mainJson, "mainJson", "", "-mainJson=中国.json")
	if mainJson == "" {
		mainJson = "中国.json"
	}

	path := filepath.Join(pwd, "geo")
	buf, err := ioutil.ReadFile(filepath.Join(path, mainJson))
	if err != nil {
		fmt.Println(err)
		return
	}

	filepath.Walk(path, func(filename string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(filename, "json") {
			return nil
		}

		if strings.HasSuffix(filename, mainJson) {
			return nil
		}

		fmt.Println(filename)
		c, _ := ioutil.ReadFile(filename)
		arr = append(arr, string(c))

		return nil
	})

	head, tail := Split(string(buf))
	for _, s := range arr {
		head = head + "," + GetConent(s)
	}

	merge := head + tail

	ioutil.WriteFile(filepath.Join(pwd, "merge.json"), []byte(merge), os.ModePerm)
}

func GetConent(s string) string {
	res := strings.Split(s, `"features":[`)
	return strings.TrimSuffix(res[1], "]}")
}

func Split(s string) (head, tail string) {
	return strings.TrimSuffix(s, "]}"), "]}"
}
