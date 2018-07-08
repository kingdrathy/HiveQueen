package filechunkstreamer

import (
	//"HiveQueen/loginit"
	"github.com/op/go-logging"
	"bufio"
	"HiveQueen/shared"
	//"fmt"
	//"io"
	//"io/ioutil"
	"os"
	//"HiveQueen/chunkscanner"
)



func check(e error) {
	if e != nil{
		//log.Error(e)
		panic(e)
	}
}
func StreamFileChunks(scanWidth int, incrementCount int, filePath string) {

	//get the logging engine going
	shared.InitializeConsoleLogger()
	var log = logging.MustGetLogger("console")
	log.Info("Logger initialized")

	//Set the slice that will hold the output of the scan
	//ScanResultOffset := []int{}
	//ScanResultText := []string{}
	ScanResults := make([]shared.ResultList,0)
	
	//set the scan width
	//This will be the number of bytes the slice will have in memory
	//at one time. Represents the size of the scan window.
	//Larger scan window will collect more data, but use more memory
	//scanWidth := 25

	//Set the increment count.
	//This is the number of bytes that the window will move each cycle
	//as it scans thru the file. 
	//Smaller numbers represent more accurate scans, but require more 
	//file interactions, which may lead to slower processing times.
	//Latency multiplies this effect
	//Larger numbers represent more speed, but may lead to missed data
	//If this is set greater than the scan width, bad things happen
	//incrementCount := 5

	//load the test file
	//dat, err := ioutil.ReadFile(filePath)
	//check(err)
	rawFile, err := os.Open(filePath)
	check(err)
	defer rawFile.Close()
	fileStats, err := rawFile.Stat()
	fileSize := int(fileStats.Size())
	check(err)
	log.Info("File opened")
	log.Debugf("File size: %d",fileSize)

	//setup the buffered reader
	//TODO research what this does for memory caching
	buffRead := bufio.NewReader(rawFile)
	//capture a snapshot of the first 40 bytes of the file.
	//Also verify that the file opened properly
	n1, err := buffRead.Peek(40)
	check(err)
	log.Debugf("40 bytes: %s",string(n1))

	//Setup the slice that will hold the bytes being scanned
	//Uses the size from the scanWidth variable
	scanSlice := make([]byte,scanWidth)

	log.Info("Preparing to build initial slices")

	//perform the inital data load of the slice
	//populates with data from the file
	for i := 0; i < scanWidth; i++ {
		temp, err := buffRead.ReadByte()
		check(err)
		//log.Debug("InitLoad temp contents: ",string(temp))
		scanSlice[i] = temp
		//log.Debug("InitLoad scanSlice contents: ",string(scanSlice[i]))
	}	
	log.Debug("Initial scan slice contents: ",string(scanSlice))

	//begin the chunk streaming process
	//len(dat)
	bytesProcessed := scanWidth
	
	ScanResults = ScanChunk(string(scanSlice), 0, ScanResults)//) append(ScanResults,
	
	
	for bytesProcessed < fileSize {
//		log.Debug("Shifted scan slice contents: ",string(scanSlice))
		//Check to see if we are about to  increment more bytes 
		//than are available. If so, shorten the increment to what 
		//is available.
		if fileSize-bytesProcessed <incrementCount {
			incrementCount = fileSize-bytesProcessed
		}
		for z := 0; z < incrementCount; z++ {
			//shift the slice by one position, dropping the first element
			scanSlice = scanSlice[1:]
			//read in the next byte from the file
			temp, err := buffRead.ReadByte()
			check(err)
			//attach the new byte to the slice
			scanSlice = append(scanSlice, temp)
//			log.Debugf("z state: %d",z)

		}
		//Scan the chunk for matching criteria
		ScanResults = ScanChunk(string(scanSlice), bytesProcessed, ScanResults)//) append(ScanResults,
		//increment the i counter by the number of bytes moved
		bytesProcessed = bytesProcessed + incrementCount
		//log.Debugf("Bytes read: %d. Increment: %d. New scan slice snapshot: %s",bytesProcessed,incrementCount,string(scanSlice))
	}
	log.Info(ScanResults)
	log.Debugf("Bytes processed: %d. Last snapshot of the slice: %s",bytesProcessed,string(scanSlice))

}


