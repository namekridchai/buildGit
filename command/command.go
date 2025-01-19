package command

import (
	"fmt"
	"os"

	"github.com/namekridchai/buildGit/data"
	"github.com/namekridchai/buildGit/enum"
	"github.com/namekridchai/buildGit/util"
)

type objectContent struct {
	objectType enum.ObjectType
	objectId   string
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

func Cat(objectId string, typo string) {
	data, err := data.GetContentfromObjId(objectId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File content: %s\n", string(data[1]))

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

		content := objectContent{objectType: objectType, objectId: hashId, fileName: file.Name()}
		objectContents = append(objectContents, content)
	}

	var contentArray string
	for _, content := range objectContents {
		c := fmt.Sprintf("%v %v %v\n", content.objectType, content.objectId, content.fileName)
		contentArray += c
	}

	hashID, err := data.Hash([]byte(contentArray), enum.Tree)
	if err != nil {
		panic(err)
	}
	return hashID

}

func getTree(rootPath string, objectId string) {
	found, err := util.IsDirExist(rootPath)
	if err != nil {
		panic(err)
	}
	if !found {
		return
	}

	// todo get content from object id

}
