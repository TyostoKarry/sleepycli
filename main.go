package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/TyostoKarry/sleepycli/internal/goodnight"
	"github.com/TyostoKarry/sleepycli/internal/help"
	"github.com/TyostoKarry/sleepycli/internal/tui"
	"github.com/spf13/pflag"
)

const version = "1.0.0"

func main() {
	help.SetupCustomHelp()

	var (
		nowFlag       bool
		wakeFlag      string
		sleepFlag     string
		fromFlag      string
		toFlag        string
		bufferFlag    int
		cyclesMaxFlag int
		cyclesMinFlag int
		goodNightFlag bool
		versionFlag   bool
	)

	pflag.BoolVarP(&nowFlag, "now", "n", false, "Calculate wake times from current time")
	pflag.StringVarP(&wakeFlag, "wake", "w", "", "Calculate bedtimes from wake time (HH:MM)")
	pflag.StringVarP(&sleepFlag, "sleep", "s", "", "Calculate wake times from sleep time (HH:MM)")
	pflag.StringVarP(&fromFlag, "from", "f", "", "Window sleep time (HH:MM), use with --to")
	pflag.StringVarP(&toFlag, "to", "t", "", "Window wake time (HH:MM), use with --from")
	pflag.IntVarP(&bufferFlag, "buffer", "b", 15, "Fall asleep buffer in minutes")
	pflag.IntVarP(&cyclesMinFlag, "cycles-min", "m", 4, "Minimum cycles to show")
	pflag.IntVarP(&cyclesMaxFlag, "cycles-max", "x", 6, "Maximum cycles to show")
	pflag.BoolVarP(&goodNightFlag, "good-night", "g", false, "Display a random good night art")
	pflag.BoolVarP(&versionFlag, "version", "v", false, "Print version")

	pflag.Parse()

	if !anyFlagSet() {
		p := tea.NewProgram(tui.InitialModel())
		m, err := p.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error running TUI:", err)
			os.Exit(1)
		}
		if model, ok := m.(tui.Model); ok && model.PrintResult != "" {
			fmt.Println(model.PrintResult)
		}
		return
	}

	if versionFlag {
		fmt.Println("sleepycli v" + version)
		os.Exit(0)
	}

	if goodNightFlag {
		fmt.Println(goodnight.RandomGoodNightArt())
		os.Exit(0)
	}

	if err := validateAndSelectMode(nowFlag, wakeFlag, sleepFlag, fromFlag, toFlag, bufferFlag, cyclesMinFlag, cyclesMaxFlag); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		fmt.Fprintln(os.Stderr, "Run 'sleepycli --help' for usage.")
		os.Exit(1)
	}
}

func anyFlagSet() bool {
	found := false
	pflag.Visit(func(f *pflag.Flag) {
		found = true
	})
	return found
}
