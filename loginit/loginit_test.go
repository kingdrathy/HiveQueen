package loginit

import(
	"os"
	"github.com/op/go-logging"
	"testing"
)

//var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,)
func setupLogger(){
	var log = logging.MustGetLogger("example")
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	logging.SetFormatter(format)
	backendLevel := logging.AddModuleLevel(backend1)
	log.SetBackend(backendLevel)
}

func TestErrorOnly(t *testing.T) {
	logging.Reset()
	setupLogger()
	var log = logging.MustGetLogger("example")
	logging.SetLevel(logging.ERROR, "example")

	log.Debug("This should not show up")
	log.Info("This should not show up")
	log.Warning("This should not show up")
	log.Error("This should show up")
	log.Critical("This should show up")
}

func TestWarnOnly(t *testing.T) {
	logging.Reset()
	setupLogger()
	var log = logging.MustGetLogger("example")
	logging.SetLevel(logging.WARNING, "example")
	log.Debug("This should not show up")
	log.Info("This should not show up")
	log.Warning("This should show up")
	log.Error("This should show up")
	log.Critical("This should show up")
}



//func InitializeLogger() {
//	var log = logging.MustGetLogger("example")

///	var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,)



//	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
//	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

///	backend1Formatter := logging.NewBackendFormatter(backend1, format)
//	backend2Formatter := logging.NewBackendFormatter(backend2, format)
//	backend2Formatter.SetLevel(logging.ERROR, "")

//	backend1Leveled := logging.AddModuleLevel(backend1)
//	backend1Leveled.SetLevel(logging.INFO, "")

//	logging.SetBackend(backend1Leveled)//, backend1Formatter)
	//logging.SetBackend(backend1)

//	log.Debugf("debug %s")
	//log.Info("info. Password: ", Password("Secret"))
//	log.Info("Logger initialized. ")
//	log.Error("debug %s")
//	return log
//}
