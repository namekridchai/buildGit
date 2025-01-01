package mainTest

import (
	"os"
	"testing"

	"github.com/namekridchai/buildGit/command"
)

func TestInitWithFileExist(t *testing.T) {
	path := ".cgit"
	command.Init()
	dir, _ := os.Stat(path)

	if dir.IsDir() {
		t.Fatal("should see .cgit  as a file not as a directory")
	}

}