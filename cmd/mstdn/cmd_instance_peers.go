package main

import (
	"context"
	"fmt"

	"github.com/misaka00251/go-mastodon"
	"github.com/urfave/cli"
)

func cmdInstancePeers(c *cli.Context) error {
	client := c.App.Metadata["client"].(*mastodon.Client)
	peers, err := client.GetInstancePeers(context.Background())
	if err != nil {
		return err
	}
	for _, peer := range peers {
		fmt.Fprintln(c.App.Writer, peer)
	}
	return nil
}
