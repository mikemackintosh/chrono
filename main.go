package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var flagMilli bool

const (
	ENV_CHRONO     = "CHRONO_TIMESTAMP"
	ENV_CHRONO_DUR = "CHRONO_DUR"
)

func init() {
	flag.BoolVar(&flagMilli, "m", false, "Show only current milliseconds")
}

func main() {
	flag.Parse()

	var t = time.Now()
	var startTime = os.Getenv(ENV_CHRONO)

	if flagMilli {
		fmt.Println(t.UnixNano() / 1000000)
		os.Exit(0)
	}

	if len(startTime) == 0 {
		fmt.Printf("unset %s;\n", ENV_CHRONO_DUR)
		fmt.Printf("export %s=%v;\n", ENV_CHRONO, t.UnixNano())
		os.Exit(0)
	}

	var duration time.Duration
	i, _ := strconv.Atoi(startTime)
	if i > 0 {
		duration = time.Since(time.Unix(0, int64(i)))
	}

	fmt.Printf("export %s=%v;\n", ENV_CHRONO_DUR, duration)
	fmt.Printf("unset %s;\n", ENV_CHRONO)
}
