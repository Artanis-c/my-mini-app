package logutils

import "log"

func LogInfo(logStr string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(" [info] " + logStr)
}

func LogErr(logStr string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(" [ERROR] " + logStr)
}
