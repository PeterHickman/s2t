package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var full bool

func divmod(a, b int64) (int64, int64) {
	remainder := a % b
	quotient := (a - remainder) / b
	return quotient, remainder
}

func human_duration(duration_in_seconds int64) string {
	var d, h, m, s int64

	s = duration_in_seconds

	m, s = divmod(s, 60)
	h, m = divmod(m, 60)
	d, h = divmod(h, 24)

	if full {
		return fmt.Sprintf("%dd %dh %dm %ds", d, h, m, s)
	} else {
		if d == 0 {
			if h == 0 {
				if m == 0 {
					return fmt.Sprintf("%ds", s)
				} else if s == 0 {
					return fmt.Sprintf("%dm", m)
				} else {
					return fmt.Sprintf("%dm %ds", m, s)
				}
			} else if m == 0 {
				return fmt.Sprintf("%dh", h)
			} else {
				return fmt.Sprintf("%dh %dm", h, m)
			}
		} else if h == 0 {
			return fmt.Sprintf("%dd", d)
		} else {
			return fmt.Sprintf("%dd %dh", d, h)
		}
	}
}

func init() {
	f := flag.Bool("full", false, "Display the unabbreviated time")

	flag.Parse()

	full = *f

	if len(flag.Args()) != 1 {
		fmt.Println("No arguments supplied")
		os.Exit(1)
	}
}

func main() {
	seconds, err := strconv.ParseInt(flag.Arg(0), 10, 64)

	if err != nil {
		fmt.Printf("Unknown number [%s]\n", flag.Arg(0))
		os.Exit(1)
	}

	println(human_duration(seconds))
}
