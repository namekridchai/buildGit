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

	expect_content := "tree" + "\x00" + content
	dirObjectID := hash("tree", content)

	hashedDirPath := GitRootdir + "/object/" + dirObjectID

	command.WriteTree(parentDir)

	actual, _ := os.ReadFile(hashedDirPath)
	if expect_content != string(actual) {
		t.Fatalf("expect %v but got %v", string(expect_content), string(actual))
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
