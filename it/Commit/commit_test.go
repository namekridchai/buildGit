package commit_test

import (
	"os"
	"strings"
	"testing"

	"github.com/namekridchai/buildGit/command"
)

func TestCommitShouldSaveContentIntoObjectDatabaseCorrectly(t *testing.T) {
	GitRootdir := ".cgit"
	os.Mkdir(GitRootdir, 0755)

	parentDir := "test"
	commitMsg := "test commit"
	objectType := "tree"
	os.Mkdir(parentDir, 0755)

	objectIDParent := command.Commit(commitMsg, parentDir)
	hashedFilePath := GitRootdir + "/object/" + objectIDParent
	_, err := os.Stat(hashedFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Fatalf("%v does not exist", hashedFilePath)
		}
		t.Fatalf("Commit should create directory successfully but got error %v", err)
	}

	actual, _ := os.ReadFile(hashedFilePath)
	if !strings.Contains(string(actual), objectType) {
		t.Fatalf("expect %v but got %v", objectType, string(actual))
	}
	if !strings.Contains(string(actual), commitMsg) {
		t.Fatalf("expect %v but got %v", commitMsg, string(actual))
	}

	headFilePath := GitRootdir + "/HEAD"
	headContent, _ := os.ReadFile(headFilePath)
	if string(headContent) != objectIDParent {
		t.Fatalf("expect HEAD to point to %v but got %v", objectIDParent, string(headContent))
	}

	objectIDChild := command.Commit(commitMsg, parentDir)
	hashedFilePathChild := GitRootdir + "/object/" + objectIDChild
	actualChild, _ := os.ReadFile(hashedFilePathChild)
	if !strings.Contains(string(actualChild), objectIDParent) {
		t.Fatalf("expect %v but got %v", objectIDParent, string(actualChild))
	}

	os.RemoveAll(GitRootdir)
	os.RemoveAll(parentDir)

}
