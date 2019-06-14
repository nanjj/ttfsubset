package main

import (
	"io/ioutil"
	"os"

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
			runes := ReadRunesFromStdin()
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
			_, err = os.Stdout.Write(data)
			if err != nil {
				return
			}
			return
		},
	}
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.Println(err)
	}
}
