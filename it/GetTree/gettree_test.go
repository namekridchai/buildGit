package mainTest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/namekridchai/buildGit/command"
)

func TestWritetreeShouldGetCorrectContentonDirectory(t *testing.T) {
	GitRootdir := ".cgit"
	os.Mkdir(GitRootdir, 0755)

	parentDir := "test"
	childFile := "test.txt"

	fileObjectID := getHashStringFromFile(parentDir + "/" + childFile)
	content := fmt.Sprintf("%s %s %s", "blob", fileObjectID, childFile)

	dirObjectID := hash("tree", content)

	command.WriteTree(parentDir)

	file, err := os.OpenFile(parentDir + "/" + childFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	command.GetTree(parentDir, dirObjectID)
	fileObjectIDActual := getHashStringFromFile(parentDir + "/" + childFile)
	if fileObjectID != fileObjectIDActual {
		t.Fatalf("expect %v but got %v", fileObjectID, fileObjectIDActual)
	}

	os.RemoveAll(GitRootdir)

}

func getHashStringFromFile(parentDir string) string {

	expect, _ := os.ReadFile(parentDir)

	return hash("blob", string(expect))
}

func hash(typo string, content string) string {
	expect_content := typo + "\x00" + content

	hash := sha256.New()
	hash.Write([]byte(expect_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
