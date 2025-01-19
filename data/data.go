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

func GetContentfromObjId(objectId string) ([]string, error) {
	path := util.GitRootdir + "/object/" + objectId

	exist, err := util.IsFileExist(path)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic("file path does not exist")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	splitedContent := strings.SplitN(string(data), "\x00", 2)
	if len(splitedContent) == 1 {
		panic("invalid content should get 2 parts after split")
	}
	return splitedContent, nil
}
