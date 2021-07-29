package gpl

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func Dup(path []string) {
	counts := make(map[int]string)
	
	for _, filename := range path {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}

		for num, line := range strings.Split(string(data), "\n") {
			if line != "" {
				counts[num] = line
			}
		}
	}

	for n, line := range counts {
		fmt.Println(n, line)
	}
}