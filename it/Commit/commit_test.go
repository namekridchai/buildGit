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

	objectID := command.Commit(commitMsg, parentDir)
	hashedFilePath := GitRootdir + "/object/" + objectID
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

	os.RemoveAll(GitRootdir)

}
