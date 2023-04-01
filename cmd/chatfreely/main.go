package main

import (
	"github.com/urfave/cli/v2"
	"github.com/writefreely/chatfreely"
	"log"
	"os"
)

var allFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "alias",
		Usage:    "Alias of the WriteFreely collection to train on",
		Required: true,
		Aliases:  []string{"c"},
	},
	&cli.StringFlag{
		Name:     "instance",
		Usage:    "WriteFreely instance to train on",
		Required: false,
		Aliases:  []string{"i"},
	},
	&cli.StringFlag{
		Name:     "order",
		Usage:    "Markov chain order",
		Required: false,
		Aliases:  []string{"o"},
	},
}

func main() {
	app := &cli.App{
		Name:   "ChatFreely",
		Usage:  "Generative \"AI\" that learns from WriteFreely blogs.",
		Action: cmdTrain,
		Commands: []*cli.Command{
			{
				Name:   "train",
				Usage:  "Train the markov chain.",
				Flags:  allFlags,
				Action: cmdTrain,
			},
			{
				Name:    "generate",
				Aliases: []string{"gen"},
				Usage:   "Generate a blog post from training data.",
				Flags:   allFlags,
				Action:  cmdGenerate,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func cmdTrain(ctx *cli.Context) error {
	alias := ctx.String("alias")
	instance := ctx.String("instance")
	order := ctx.Int("order")
	if order == 0 {
		order = 1
	}
	chain, err := chatfreely.BuildModel(alias, instance, order)
	if err != nil {
		return err
	}
	return saveModel(chain, alias)
}

func cmdGenerate(ctx *cli.Context) error {
	alias := ctx.String("alias")
	order := ctx.Int("order")
	if order == 0 {
		order = 1
	}
	chain, err := loadModel(alias, order)
	if err != nil {
		return err
	}
	err = chatfreely.PrintBlogPost(chain)
	if err != nil {
		return err
	}
	return nil
}
