package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please, provide a module name.")
		return
	}
	os.Mkdir(os.Args[1], 0755)
	os.Chdir(os.Args[1])
	cmd := exec.Command("go", "mod", "init", os.Args[1])
	os.WriteFile("main.go", []byte("package main\n\nfunc main() {\n\tprintln(\"Hello, World!\")\n}\n"), 0644)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))

}
