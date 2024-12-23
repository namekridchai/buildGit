// main.go
package main

import (
	"errors"
	"fmt"
	"os"
)

var(
	dir  = ".cgit"
)

func Init() {
    fmt.Println("init custom git")
	exist,err := IsDirExist(dir)
	if err!= nil{
		return 
	}

	if !exist{
		err := os.Mkdir(dir, 0755)
		if err!= nil{
			fmt.Println("Error creating directory:", err)
		}
	}
	
}

func IsDirExist(dirName string) (bool,error){
		info, err := os.Stat(dirName)
		if err != nil {
			if os.IsNotExist(err) {
				return false,nil
			} else {
				fmt.Println("there is err:", err)
				return false,err
			}
		} 
		if info.IsDir() {
			return true,nil
		} else {
			msg:=fmt.Sprintf("path exists but is not a directory:%v",dirName)
			return false,errors.New(msg)
		}
		
}


func main() {
    // Check the number of command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./main.go [command]")
        os.Exit(1)
    }

    // The first argument is the command
    command := os.Args[1]

    switch command {
	case "init":
		Init()
    default:
        fmt.Println("Unknown command:", command)
        os.Exit(1)
    }
}