package data

import (
	"crypto/sha256"
	"encoding/hex"

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
