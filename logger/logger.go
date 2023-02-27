package logger

//////////////////////////////////////////////////
// 2023.02.28 @yjkim
// v0.0.1: init
//
//////////////////////////////////////////////////
import (
	"fmt"
	"log"
	"os"
	"time"
)

//////////////////////////////////////////////////
const (
	FILE = 0
)

//////////////////////////////////////////////////
type Logger struct {
	fileName  string
	Prefix    string
	Directory string
	Date      *string
	logFD     *os.File
}

//////////////////////////////////////////////////
func (l *Logger) InitStandardLogger(logType uint) {
	if logType == 0 {
		if l.Directory == "" {
			l.Directory = "logs"
		}
		if _, err := os.Stat(l.Directory); os.IsNotExist(err) {
			err = os.MkdirAll(l.Directory, 0755)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
		go func() {
			defer l.logFD.Close()
			for {
				l.logLotate()
				time.Sleep(time.Second * 10)
			}
		}()
	}
	time.Sleep(time.Second * 1)
	log.Println("==============================")
}

//////////////////////////////////////////////////
func (l *Logger) logLotate() {
	var err error
	tmFormat := "20060102"
	timeDate := time.Now().Format(tmFormat)
	if l.Date == nil {
		l.Date = &timeDate
	}
	_, err = os.Stat(l.fileName)
	if l.fileName == "" || *l.Date != timeDate || err != nil {
		if l.logFD != nil {
			l.logFD.Close()
		}
		l.fileName = fmt.Sprintf("%s/%s-%s.log", l.Directory, l.Prefix, timeDate)
		l.logFD, err = os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		l.Date = &timeDate
		if err != nil {
			log.Fatalln(err)
		}
		log.SetOutput(l.logFD)
	}
}
