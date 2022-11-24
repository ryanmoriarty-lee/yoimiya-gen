package foo

import (
	"fmt"
	"github.com/ryanmoriarty-lee/yoimiya-gen/framework/cobra"
)

var FooCommand = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("this is foo command")
		return nil
	},
}
