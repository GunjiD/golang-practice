package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/tempconv"
)

func main() {
	args := os.Args[1:]
	// argument check.
	if len(args) == 0 {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		input := s.Text()
		l := strings.Split(input, " ")
		for _, arg := range l {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		}
	}
}
