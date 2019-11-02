package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(stdout io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		template = flags.String("template", "", "plush template to render")
		outfile  = flags.String("out", "", "output file (default: stdout)")
		pkg      = flags.String("pkg", "", "explicit package name (default: inferred)")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	def, err := newParser(flags.Args()...).parse()
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
	out, err := render(string(b), def)
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
	}
	if _, err := io.WriteString(w, out); err != nil {
		return err
	}
	return nil
}
