package chunkscanner

import "testing"

func TestBasic(t *testing.T){
	ScanChunk("This is just a test")
	ScanChunk("This is justy a test")
	ScanChunk("This is just a test config")
	ScanChunk("This is justy a test config")
}
