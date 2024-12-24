package git

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var (
	dir = ".cgit"
)

func Init() {
	fmt.Println("init custom git")
	exist, err := IsDirExist(dir)
	if err != nil {
		return
	}

	if !exist {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}

}

func Hash(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	hash := sha256.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	_, err = os.Stat(dir+"/object")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(dir+"/object",0755)
			if err!=nil{
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	
	file, err := os.Create(dir+"/object/"+hashString)
	if err != nil {
		panic(err)
	}
	file.Close()
}

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
