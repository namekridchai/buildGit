package command

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/namekridchai/buildGit/util"
)

var (
	dir = ".cgit"
)

func Init() {
	fmt.Println("init custom git")
	exist, err := util.IsDirExist(dir)
	if err != nil {
		return
	}

	if !exist {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}

}

func Hash(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	hash := sha256.New()
	hash.Write([]byte(content))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	savedDirectory := dir + "/object/"
	_, err = os.Stat(savedDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(savedDirectory, 0755)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	file, err := os.Create(dir + "/object/" + hashString)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(content)
		panic(err)
	}

	file.Close()
}

func Cat(hash string){
	
}
