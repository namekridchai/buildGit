package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/namekridchai/buildGit/data"
	"github.com/namekridchai/buildGit/enum"
	"github.com/namekridchai/buildGit/util"
)

type objectContent struct {
	objectType enum.ObjectType
	hashId     string
	fileName   string
}

func Init() {
	fmt.Println("init custom git")
	err := util.CreatDirIfNotExist(util.GitRootdir)
	if err != nil {
		return
	}
}

func Hash(filePath string, typo enum.ObjectType) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	hashID, err := data.Hash(content, typo)
	if err != nil {
		panic(err)
	}
	return hashID

}

func Cat(hash string, typo string) {
	path := util.GitRootdir + "/object/" + hash

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

func WriteTree(rootPath string) string {
	found, err := util.IsDirExist(rootPath)
	if err != nil {
		panic(err)
	}
	if !found {
		return ""
	}

	files, err := os.ReadDir(rootPath)
	if err != nil {
		panic(err)
	}

	var objectContents []objectContent

	for _, file := range files {
		if file.Name() == ".cgit" || file.Name() == ".git" {
			continue
		}
		fmt.Println(file.Name())
		path := rootPath + "/" + file.Name()

		var hashId string
		var objectType enum.ObjectType
		if file.IsDir() {
			objectType = enum.Tree
			hashId = WriteTree(path)
		} else {
			objectType = enum.Blob
			hashId = Hash(path, enum.Blob)
		}

		content := objectContent{objectType: objectType, hashId: hashId, fileName: file.Name()}
		objectContents = append(objectContents, content)
	}

	var contentArray string
	for _, content := range objectContents {
		c := fmt.Sprintf("%v %v %v\n", content.objectType, content.hashId, content.fileName)
		contentArray += c
	}

	hashID, err := data.Hash([]byte(contentArray), enum.Tree)
	if err != nil {
		panic(err)
	}
	return hashID

}
