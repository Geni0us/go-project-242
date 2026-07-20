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
	var human bool
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit)",
				Aliases:     []string{"H"},
				Destination: &human,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			for _, path := range cmd.Args().Slice() {
				res, err := code.GetPathSize(path, false, human, false)
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
