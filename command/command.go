package command

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/namekridchai/buildGit/util"
)

var (
	dir  = ".cgit"
	blob = "blob"
)

func Init() {
	fmt.Println("init custom git")
	err := util.CreatDirIfNotExist(dir)
	if err != nil {
		return
	}
}

func Hash(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	save_content := blob + "\x00" + string(content)
	hash := sha256.New()
	hash.Write([]byte(save_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	savedDirectory := dir + "/object/"
	err = util.CreatDirIfNotExist(savedDirectory)
	if err != nil {
		return
	}

	file, err := os.Create(dir + "/object/" + hashString)
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte(save_content))
	if err != nil {
		fmt.Println(save_content)
		panic(err)
	}

	file.Close()
}

func Cat(hash string, typo string) {
	path := dir + "/object/" + hash

	exist, err := util.IsFileExist(path)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic("file path does not exist")
	}
	data, err := os.ReadFile(path)

	splitedContent := strings.SplitN(string(data), "\x00", 2)
	if len(splitedContent) == 1 {
		panic("get array length of 1 after split content in file")
	}
	if splitedContent[0] != typo {
		errMsg := fmt.Sprintf("cat file get mismatch type expect %v get %v", typo, splitedContent[0])
		panic(errMsg)

	}

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File content: %s\n", string(splitedContent[1]))

}

func WriteTree(path string) {
	found, err := util.IsDirExist(path)
	if err != nil {
		panic(err)
	}
	if !found {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			WriteTree(path + "/" + file.Name())
		}
	}

}
