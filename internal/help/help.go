package help

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

func SetupCustomHelp() {
	fs := pflag.CommandLine
	fs.SortFlags = false

	pflag.Usage = func() {
		name := filepath.Base(os.Args[0])
		out := fs.Output()
		if out == nil {
			out = os.Stderr
		}

		fmt.Fprintf(out, "%s — sleep cycle calculator\n\n", name)

		fmt.Fprintf(out, "Usage:\n")
		fmt.Fprintf(out, "  %s [mode] [options]\n\n", name)

		fmt.Fprintf(out, "Choose exactly one mode:\n")
		fmt.Fprintf(out, "  -n, --now                    Calculate wake times from the current time\n")
		fmt.Fprintf(out, "  -w, --wake HH:MM             Calculate bedtimes for a target wake time\n")
		fmt.Fprintf(out, "  -s, --sleep HH:MM            Calculate wake times for a target sleep time\n")
		fmt.Fprintf(out, "  -f, --from HH:MM             Start of sleep window. Use together with --to\n")
		fmt.Fprintf(out, "  -t, --to HH:MM               End of sleep window. Use together with --from\n\n")

		fmt.Fprintf(out, "Options:\n")
		fmt.Fprintf(out, "  -b, --buffer int             Minutes to fall asleep (default: 15)\n")
		fmt.Fprintf(out, "  -m, --cycles-min int         Minimum cycles to show (default: 4)\n")
		fmt.Fprintf(out, "  -x, --cycles-max int         Maximum cycles to show (default: 6)\n\n")

		fmt.Fprintf(out, "Other:\n")
		fmt.Fprintf(out, "  -g, --good-night             Print random good night art\n")
		fmt.Fprintf(out, "  -v, --version                Print version\n")
		fmt.Fprintf(out, "  -h, --help                   Show help\n\n")

		fmt.Fprintf(out, "Examples:\n")
		fmt.Fprintf(out, "  %s --now\n", name)
		fmt.Fprintf(out, "  %s --wake 07:00\n", name)
		fmt.Fprintf(out, "  %s --sleep 22:30\n", name)
		fmt.Fprintf(out, "  %s --from 22:00 --to 07:00\n", name)
		fmt.Fprintf(out, "  %s --wake 07:00 --buffer 20 --cycles-min 5 --cycles-max 6\n\n", name)

		fmt.Fprintf(out, "Notes:\n")
		fmt.Fprintf(out, "  - Time format is 24-hour HH:MM\n")
		fmt.Fprintf(out, "  - Short hours like 7:00 are accepted\n")
		fmt.Fprintf(out, "  - Modes cannot be combined\n")
	}
}
