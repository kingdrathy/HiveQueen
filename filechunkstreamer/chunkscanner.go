package filechunkstreamer

import(
	//"fmt"
	"HiveQueen/loginit"
	"regexp"
	"github.com/op/go-logging"
	"HiveQueen/shared"
	
)

func ScanChunk(testString string, ByteIndex int, ScanResults []shared.ResultList)[]shared.ResultList{
	//get the logger going
	loginit.InitializeConsoleLogger()
	var log = logging.MustGetLogger("console")
	log.Debug("Program started. Logger initialized")

	//Compile the regex to scan for
	//TODO Make this driven by an external object
	//TODO Store results in an object
	//For a test, look for a 5 char word
	//var regexPattern = regexp.MustCompile(`(asdfghjklp).{0,100}(asdfghjklp)`) //(.){50}\w{10}
	var regexPattern = regexp.MustCompile(`\d{3}\-\d{2}\-\d{4}`)


	//Compile the dictionary string to scan for
	//TODO Make this also driven by an external object
	//TODO test using standard string matching
	//TODO Store these results in the same type of object
	//for a test, look for "config" or "value"
	var dictPattern = regexp.MustCompile(`liberty`)

	//Compare string to patterns to ID a match
	if regexPattern.MatchString(testString) {
		//Capture the index of the matches 
		MatchLocations := regexPattern.FindAllStringIndex(testString, -1)
		log.Debugf("Number of matches in this slice: '%d'",len(MatchLocations))
		for z := 0; z < len(MatchLocations); z++ {
			//MatchLocations[z][0]+ByteIndex
			//check and see if we already have an entry for this match
			var StartingOffset int = MatchLocations[z][0]+ByteIndex
			if !CheckIfSliceContains(ScanResults,StartingOffset) {
				//var shared.ResultList
				ScanResults = append(ScanResults, shared.ResultList{ResultOffset:MatchLocations[z][0]+ByteIndex, ResultText:testString[MatchLocations[z][0]:MatchLocations[z][1]]})
				log.Infof("Adding entry to ScanResults: Offset: '%d', Text: '%s'",MatchLocations[z][0]+ByteIndex,testString[MatchLocations[z][0]:MatchLocations[z][1]])
				}
			
			//log.Infof("Start: '%d', End: '%d'. Text: '%s'",MatchLocations[z][0]+ByteIndex,MatchLocations[z][1]+ByteIndex,testString[MatchLocations[z][0]:MatchLocations[z][1]])
			//log.Infof("Regex match for '%s' at '%i'",regexPattern.FindAllString(testString, -1),MatchLocations)
		}
		//log.Infof("Regex match for '%s' at '%i'",regexPattern.FindString(testString),regexPattern.FindStringIndex(testString))
		//log.Infof("Regex match for '%s'",regexPattern.FindString(testString))
//		log.Infof("Regex Match for '%s': %v",testString,true)
	} else {
		log.Debugf("Regex Match for '%s': %v",testString,false)
	}
	if dictPattern.MatchString(testString) {
		//log.Infof("Dict match for '%s' at '%i'",dictPattern.FindString(testString),dictPattern.FindStringIndex(testString))
		//log.Infof("Dict match for '%s': %v",testString,true)
	} else {
		log.Debugf("Dict match for '%s': %v",testString,false)
	}
	return ScanResults

}

func CheckIfSliceContains(ScanResults []shared.ResultList, Offset int) bool{

	//Initialize logger
	//loginit.InitializeConsoleLogger()
	//var log = logging.MustGetLogger("console")
	
	for q := range ScanResults {
		if ScanResults[q].ResultOffset == Offset{
			return true
			}
		}
	return false
}