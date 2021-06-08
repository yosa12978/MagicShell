package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintLogo(username string) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()
	fmt.Println(` ███╗   ███╗ █████╗  ██████╗ ██╗ ██████╗ `)
	fmt.Println(` ████╗ ████║██╔══██╗██╔════╝ ██║██╔════╝ `)
	fmt.Println(` ██╔████╔██║███████║██║  ███╗██║██║      `)
	fmt.Println(` ██║╚██╔╝██║██╔══██║██║   ██║██║██║      `)
	fmt.Println(` ██║ ╚═╝ ██║██║  ██║╚██████╔╝██║╚██████╗ `)
	fmt.Println(` ╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝ ╚═════╝ `)
	fmt.Println()
	fmt.Println(` ███████╗██╗  ██╗███████╗██╗     ██╗     `)
	fmt.Println(` ██╔════╝██║  ██║██╔════╝██║     ██║     `)
	fmt.Println(` ███████╗███████║█████╗  ██║     ██║     `)
	fmt.Println(` ╚════██║██╔══██║██╔══╝  ██║     ██║     `)
	fmt.Println(` ███████║██║  ██║███████╗███████╗███████╗`)
	fmt.Println(` ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝`)

	fmt.Printf("User - %s\n\n", username)
}
