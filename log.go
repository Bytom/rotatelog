package rotatelog

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/bytom/bytom-netstats/config"
)

const logPath = "logs"

// InitLogFile init logrus with hook
func InitLogFile(module string, logs config.Logs) error {
	rotateTime, maxAge, err := logs.Durations()
	if err != nil {
		return err
	}

	if err := clearLockFiles(logPath); err != nil {
		return err
	}

	logrus.AddHook(newRotateHook(logPath, module, rotateTime, maxAge))
	logrus.SetOutput(ioutil.Discard)
	logLevel, err := logs.Level()
	if err != nil {
		logrus.WithField("error", err).Fatal("wrong log level")
	}

	logrus.SetLevel(logLevel)
	fmt.Printf("all logs are output in the %s directory, log level:%s\n", logPath, logs.LogLevel)
	return nil
}

func clearLockFiles(logPath string) error {
	files, err := ioutil.ReadDir(logPath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	for _, file := range files {
		if ok := strings.HasSuffix(file.Name(), "_lock"); ok {
			if err := os.Remove(filepath.Join(logPath, file.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}
