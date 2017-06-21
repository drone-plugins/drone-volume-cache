package main

import (
	"os"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var version = "1.0.0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "volume cache plugin"
	app.Usage = "valume cache plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{

		//
		// Cache information
		//

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
			Name:  "ttl",
			Usage: "cache ttl in days",
			Value: 30,
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug plugin output",
			EnvVar: "PLUGIN_DEBUG",
		},

		//
		// Build information (for setting defaults)
		//

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
		ttl:      time.Duration(c.Int("ttl")) * 24 * 30,
		storage:  &localCache{},
	}
	if p.file == "" {
		p.file = path.Join(
			c.String("repo-owner"),
			c.String("repo-name"),
			c.String("repo-branch")+".tar",
		)
	}
	if p.fallback == "" {
		p.fallback = path.Join(
			c.String("repo-owner"),
			c.String("repo-name"),
			c.String("repo-branch")+".tar",
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

var x = `
pipeline:
  cache:
    image: plugins/local-cache
    path: /cache
    file: dev.tar.gz
    fallback_to: master.tar.gz
    restore: true
    rebuild: false
    flush: false
    ttl: 30
`
