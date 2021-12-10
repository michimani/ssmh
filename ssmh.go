package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"ssmhistory/cmd"
	"ssmhistory/internal"
	"ssmhistory/types"
)

var (
	inputN   *int    = flag.Int("n", 0, "max items")
	inputR   *string = flag.String("r", "", "region")
	version  string
	revision string
)

func usage() {
	format := `
                      _     _     _
    ___ ___ _ __ ___ | |__ (_)___| |_ ___  _ __ _   _
   / __/ __| '_ ' _ \| '_ \| / __| __/ _ \| '__| | | |
   \__ \__ \ | | | | | | | | \__ \ || (_) | |  | |_| |
   |___/___/_| |_| |_|_| |_|_|___/\__\___/|_|   \__, |
                                                |___/
   Version: %s-%s

Usage:
  ssmh [flags] [values]
Flags:
  -n (integer)
    Specifies the number of session histories to retrieve,
    a value greater than 0.

  -r (string)
    Specify the region to retrieve. If not specified, the
    region set in the environment variable "AWS_DEFAULT_REGION"
    will be the target region.

Exapmple:
  ssmh -n 10 -r ap-northeast-1
`
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, version, revision))
}

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run())
}

func run() int {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if inputR != nil && *inputR != "" {
		region = *inputR
	}

	in := &types.NewClientInput{
		Context: context.Background(),
		Region:  region,
	}

	client, err := internal.NewClient(in)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	p := types.HistoryParams{}

	maxResults := types.MaxResults(*inputN)
	if maxResults.Valid() {
		p.MaxResults = &maxResults
	}

	if err := cmd.ListSessionHistory(context.Background(), client, &p); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	return 0
}
