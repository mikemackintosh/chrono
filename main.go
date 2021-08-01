package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	flagMilli   bool
	flagVersion bool
)

func init() {
	flag.BoolVar(&flagMilli, "m", false, "Show current milliseconds")
	flag.BoolVar(&flagVersion, "v", false, "Show version")
}

func main() {
	var duration time.Duration

	flag.Parse()

	if flagVersion {
		fmt.Printf("1.0")
		os.Exit(0)
	}

	var t = time.Now().UnixMilli()
	if flagMilli {
		fmt.Println(t)
		os.Exit(0)
	}

	var startTime = os.Getenv("START_TIME")
	if len(startTime) == 0 {
		fmt.Printf("0ms")
		os.Exit(0)
	}
	i, _ := strconv.Atoi(startTime)

	if i > 0 {
		duration = time.Since(time.UnixMilli(int64(i)))
	}

	fmt.Printf("%s", humanizeDuration(duration))
}

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65
func humanizeDuration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))
	milliseconds := int64(math.Mod(float64(duration.Milliseconds()), 1000))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"d", days},
		{"h", hours},
		{"m", minutes},
		{"s", seconds},
		{"ms", milliseconds},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		default:
			parts = append(parts, fmt.Sprintf("%d%s", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}
