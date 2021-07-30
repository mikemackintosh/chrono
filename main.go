package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	ENV_CHRONO     = "CHRONO_TIMESTAMP"
	ENV_CHRONO_DUR = "CHRONO_DUR"
)

func main() {
	var t = time.Now()
	var startTime = os.Getenv(ENV_CHRONO)

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
