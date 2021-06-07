package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yosa12978/MagicShell/internal/app/handlers"
)

func Run() {
	mainHandler := handlers.MainHandler{}
	reader := bufio.NewScanner(os.Stdin)

	mainHandler.ClearHandler()

	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("%s [ ", dir)

		reader.Scan()
		data := reader.Text()
		info := strings.Split(data, " ")

		switch info[0] {

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
		default:
			mainHandler.ProgrammHandler(info)

		}
	}
}
