package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func IsDirExist(dirName string) (bool, error) {
	info, found, err := IsPathExist(dirName)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}
	if info.IsDir() {
		return true, nil
	} else {
		msg := fmt.Sprintf("path exists but is not a directory:%v", dirName)
		return false, errors.New(msg)
	}

}

func IsPathExist(path string) (fs.FileInfo, bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return info, false, nil
		} else {
			fmt.Println("there is err:", err)
			return info, false, err
		}
	}
	return info, true, nil
}

func IsFileExist(filePath string) (bool, error) {
	info, found, err := IsPathExist(filePath)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}
	if !info.IsDir() {
		return true, nil
	} else {
		msg := fmt.Sprintf("path exists but is not a file:%v", filePath)
		return false, errors.New(msg)
	}
}

func CreatDirIfNotExist(dirname string) error {
	exist, err := IsDirExist(dirname)
	if err != nil {
		return err
	}

	if !exist {
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}
	return nil
}

func CreateAndWriteFile(path string, content string) (err error) {

	var file *os.File
	file, err = os.Create(path)

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			fmt.Printf("there is an err %v", err)
		}
	}(file)

	if err != nil {
		return
	}
	_, err = file.Write([]byte(content))
	if err != nil {
		return
	}
	return

}
