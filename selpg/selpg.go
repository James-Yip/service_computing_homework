/*=================================================================

Program name:
	selpg (SELect PaGes implemented by Golang)

Purpose:
	Extract a specified range of pages from an input text file.

===================================================================*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

/*================================ types =========================*/

type sp_flags struct {
	start_page  int
	end_page    int
	page_len    int
	useFormFeed bool
	in_filename string
	print_dest  string
}

/*=============================== globals =========================*/

var progname string = "selpg" // program name, for error messages

/*================================ main ===========================*/

func main() {
	var sf sp_flags
	parse_input(&sf)
	check_input(sf)
	process_input(sf)
}

/*============================ parse_input ========================*/
// parse the input flags & arguments
// and use them to initialize struct sf
func parse_input(sf *sp_flags) {
	flag.IntVar(&sf.start_page, "s", -1, "the start page number to extract an input text(mandatory)")
	flag.IntVar(&sf.end_page, "e", -1, "the end page number to extract an input text(mandatory)")
	flag.IntVar(&sf.page_len, "l", 72, "#lines of each page")
	flag.BoolVar(&sf.useFormFeed, "f", false, "use form feed to define the pages of an input text")
	flag.StringVar(&sf.print_dest, "d", "", "the destination to output the selected pages")
	flag.Parse()
	sf.in_filename = flag.Arg(0)
}

/*============================ check_input ========================*/
// check the validity of the input flags
func check_input(sf sp_flags) {
	errMessage := ""
	switch {
	case flag.NArg() > 1:
		errMessage = "Only support one input text file."
	case sf.start_page < 1 || sf.end_page < 1:
		errMessage = "The start page / end page should be positive integer."
	case sf.start_page > sf.end_page:
		errMessage = "The end page should be greater than the start page."
	case sf.page_len < 1:
		errMessage = "Page length should be positive integer."
	case sf.useFormFeed && sf.page_len != 72:
		errMessage = "Flags -l & -f are mutually exclusive."
	}
	if errMessage != "" {
		fmt.Fprintf(os.Stderr, "%s: %s\n", progname, errMessage)
		flag.Usage()
		os.Exit(1)
	}
}

/*========================== process_input ========================*/

func process_input(sf sp_flags) {
	// declaration
	var fin io.ReadCloser      // input stream
	var fout io.WriteCloser    // output stream
	var line_ctr int           // line counter
	var page_ctr int           // page counter
	var err error              // error message
	var cmd *exec.Cmd          // execute extra command if flag -d is given
	var scanner *bufio.Scanner // scanner used to scan the input stream

	// set the input source
	if sf.in_filename == "" {
		fin = os.Stdin
	} else {
		fin, err = os.Open(sf.in_filename)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"%s: fail to open %s(%s)\n",
				progname, sf.in_filename, err)
			os.Exit(2)
		}
	}

	// set the output source
	if sf.print_dest == "" {
		fout = os.Stdout
	} else {
		cmd = exec.Command("sh", "-c", sf.print_dest)
		fout, err = cmd.StdinPipe()
		if fout == nil {
			fmt.Fprintf(
				os.Stderr,
				"%s: fail to open pipe(%s)\n",
				progname, sf.in_filename, err)
			os.Exit(3)
		}
	}

	// start processing
	scanner = bufio.NewScanner(fin)
	if !sf.useFormFeed {
		line_ctr = 0
		page_ctr = 1
		// process type1 input text
		for scanner.Scan() {
			line_ctr++
			if line_ctr > sf.page_len {
				page_ctr++
				line_ctr = 1
			}
			if (page_ctr >= sf.start_page) && (page_ctr <= sf.end_page) {
				fmt.Fprintf(fout, "%s\n", scanner.Text())
			}
		}
	} else {
		page_ctr = 1
		// process type2 input text
		for scanner.Scan() {
			for _, c := range scanner.Text() {
				if c == '\f' {
					page_ctr++
				} else if (page_ctr >= sf.start_page) && (page_ctr <= sf.end_page) {
					fmt.Fprintf(fout, "%s", string(c))
				}
			}
			if (page_ctr >= sf.start_page) && (page_ctr <= sf.end_page) {
				fmt.Fprintf(fout, "\n")
			}
		}
	}
	// end of processing

	if page_ctr < sf.start_page {
		// output error message
		fmt.Fprintf(
			os.Stderr,
			"%s: start_page(%d) greater than total pages (%d), no output written\n",
			progname, sf.start_page, page_ctr)
	} else if page_ctr < sf.end_page {
		// output error message
		fmt.Fprintf(
			os.Stderr,
			"%s: end_page(%d) greater than total pages (%d), less output than expected\n",
			progname, sf.end_page, page_ctr)
	}

	// execute extra command if flag -d is given
	if sf.print_dest != "" {
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: print fails\n", progname)
		}
	}

	fin.Close()
	fout.Close()
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}
