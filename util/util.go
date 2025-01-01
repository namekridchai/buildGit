package util

import (
	"errors"
	"fmt"
	"os"
)

func IsDirExist(dirName string) (bool, error) {
	info, err := os.Stat(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			fmt.Println("there is err:", err)
			return false, err
		}
	}
	if info.IsDir() {
		return true, nil
	} else {
		msg := fmt.Sprintf("path exists but is not a directory:%v", dirName)
		return false, errors.New(msg)
	}

}
