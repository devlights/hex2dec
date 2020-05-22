package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/devlights/gomy/convert"
)

// Convert -- 指定された文字列を16進数文字列とし10進数文字列へ変換します.
func Convert(val string) (string, error) {
	return convert.Hex2Dec(val, "", 0)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: hex2dec hex-value(e.g: 0xFFFF)")
		os.Exit(1)
	}

	v := args[0]
	i, err := Convert(v)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\t-->\t%s\n", v, i)
	os.Exit(0)
}
