package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GetModuleName() string {
	file, _ := os.Open(filepath.Base("../../go.mod"))

	defer file.Close()

	f := bufio.NewReader(file)
	b, _, _ := f.ReadLine()

	return string(b[7:])
}

func IsModuleExist() (bool, error) {
	file, error := os.Open(filepath.Base("../../go.mod"))
	if error != nil {
		return false, fmt.Errorf("please initial module first")
	}
	defer file.Close()

	return true, nil
}
