package help

import (
	"bytes"
	"fmt"
	"io"
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

		if err := writeHelp(out, name); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to write help output: %v\n", err)
		}
	}
}

func writeHelp(w io.Writer, name string) error {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%s — sleep cycle calculator\n\n", name)
	b.WriteString("Usage:\n")
	fmt.Fprintf(&b, "  %s                    Launch interactive mode\n", name)
	fmt.Fprintf(&b, "  %s [mode] [options]\n\n", name)

	b.WriteString("Choose exactly one mode:\n")
	b.WriteString("  -n, --now                    Calculate wake times from the current time\n")
	b.WriteString("  -w, --wake HH:MM             Calculate bedtimes for a target wake time\n")
	b.WriteString("  -s, --sleep HH:MM            Calculate wake times for a target sleep time\n")
	b.WriteString("  -f, --from HH:MM             Start of sleep window. Use together with --to\n")
	b.WriteString("  -t, --to HH:MM               End of sleep window. Use together with --from\n\n")

	b.WriteString("Options:\n")
	b.WriteString("  -b, --buffer int             Minutes to fall asleep (default: 15)\n")
	b.WriteString("  -m, --cycles-min int         Minimum cycles to show (default: 4)\n")
	b.WriteString("  -x, --cycles-max int         Maximum cycles to show (default: 6)\n\n")

	b.WriteString("Other:\n")
	b.WriteString("  -g, --good-night             Print random good night art\n")
	b.WriteString("  -v, --version                Print version\n")
	b.WriteString("  -h, --help                   Show help\n\n")

	b.WriteString("Examples:\n")
	fmt.Fprintf(&b, "  %s --now\n", name)
	fmt.Fprintf(&b, "  %s --wake 07:00\n", name)
	fmt.Fprintf(&b, "  %s --sleep 22:30\n", name)
	fmt.Fprintf(&b, "  %s --from 22:00 --to 07:00\n", name)
	fmt.Fprintf(&b, "  %s --wake 07:00 --buffer 20 --cycles-min 5 --cycles-max 6\n\n", name)

	b.WriteString("Notes:\n")
	b.WriteString("  - Time format is 24-hour HH:MM\n")
	b.WriteString("  - Short hours like 7:00 are accepted\n")
	b.WriteString("  - Modes cannot be combined\n")

	_, err := io.WriteString(w, b.String())
	return err
}
