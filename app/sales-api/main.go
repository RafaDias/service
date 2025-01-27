package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"log"
	"os"
	"time"
	"github.com/pkg/errors"
)

func main() {
	log := log.New(os.Stdout, "SALES : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		log.Println("main: error", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	// ==========================
	// Configuration

	var cfg struct{
		conf.Version
		APIHost string `conf:"default:0.0.0.0:3000"`
		DebugHost string `conf:"default:0.0.0.0:4000"`
		ReadTimeout time.Duration `conf:"default:5s"`
		WriteTimeout time.Duration `conf:"default:5s"`
		ShutdownTimeout time.Duration `conf:"default:5s"`
	}
	cfg.Version.Desc = "copyright information here"
	if err := conf.Parse(os.Args[1:], "SALES", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("SALES", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString("SALES", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}
	return nil
}
