package command

import (
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/cobra"
	"log"
	"os"
	"os/exec"
)

// npm just run local go bin
var npmCommand = &cobra.Command{
	Use:   "npm",
	Short: "运行 PATH/npm 的命令",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("yoimiya npm: should install npm in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
