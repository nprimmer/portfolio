package main

import (
	"flag"
	"fmt"
)

func main() {
	f := false
	showFlag := flag.Bool("show", f, "show flag")
	helpFlag := flag.Bool("help", f, "help")
	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: ./main [options]")
		flag.PrintDefaults()
		return
	}

	if !*showFlag {
		fmt.Println("Lightcycle Arena Password Storage System")
		return
	}

	fmt.Println("gc24{9261fb7d-22f0-4b3b-a5e7-a7316d1c4904}")
}
