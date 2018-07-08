package main

import (
	"HiveQueen/filechunkstreamer"
	"flag"
	//"HiveQueen/shared"
)



func main() {
	var scanWidth int
	var incrementCount int
	var filePath string
	flag.IntVar(&scanWidth, "scanwidth", 500, "Set the size of the scan width")
	flag.IntVar(&incrementCount, "incrementcount", 5,"set the number of bytes to increment by")
	flag.StringVar(&filePath, "filepath", "c:/temp/CommonSense.txt","set the file to be parsed")
	flag.Parse()
	//scanWidth := 25
	//incrementCount := 5
	//filePath := "/tmp/dat"
	filechunkstreamer.StreamFileChunks(scanWidth,incrementCount,filePath)
}
