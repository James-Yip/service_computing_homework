package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type sp_flags struct {
	start_page  int
	end_page    int
	page_len    int
	useFormFeed bool
	in_filename string
	print_dest  string
}

func main() {
	var sf sp_flags
	// parse the input flags and use them to initialize struct sf
	parse_flags(&sf)
	// check the validation of the input flags
	check_validation(sf)
	process_input(sf)
}

func parse_flags(sf *sp_flags) {
	flag.IntVar(&sf.start_page, "s", -1, "the start page number to extract an input text(mandatory)")
	flag.IntVar(&sf.end_page, "e", -1, "the end page number to extract an input text(mandatory)")
	flag.IntVar(&sf.page_len, "l", 72, "#lines of each page")
	flag.BoolVar(&sf.useFormFeed, "f", false, "use form feed to define the pages of an input text")
	flag.StringVar(&sf.print_dest, "d", "", "the destination to print the select pages")
	flag.Parse()
	sf.in_filename = flag.Arg(0)
	fmt.Printf("The size of args is: %d\n", flag.NArg())
}

func check_validation(sf sp_flags) {
	errMessage := ""
	switch {
	case flag.NFlag() < 2:
		errMessage = "not enough arguments"
	case sf.start_page < 1 || sf.end_page < 1:
		errMessage = "The start_page/end_page should be positive integer."
	case sf.start_page > sf.end_page:
		errMessage = "The start_page can't larger than the end_page."
	case sf.useFormFeed && sf.page_len != 72:
		errMessage = "-l & -f is mutually exclusive."
	}
	if errMessage != "" {
		fmt.Printf("%s\n", errMessage)
		flag.Usage()
		os.Exit(1)
	}
}

func process_input(sf sp_flags) {
	// set the input source
	fin, err := os.Open(sf.in_filename)
	if sf.in_filename == "" {
		fin = os.Stdin
	} else if fin == nil {
		panic(err)
		os.Exit(2)
	}

	line_ctr := 0
	page_ctr := 1
	// start processing
	scanner := bufio.NewScanner(fin)
	if !sf.useFormFeed {
		for scanner.Scan() {
			line_ctr++
			if line_ctr > sf.page_len {
				page_ctr++
				line_ctr = 1
			}
			if (page_ctr >= sf.start_page) && (page_ctr <= sf.end_page) {
				fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
			}
		}
	} else {
		for _, c := range scanner.Text() {
			if c == '\f' {
				page_ctr++
			}
			if (page_ctr >= sf.start_page) && (page_ctr <= sf.end_page) {
				fmt.Fprintf(os.Stdout, "%s\n", string(c))
			}
		}
	}
	// end of processing

	if page_ctr < sf.start_page {
		// output error message
		fmt.Fprintf(os.Stderr, "start_page (%d) larger than total pages (%d), no output written\n", sf.start_page, page_ctr)
	} else if page_ctr < sf.end_page {
		// output error message
		fmt.Fprintf(os.Stderr, "end_page (%d) larger than total pages (%d), less output than expected\n", sf.end_page, page_ctr)
	}
	if fin != nil {
		fin.Close()
	}
	fmt.Printf("selpg: done\n")
}
