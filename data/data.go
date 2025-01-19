package data

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/namekridchai/buildGit/enum"
	"github.com/namekridchai/buildGit/util"
)

type objectContent struct {
	ObjectType enum.ObjectType
	Content    string
}

func Hash(content []byte, typo enum.ObjectType) (hashID string, err error) {

	save_content := typo.GetObjectType() + "\x00" + string(content)
	hash := sha256.New()
	hash.Write([]byte(save_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	savedDirectory := util.GitRootdir + "/object/"
	err = util.CreatDirIfNotExist(savedDirectory)
	if err != nil {
		return "", err
	}

	newFilePath := savedDirectory + hashString
	if err = util.CreateAndWriteFile(newFilePath, save_content); err != nil {
		return "", err
	}

	return hashString, nil
}

func GetContentfromObjId(objectId string) (objectContent, error) {
	path := util.GitRootdir + "/object/" + objectId

	exist, err := util.IsFileExist(path)
	if err != nil {
		return objectContent{}, err
	}
	if !exist {
		return objectContent{}, fmt.Errorf("file not found")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return objectContent{}, err
	}

	splitedContent := strings.SplitN(string(data), "\x00", 2)
	if len(splitedContent) == 1 {
		panic("invalid content should get 2 parts after split")
	}
	objType, ok := enum.GetObjectType(splitedContent[0])
	if !ok {
		return objectContent{}, fmt.Errorf("invalid object type %v", splitedContent[0])
	}
	return objectContent{ObjectType: objType, Content: splitedContent[1]}, nil
}
