package logutils

import "log"

func LogInfo(logStr string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("\"\\n\\n\\x1b[31m[info]")
	log.Println(logStr)
}

func LogErr(logStr string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[error]")
	log.Println(logStr)
}
