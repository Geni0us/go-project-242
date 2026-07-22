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
	var human, all, recurcive bool
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText: "hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "recursive",
				Value:       false,
				Usage:       "recursive size of directories (default: false)",
				Aliases:     []string{"r"},
				Destination: &recurcive,
			},
			&cli.BoolFlag{
				Name:        "human",
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit) (default: false)",
				Aliases:     []string{"H"},
				Destination: &human,
			},
			&cli.BoolFlag{
				Name:        "all",
				Value:       false,
				Usage:       "include hidden files and directories (default: false)",
				Aliases:     []string{"a"},
				Destination: &all,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			for _, path := range cmd.Args().Slice() {
				res, err := code.GetPathSize(path, recurcive, human, all)
				if err == nil {
					fmt.Printf("%s\t%s\n", res, path)
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
