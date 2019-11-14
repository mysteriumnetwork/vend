package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/mysteriumnetwork/vend/file"
)

// Options contains CLI arguments passed to the program.
type Options struct {
	Help        bool
	PkgOnly     bool
	ReplaceDeps []file.ReplaceDep
}

type replaceFlag struct {
	deps []file.ReplaceDep
}

func (r replaceFlag) String() string {
	return ""
}

func (r *replaceFlag) Set(val string) error {
	split := strings.Split(val, "=")
	if len(split) != 2 {
		return errors.New("invalid replace argument. See --help for information")
	}
	r.deps = append(r.deps, file.ReplaceDep{
		Path:     split[0],
		WithPath: split[1],
	})
	return nil
}

// ParseOptions parses the command line options and returns a struct filled with
// the relevant options.
func ParseOptions() Options {
	var opt Options

	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.BoolVar(&opt.PkgOnly, "package", false, "Only vendor package level dependencies.")
	var rf = &replaceFlag{}
	flag.Var(rf, "replace", "Replace deps. Sample: -replace original/path=replacement/path")
	flag.Parse()
	opt.ReplaceDeps = rf.deps

	return opt
}

// PrintUsage prints the usage of this tool.
func (opt *Options) PrintUsage() {
	const banner string = `                     _
__   _____ _ __   __| |
\ \ / / _ \ '_ \ / _' |
 \ V /  __/ | | | (_| |
  \_/ \___|_| |_|\__,_|

`

	color.Green(banner)
	fmt.Printf("A small command line utility for fully vendoring module dependencies\n\n")

	flag.Usage()
}
