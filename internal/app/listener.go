package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/yosa12978/MagicShell/internal/app/handlers"
	"github.com/yosa12978/MagicShell/pkg/helpers"
)

func Run() {
	mainHandler := handlers.MainHandler{}

	userdir, _ := os.UserHomeDir()
	os.Chdir(userdir)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	helpers.PrintLogo(user.Username)

	reader := bufio.NewScanner(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("%s --> ", dir)

		reader.Scan()
		data := reader.Text()
		info := strings.Split(data, " ")

		switch strings.ToLower(info[0]) {

		case "cd":
			mainHandler.CdHandler(info[1:])
		case "exit":
			mainHandler.ExitHandler()
		case "ls":
			mainHandler.LsHandler(dir)
		case "echo":
			mainHandler.EchoHandler(info[1:])
		case "clear":
			mainHandler.ClearHandler()
		case "touch":
			mainHandler.TouchHandler(info[1])
		case "mkdir":
			mainHandler.MkdirHandler(info[1])
		case "cat":
			mainHandler.CatHandler(info[1:])
		case "setenv":
			mainHandler.SetenvHandler(info[1], info[2])
		case "getenv":
			mainHandler.GetenvHandler(info[1])
		case "sha256":
			mainHandler.Sha256Handler(strings.Join(info[1:], " "))
		case "sha1":
			mainHandler.Sha1Handler(strings.Join(info[1:], " "))
		case "md5":
			mainHandler.MD5Handler(strings.Join(info[1:], " "))
		case "help":
			mainHandler.HelpHandler()
		default:
			mainHandler.ProgrammHandler(helpers.ConcatArray(info))

		}
	}
}
