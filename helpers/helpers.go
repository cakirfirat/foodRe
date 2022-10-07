package helpers

import (
	"flag"
	"log"
	"os"
)

var APP_NAME = "foodRe"

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

var (
	Log *log.Logger = Loggerx()
)

func Loggerx() *log.Logger {
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	//En el caso que la variable de entorno exista, el sistema usa la configuración del docker.
	if LOG_FILE_LOCATION == "" {
		LOG_FILE_LOCATION = "../logs/" + APP_NAME + ".log"
	} else {
		LOG_FILE_LOCATION = LOG_FILE_LOCATION + APP_NAME + ".log"
	}
	flag.Parse()
	//Si el archivo existe se rehusa, es decir, no elimina el archivo log y crea uno nuevo.
	if _, err := os.Stat(LOG_FILE_LOCATION); os.IsNotExist(err) {
		file, err1 := os.Create(LOG_FILE_LOCATION)
		if err1 != nil {
			panic(err1)
		}
		//si no existe,se crea uno nuevo.
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		//si existe se rehusa.
		file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
