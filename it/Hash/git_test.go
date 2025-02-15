package mainTest

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"

	"github.com/namekridchai/buildGit/command"
)

func TestHashShouldGetCorrectFileName(t *testing.T) {
	GitRootdir := ".cgit"
	os.Mkdir(GitRootdir, 0755)
	path := "git_test.go"
	content, _ := os.ReadFile(path)
	expect_content := "blob" + "\x00" + string(content)

	hash := sha256.New()
	hash.Write([]byte(expect_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	targetPath := GitRootdir + "/object/" + hashString
	// actual, _ := os.ReadFile(targetPath)

	command.Hash(path, "blob")

	dir, err := os.Stat(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%v does not exist", targetPath)
		}
		t.Fatalf("Hash should create directory successfully but got error %v", err)
	}
	if dir.IsDir() {
		t.Fatal("hash does not create file but directory")
	}

	os.RemoveAll(GitRootdir)

}
