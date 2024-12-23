package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime/debug"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  "test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("test called %v\n", os.Args)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "version",
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := debug.ReadBuildInfo()
		fmt.Println(version.Main.Version)
	},
}

func main() {
	root := &cobra.Command{}
	root.AddCommand(testCmd, versionCmd)
	fmt.Println(root.Version)
	if err := root.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Oops. An error while executing command '%s'\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}
