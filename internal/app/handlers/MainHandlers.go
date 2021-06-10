package handlers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type MainHandler struct{}

func (mh *MainHandler) ExitHandler() {
	os.Exit(0)
}

func (mh *MainHandler) CdHandler(args []string) {
	os.Chdir(strings.Join(args, " "))
}

func (mh *MainHandler) LsHandler(dir string) {
	finfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Files of directory %s\n----------\n", dir)
	for _, file := range finfo {
		if file.IsDir() {
			fmt.Printf("%s [Directory]\n", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}

func (mh *MainHandler) ClearHandler() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (mh *MainHandler) EchoHandler(elems []string) {
	fmt.Println(strings.Join(elems, " "))
}

func (mh *MainHandler) ProgrammHandler(info []string) {
	cmd := exec.Command(info[0], info[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (mh *MainHandler) TouchHandler(name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()
	fmt.Printf("File %s created\n", name)
}

func (mh *MainHandler) MkdirHandler(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Dir %s created\n", name)
}

func (mh *MainHandler) CatHandler(filenames []string) {
	for _, name := range filenames {
		file, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(file))
	}
}

func (mh *MainHandler) SetenvHandler(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Created env var %s\n", key)
}

func (mh *MainHandler) GetenvHandler(key string) {
	res := os.Getenv(key)
	fmt.Println(res)
}

func (mh *MainHandler) Sha256Handler(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := sha256.Sum256(data)
	fmt.Printf("%x\n", res)
}

func (mh *MainHandler) Sha1Handler(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := sha1.Sum(data)
	fmt.Printf("%x\n", res)
}

func (mh *MainHandler) MD5Handler(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := md5.Sum(data)
	fmt.Printf("%x\n", res)
}

func (mh *MainHandler) HelpHandler() {
	data, err := ioutil.ReadFile("help.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(data))
}
