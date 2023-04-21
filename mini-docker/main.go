package main

import (
	"basicGo/mini-docker/container"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `mini-docker is a docker test`

func main() {
	app := cli.NewApp()
	app.Name = "mini-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		InitCommand,
		RunCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	log.Println("#### finished ####")
}

var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process in container",
	Action: func(context *cli.Context) error {
		log.Info("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command is %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err

	},
}

var RunCommand = cli.Command{
	Name:  "run",
	Usage: `run container`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},

	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}

		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
