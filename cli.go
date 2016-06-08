package main

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"gopkg.in/amz.v3/aws"
	"gopkg.in/codegangsta/cli.v2" // imports as package "cli"
	"os"
	"strings"
	"time"
)

const (
	flagBucket = "bucket"
	flagPath   = "path"
	flagKey    = "key"
	flagSecret = "secret"
	flagRegion = "region"
)

func cliDef() {
	app := cli.NewApp()
	app.Name = "s3size"
	app.Version = "0.1"
	app.Authors = []cli.Author{cli.Author{Name: "Enrique Bris", Email: "enrique@bris.io"}}
	app.Usage = "S3 bucket size calculator"
	app.Setup()
	app.Action = cliMainAction

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  flagBucket,
			Value: "",
			Usage: "bucket name",
		},

		cli.StringFlag{
			Name:  flagPath,
			Value: "",
			Usage: "path",
		},

		cli.StringFlag{
			Name:  flagKey,
			Value: "",
			Usage: "S3 access key",
		},

		cli.StringFlag{
			Name:  flagSecret,
			Value: "",
			Usage: "S3 secret key",
		},

		cli.StringFlag{
			Name:  flagRegion,
			Value: "us-east-1",
			Usage: "region",
		},
	}

	app.Run(os.Args)
}

func getRegion(regionName string) (region aws.Region, err error) {
	region, ok := aws.Regions[regionName]
	if !ok {
		err = errors.New("Unknown region. Check http://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region")
	}

	return
}

func validateBucket(bucket string) error {
	if strings.Trim(bucket, " ") == "" {
		return errors.New("Invalid bucket name (no name)")
	}
	return nil
}

func cliError(err error) {
	fmt.Printf("%s\nTry \"help\" to get all the options.\n", err.Error())
}

func cliMainAction(c *cli.Context) (err error) {

	bucket := c.String(flagBucket)
	err = validateBucket(bucket)
	if err != nil {
		cliError(err)
		return
	}

	path := c.String(flagPath)

	key := c.String(flagKey)
	if strings.Trim(key, " ") == "" {
		err = errors.New("Invalid key (empty key)")
		cliError(err)
		return
	}

	secret := c.String(flagSecret)
	if strings.Trim(secret, " ") == "" {
		err = errors.New("Invalid secret (empty secret)")
		cliError(err)
		return
	}

	region, err := getRegion(c.String(flagRegion))
	if err != nil {
		cliError(err)
		return
	}

	// initialize spinner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "work in progress "
	s.Color("green")
	s.Start()

	// try to calculate the size
	err = getBucketList(bucket, key, secret, region, path, "", "", 0)

	// stop spinner
	s.Stop()

	if err != nil {
		err = errors.New("AWS message: " + err.Error())
		cliError(err)
		return
	}

	return nil
}
