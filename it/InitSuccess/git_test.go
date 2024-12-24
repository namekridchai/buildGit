package mainTest

import (
	"os"
	"testing"

	"github.com/namekridchai/buildGit/git"
)

func TestInitSuccess(t *testing.T) {
	path := ".cgit"
    git.Init()
	dir, err := os.Stat(path)
	if err != nil {
        if os.IsNotExist(err) {
            t.Fatalf(".cgit does not exist")
        }
		t.Fatalf("Init should create directory successfully but got error %v",err)
	}
    if !dir.IsDir(){
        t.Fatal("init does not create directory but file")
    }
    os.Remove(path)
	
}
