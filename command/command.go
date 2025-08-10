package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/namekridchai/buildGit/data"
	"github.com/namekridchai/buildGit/enum"
	"github.com/namekridchai/buildGit/util"
)

type objectDirContent struct {
	objectType enum.ObjectType
	objectId   string
	fileName   string
}

func Init() {
	fmt.Println("init custom git")
	err := util.CreateDirIfNotExist(util.GitRootdir)
	if err != nil {
		return
	}
}

func LogCommit() {
	headFilePath := util.GitRootdir + "/HEAD"
	exist, err := util.IsFileExist(headFilePath)
	if err != nil {
		panic(err)
	}
	if !exist {
		fmt.Println("No commits found.")
		return
	}

	headContent, _ := os.ReadFile(headFilePath)
	objectId := string(headContent)

	for objectId != "" {
		content, err := data.GetContentfromObjId(objectId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Commit ID: %s\n", objectId)

		lines := strings.Split(content.Content, "\n")
		fmt.Printf("Message: %s\n", lines[len(lines)-1])

		if strings.HasPrefix(lines[1], "parent ") {
			objectId = strings.TrimPrefix(lines[1], "parent ")

		} else {
			objectId = ""
		}
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

	if string(data.ObjectType) != typo {
		errMsg := fmt.Sprintf("cat file get mismatch type expect %v get %v", typo, data.ObjectType)
		panic(errMsg)
	}

	fmt.Printf("File content: %s\n", string(data.Content))

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

	var objectContents []objectDirContent

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

		content := objectDirContent{objectType: objectType, objectId: hashId, fileName: file.Name()}
		objectContents = append(objectContents, content)
	}

	var contentArray string
	for index, content := range objectContents {
		c := fmt.Sprintf("%v %v %v", content.objectType, content.objectId, content.fileName)
		contentArray += c
		if index != len(objectContents)-1 {
			contentArray += "\n"
		}
	}

	hashID, err := data.Hash([]byte(contentArray), enum.Tree)
	if err != nil {
		panic(err)
	}
	return hashID

}

func GetTree(rootPath string, objectId string) {
	found, err := util.IsDirExist(rootPath)

	if err != nil {
		panic(err)
	}
	if !found {
		return
	}

	err = clearDirectory(rootPath)
	if err != nil {
		panic(err)
	}

	directoryContent, err := data.GetContentfromObjId(objectId)
	if err != nil {
		panic(err)
	}

	typo, ok := enum.GetObjectType("tree")
	if !ok {
		panic("invalid object type")
	}

	if directoryContent.ObjectType != typo {
		errMsg := fmt.Sprintf("file  mismatch type expect %v get %v", typo, directoryContent.ObjectType)
		panic(errMsg)
	}

	lines := strings.Split(directoryContent.Content, "\n")
	for _, line := range lines {

		content := toObjectDirContent(line)
		if content.objectType == enum.Blob {
			fileContent, err := data.GetContentfromObjId(content.objectId)
			if err != nil {
				panic(err)
			}

			file, err := os.OpenFile(rootPath+"/"+content.fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = file.WriteString(fileContent.Content)
			if err != nil {
				panic(err)

			}
		} else {
			GetTree(rootPath+"/"+content.fileName, content.objectId)
		}
	}

}
func Commit(msg string, rootPath string) string {
	found, err := util.IsDirExist(rootPath)

	if err != nil {
		panic(err)
	}
	if !found {
		return ""
	}

	headFilePath := util.GitRootdir + "/HEAD"

	exist, err := util.IsFileExist(headFilePath)
	if err != nil {
		panic(err)
	}
	var hashID string
	var content string

	if !exist {
		content = fmt.Sprintf("tree %v\n\ncommit %v", WriteTree(rootPath), msg)
	} else {
		headObjectID, err := os.ReadFile(headFilePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			panic(err)
		}

		content = fmt.Sprintf("tree %v\nparent %v\ncommit %v", WriteTree(rootPath), string(headObjectID), msg)

	}
	hashID, err = data.Hash([]byte(content), enum.Commit)
	if err != nil {
		panic(err)
	}
	util.CreateAndWriteFile(headFilePath, hashID)

	return hashID
}

func clearDirectory(rootPath string) error {
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.RemoveAll(rootPath + "/" + file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func toObjectDirContent(line string) objectDirContent {
	splited := strings.Split(line, " ")
	if len(splited) != 3 {
		panic("invalid line")
	}
	objType, ok := enum.GetObjectType(splited[0])
	if !ok {
		panic("invalid object type")
	}
	return objectDirContent{objectType: objType, objectId: splited[1], fileName: splited[2]}
}
