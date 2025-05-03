package mainTest

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"

	"github.com/namekridchai/buildGit/command"
)

const path = "hash_test.go"

func TestHashShouldGetCorrectFileName(t *testing.T) {
	GitRootdir := ".cgit"
	os.Mkdir(GitRootdir, 0755)

	expect, _ := os.ReadFile(path)
	expect_content := "blob" + "\x00" + string(expect)

	hash := sha256.New()
	hash.Write([]byte(expect_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	hashedFilePath := GitRootdir + "/object/" + hashString
	// actual, _ := os.ReadFile(hashedFilePath)

	command.Hash(path, "blob")

	dir, err := os.Stat(hashedFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%v does not exist", hashedFilePath)
		}
		t.Fatalf("Hash should create directory successfully but got error %v", err)
	}
	if dir.IsDir() {
		t.Fatal("hash does not create file but directory")
	}
	
	os.RemoveAll(GitRootdir)

}

func TestHashShouldGetCorrectContent(t *testing.T) {
	GitRootdir := ".cgit"
	os.Mkdir(GitRootdir, 0755)
	expect, _ := os.ReadFile(path)
	expect_content := "blob" + "\x00" + string(expect)

	hash := sha256.New()
	hash.Write([]byte(expect_content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	hashedFilePath := GitRootdir + "/object/" + hashString

	command.Hash(path, "blob")
	actual, _ := os.ReadFile(hashedFilePath)

	if string(expect_content) != string(actual) {
		t.Fatalf("expect %v but got %v", string(expect_content), string(actual))
	}

	os.RemoveAll(GitRootdir)

}
