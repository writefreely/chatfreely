package chatfreely

import (
	"fmt"
	"github.com/mb-14/gomarkov"
	"log"
	"strings"
)

// PrintBlogPost prints out a post from the given training data.
func PrintBlogPost(chain *gomarkov.Chain) error {
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

// GenerateBlogPost generates and returns a post from the given training data.
func GenerateBlogPost(chain *gomarkov.Chain) (string, error) {
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
			return "", err
		}
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[order:len(tokens)-1], " "), nil
}
