package loginit

import(
	"os"
	"github.com/op/go-logging"

)

type Password string
func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func InitializeConsoleLogger() {
	//var log = logging.MustGetLogger("console")

	var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,)


	//setup the log backends
	//Errors go to standard error out, everything else to standard out
	consoleBackend := logging.NewLogBackend(os.Stdout, "", 0)
	consoleErrorBackend := logging.NewLogBackend(os.Stderr, "", 0)

	//Everything should have the same log format
	logging.SetFormatter(format)

///	backend1Formatter := logging.NewBackendFormatter(backend1, format)
//	backend2Formatter := logging.NewBackendFormatter(backend2, format)
//	backend2Formatter.SetLevel(logging.ERROR, "")

	//set the backend levels. 
	consoleBackendLevel := logging.AddModuleLevel(consoleBackend)
	consoleBackendLevel.SetLevel(logging.INFO, "")

	consoleErrorBackendLevel := logging.AddModuleLevel(consoleErrorBackend)
	consoleErrorBackendLevel.SetLevel(logging.ERROR, "")

	//backend1Leveled := logging.AddModuleLevel(backend1)
//	backend1Leveled.SetLevel(logging.INFO, "")

//	log.SetBackend(consoleBackendLevel,consoleErrorBackendLevel)//, backend1Formatter)
	logging.SetBackend(consoleBackendLevel,consoleErrorBackendLevel)
	//logging.SetBackend(backend1)

//	log.Debugf("debug %s")
	//log.Info("info. Password: ", Password("Secret"))
//	log.Info("Logger initialized. ")
//	log.Error("debug %s")
//	return log
}
