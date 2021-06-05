package config

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var errInvalidEnvVarFormat = errors.New("invalid format of environment variable")

// LoadFile reads dotenv passed in argument from cwd and exports them as env variables
func LoadFile(filename string) {
	if _, err := os.Stat(filename); err != nil {
		return
	}
	if err := ReadFile(filename); err != nil {
		log.Panic(err.Error())
	}
}

// ReadFile reads passed file path
func ReadFile(file string) (err error) {
	f, err := os.Open(filepath.Clean(file))
	if err != nil {
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		err = readLine(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func readLine(line string) error {
	line = strings.Trim(line, "\n ")
	if len(line) == 0 {
		return nil
	}
	if i := strings.Index(line, "#"); i > -1 {
		return readLine(line[:i])
	}
	parts := strings.Split(line, "=")
	if len(parts) < 2 {
		return errInvalidEnvVarFormat
	}
	parts[1] = strings.ReplaceAll(parts[1], "'", "")
	parts[1] = strings.ReplaceAll(parts[1], `""`, "")
	return os.Setenv(parts[0], parts[1])
}
