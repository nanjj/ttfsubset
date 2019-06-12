package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"
)

var (
	RootCmd *cobra.Command
)

func init() {
	RootCmd = &cobra.Command{
		Use:  "ttfsubset <original ttf>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			infile := args[0]
			flags := cmd.Flags()
			output, err := flags.GetString("output")
			if err != nil {
				return
			}
			runes, err := flags.GetString("runes")
			if err != nil {
				return
			}
			if runes == "" {
				runes, err = GetAllRunes(infile)
				if err != nil {
					return
				}
			}
			in, err := ioutil.ReadFile(infile)
			if err != nil {
				return
			}
			data := gofpdf.UTF8CutFont(in, runes)
			if output != "" {
				err = ioutil.WriteFile(output, data, 0644)
				if err != nil {
					return
				}
			} else {
				fmt.Println(FormatBytes(data))
			}
			return
		},
	}
	flags := RootCmd.Flags()
	flags.StringP("output", "o", "", "output file")
	flags.StringP("runes", "r", "", "runes to keep, default to keep all")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.Println(err)
	}
}
