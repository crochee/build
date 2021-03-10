package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.NewApp()
	app := cli.NewApp()
	app.Name = "build"
	app.Version = Version
	app.Usage = "build project"

	app.Commands = []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build project",
			Action:  run,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "url",
					Aliases: []string{"u"},
					Usage:   "git url",
					EnvVars: []string{"GIT_URL"},
				},
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "cid file",
					EnvVars: []string{"CID_FILE"},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Print(err)
	}
}

func run(ctx *cli.Context) error {
	uri := ctx.String("url")
	if uri == "" {
		return errors.New("url is nil")
	}
	if err := Exec("git", "clone", uri); err != nil {
		return err
	}
	index := strings.LastIndex(uri, "/")
	project := strings.TrimSuffix(uri[index:], path.Ext(uri[index:]))
	file := ctx.String("file")
	if file == "" {
		file = fmt.Sprintf("%s/.cid/%s/%s.yml", project, project)
	} else {
		file = fmt.Sprintf("%s/.cid/%s", project, file)
	}
	if err := Exec("cd", project); err != nil {
		return err
	}
	return Shell(file)
}
