package main

import (
	"fmt"
	"github.com/mb-14/gomarkov"
	"github.com/urfave/cli/v2"
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
		Name:     "order",
		Usage:    "Alias of the WriteFreely collection to train on",
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
	order := ctx.Int("order")
	if order == 0 {
		order = 1
	}
	chain, err := buildModel(alias, order)
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
	err = generateBlogPost(chain)
	if err != nil {
		return err
	}
	return nil
}

func generateBlogPost(chain *gomarkov.Chain) error {
	order := chain.Order
	if order == 0 {
		order = 1
	}
	log.Printf("Generating post. Order %d", order)
	tokens := []string{gomarkov.StartToken}
	for i := 1; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, err := chain.Generate(tokens[(len(tokens) - order):])
		if err != nil {
			return err
		}
		if next != gomarkov.EndToken {
			fmt.Print(next + " ")
		}
		tokens = append(tokens, next)
	}
	//fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
	return nil
}
