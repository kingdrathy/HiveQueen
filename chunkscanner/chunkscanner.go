package chunkscanner

import(
	//"fmt"
	"HiveQueen/loginit"
	"regexp"
	"github.com/op/go-logging"
)


func check(e error) {
	if e != nil{
		//log.Error(e)
		panic(e)
	}
}


func ScanChunk(testString string){
	//get the logger going
	loginit.InitializeConsoleLogger()
	var log = logging.MustGetLogger("console")
	log.Debug("Program started. Logger initialized")

	//Compile the regex to scan for
	//TODO Make this driven by an external object
	//TODO Store results in an object
	//For a test, look for a 5 char word
	var regexPattern = regexp.MustCompile(`\W\w{15}\W`)


	//Compile the dictionary string to scan for
	//TODO Make this also driven by an external object
	//TODO test using standard string matching
	//TODO Store these results in the same type of object
	//for a test, look for "config" or "value"
	var dictPattern = regexp.MustCompile(`Moscow`)

	//Compare string to patterns to ID a match
	if regexPattern.MatchString(testString) {
		log.Infof("Regex match for '%s'",regexPattern.FindString(testString))
//		log.Infof("Regex Match for '%s': %v",testString,true)
	} else {
		log.Debugf("Regex Match for '%s': %v",testString,false)
	}
	if dictPattern.MatchString(testString) {
		log.Infof("Dict match for '%s': %v",testString,true)
	} else {
		log.Debugf("Dict match for '%s': %v",testString,false)
	}

}
