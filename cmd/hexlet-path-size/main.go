package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Action: func(_ context.Context, cmd *cli.Command) error {
			for _, path := range cmd.Args().Slice() {
				res, err := code.GetPathSize(path, false, true, false)
				if err == nil {
					fmt.Println(res)
				} else {
					log.Fatal(err)
				}
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
