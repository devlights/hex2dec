package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Do は、指定された16進数文字列を10進数に変換します。
func Do(v string) (int64, error) {
	if len(v) == 0 {
		return 0, nil
	}

	if strings.HasPrefix(v, "0x") {
		v = strings.Replace(v, "0x", "", 1)
	}

	return strconv.ParseInt(v, 16, 32)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: hex2dec hex-value(e.g: 0xFFFF)")
		os.Exit(1)
	}

	v := args[0]
	i, err := Do(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\t-->\t%d\n", v, i)
	os.Exit(0)
}
