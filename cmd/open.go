/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"mh4af/internal/model"
	"mh4af/internal/validation"
	"mh4af/pkg/common"
	"mh4af/pkg/open"
	"os"
)

const port = 54321

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open browser",
	Long:  `open browser`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Unknown Error Occured: %s\n", err)
			os.Exit(1)
		}
		if filepath == "" {
			fmt.Println("filepath is required")
			os.Exit(1)
		}

		yaml := model.Definition{}
		err = common.ReadYaml(&yaml, filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		model, err := validation.Validate(yaml)
		if err != nil {
			panic(err)
		}

		listen := make(chan bool)
		go func() {
			<-listen
			url := fmt.Sprintf("http://localhost:%d", port)
			if err = browser.OpenURL(url); err != nil {
				panic(err)
			}
		}()
		listen <- true

		open.RunServer(port, *model)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
	openCmd.Flags().StringP("file", "f", "", "input schema filepath")
}
