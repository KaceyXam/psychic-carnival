package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	mdfile, err := os.ReadFile("test-files/hello-world.md")
	if err != nil {
		panic(err)
	}

	parsedmd := ParseMarkdown(mdfile)

	for _, item := range parsedmd {
		fmt.Println(item)
	}

}

func ParseMarkdown(md []byte) []string {
	res := strings.Split(WindowsSucks(string(md)), "\n")
	return DeleteEmpty(res)
}

func WindowsSucks(file string) string {
	return strings.ReplaceAll(string(file), "\r\n", "\n")
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}