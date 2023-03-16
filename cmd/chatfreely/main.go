package main

import (
	"flag"
	"fmt"
	"github.com/mb-14/gomarkov"
	"log"
	"os"
)

func main() {
	train := flag.String("train", "", "Train the markov chain on the given username")
	prompt := flag.String("in", "", "Prompt to generate post")
	flag.Parse()

	if *train != "" {
		err := cmdTrain(*train)
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	chain, err := loadModel()
	if err != nil {
		log.Fatalln(err)
	}
	err = generateBlogPost(chain, *prompt)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}

func cmdTrain(alias string) error {
	chain, err := buildModel(alias)
	if err != nil {
		return err
	}
	return saveModel(chain)
}

func generateBlogPost(chain *gomarkov.Chain, prompt string) error {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, err := chain.Generate(tokens[(len(tokens) - 1):]) //strings.Split(prompt, " "))
		if err != nil {
			return err
		}
		if next != gomarkov.EndToken {
			fmt.Print(next + " ")
		}
		tokens = append(tokens, next)
	}
	/*
		next, err := chain.Generate(strings.Split(prompt, " "))
		if err != nil {
			return err
		}
		fmt.Print(next + " ")
	*/
	//fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
	return nil
}
