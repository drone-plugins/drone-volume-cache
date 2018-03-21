package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "volume cache plugin"
	app.Usage = "volume cache plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "path",
			Usage:  "path",
			Value:  "/cache",
			EnvVar: "PLUGIN_PATH",
		},
		cli.StringFlag{
			Name:   "file",
			Usage:  "file",
			EnvVar: "PLUGIN_FILE",
		},
		cli.StringFlag{
			Name:   "fallback-to",
			Usage:  "fallback-to",
			EnvVar: "PLUGIN_FALLBACK_TO",
		},
		cli.StringSliceFlag{
			Name:   "mount",
			Usage:  "cache directories",
			EnvVar: "PLUGIN_MOUNT",
		},
		cli.BoolFlag{
			Name:   "rebuild",
			Usage:  "rebuild the cache directories",
			EnvVar: "PLUGIN_REBUILD",
		},
		cli.BoolFlag{
			Name:   "restore",
			Usage:  "restore the cache directories",
			EnvVar: "PLUGIN_RESTORE",
		},
		cli.BoolFlag{
			Name:   "flush",
			Usage:  "flush the cache directories",
			EnvVar: "PLUGIN_FLUSH",
		},
		cli.IntFlag{
			Name:   "ttl",
			Usage:  "cache ttl in days",
			Value:  30,
			EnvVar: "PLUGIN_TTL",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug plugin output",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.StringFlag{
			Name:   "repo-owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo-name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit-branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	p := &plugin{
		mount:    c.StringSlice("mount"),
		path:     c.String("path"),
		file:     c.String("file"),
		fallback: c.String("fallback-to"),
		rebuild:  c.Bool("rebuild"),
		restore:  c.Bool("restore"),
		flush:    c.Bool("flush"),
		ttl:      time.Duration(c.Int("ttl")) * 24 * time.Hour,
		storage:  &localCache{},
	}

	if p.file == "" {
		p.file = path.Join(
			c.String("repo-owner"),
			c.String("repo-name"),
			c.String("commit-branch")+".tar",
		)
	}

	if p.fallback == "" {
		p.fallback = path.Join(
			c.String("repo-owner"),
			c.String("repo-name"),
			c.String("commit-branch")+".tar",
		)
	}

	if !path.IsAbs(p.file) {
		p.file = path.Join(p.path, p.file)
	}

	if !path.IsAbs(p.fallback) {
		p.fallback = path.Join(p.path, p.fallback)
	}

	return p.Exec()
}
