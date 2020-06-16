package main

import (
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Println(args[0] + ` usage:
	oto [flags] paths [[path2] [path3]...]`)
		fmt.Println(`
flags:`)
		flags.PrintDefaults()
	}
	var (
		template   = flags.String("template", "", "plush template to render")
		outfile    = flags.String("out", "", "output file (default: stdout)")
		pkg        = flags.String("pkg", "", "explicit package name (default: inferred)")
		v          = flags.Bool("v", false, "verbose output")
		paramsStr  = flags.String("params", "", "list of parameters in the format: \"key:value,key:value\"")
		ignoreList = flags.String("ignore", "", "comma separated list of interfaces to ignore")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if *template == "" {
		return errors.New("missing template")
	}
	params, err := parseParams(*paramsStr)
	if err != nil {
		return errors.Wrap(err, "params")
	}
	parser := newParser(flags.Args()...)
	ignoreItems := strings.Split(*ignoreList, ",")
	if ignoreItems[0] != "" {
		parser.ExcludeInterfaces = ignoreItems
	}
	parser.Verbose = *v
	if parser.Verbose {
		fmt.Println("oto - github.com/pacedotdev/oto")
	}
	def, err := parser.parse()
	if err != nil {
		return err
	}
	if *pkg != "" {
		def.PackageName = *pkg
	}
	b, err := ioutil.ReadFile(*template)
	if err != nil {
		return err
	}
	out, err := render(string(b), def, params)
	if err != nil {
		return err
	}
	var w io.Writer = stdout
	if *outfile != "" {
		f, err := os.Create(*outfile)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f

		// Apply gofmt to .go files
		if strings.HasSuffix(*outfile, ".go") {
			outb, err := format.Source([]byte(out))
			if err != nil {
				return err
			}
			out = string(outb)
		}
	}
	if _, err := io.WriteString(w, out); err != nil {
		return err
	}
	if parser.Verbose {
		var methodsCount int
		for i := range def.Services {
			methodsCount += len(def.Services[i].Methods)
		}
		fmt.Println()
		fmt.Printf("\tTotal services: %d", len(def.Services))
		fmt.Printf("\tTotal Methods: %d", methodsCount)
		fmt.Printf("\tTotal Objects: %d\n", len(def.Objects))
		fmt.Printf("\tOutput size: %s\n", humanize.Bytes(uint64(len(out))))
	}
	return nil
}

// parseParams returns a map of data parsed from the params string.
func parseParams(s string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if s == "" {
		// empty map for an empty string
		return params, nil
	}
	pairs := strings.Split(s, ",")
	for i := range pairs {
		pair := strings.TrimSpace(pairs[i])
		segs := strings.Split(pair, ":")
		if len(segs) != 2 {
			return nil, errors.New("malformed params")
		}
		params[strings.TrimSpace(segs[0])] = strings.TrimSpace(segs[1])
	}
	return params, nil
}
