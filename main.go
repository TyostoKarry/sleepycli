package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const version = "0.1.0"

func main() {
	var (
		wakeFlag      string
		sleepFlag     string
		fromFlag      string
		toFlag        string
		bufferFlag    int
		cyclesMaxFlag int
		cyclesMinFlag int
		versionFlag   bool
	)

	pflag.StringVarP(&wakeFlag, "wake", "w", "", "Calculate bedtimes from wake time (HH:MM)")
	pflag.StringVarP(&sleepFlag, "sleep", "s", "", "Calculate wake times from sleep time (HH:MM)")
	pflag.StringVarP(&fromFlag, "from", "f", "", "Window sleep time (HH:MM), use with --to")
	pflag.StringVarP(&toFlag, "to", "t", "", "Window wake time (HH:MM), use with --from")
	pflag.IntVarP(&bufferFlag, "buffer", "b", 15, "Fall asleep buffer in minutes")
	pflag.IntVarP(&cyclesMinFlag, "cycles-min", "n", 4, "Minimum cycles to show")
	pflag.IntVarP(&cyclesMaxFlag, "cycles-max", "x", 6, "Maximum cycles to show")
	pflag.BoolVarP(&versionFlag, "version", "v", false, "Print version")

	pflag.Parse()

	if versionFlag {
		fmt.Println("sleepycli v" + version)
		os.Exit(0)
	}

	if err := validateAndSelectMode(wakeFlag, sleepFlag, fromFlag, toFlag, bufferFlag, cyclesMinFlag, cyclesMaxFlag); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
