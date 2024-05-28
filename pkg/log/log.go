package log

import (
	"log"
	"os"
)

const (
	logDirPath  = "logs"
	logFileName = "app.log"
)

func LogInit() {
	file, err := getLogFileWriter(logDirPath, logFileName)
	if err != nil {
		log.Fatal("unexpected error during log file setup")
	}
	log.SetOutput(file)
}

func getLogFileWriter(path string, fileName string) (*os.File, error) {
	filePath := path + "/" + fileName
	logDirSetUp(path)
	logFileSetUp(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("unexpected error during creating log file Writer, on path: %s", filePath)
	}
	return file, err
}

func logDirSetUp(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatalf("can't create log directory on path: %s ", path)
		}
	} else if err != nil {
		log.Fatalf("unexpected error on path: %s", path)
	}
}

func logFileSetUp(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("can't create a log file on path: %s", filePath)
		}
		defer file.Close()
	}
}
