package main

import (
	"github.com/mb-14/gomarkov"
	"github.com/writeas/go-writeas/v2"
	"log"
	"os"
	"strings"
	"sync"
)

func fetchBlogPosts(alias string) ([]writeas.Post, error) {
	log.Printf("Fetching blog posts from '%s'...", alias)
	c := writeas.NewClient()
	c.SetApplicationKey(os.Getenv("APP_KEY"))
	posts, err := c.GetCollectionPosts(alias)
	if err != nil {
		return nil, err
	}
	return *posts, err
}

func buildModel(alias string) (*gomarkov.Chain, error) {
	posts, err := fetchBlogPosts(alias)
	if err != nil {
		return nil, err
	}
	chain := gomarkov.NewChain(1)
	var wg sync.WaitGroup
	wg.Add(len(posts))
	log.Printf("Adding %d posts to markov chain...", len(posts))
	for _, storyID := range posts {
		go func(p writeas.Post) {
			defer wg.Done()
			log.Printf("\"%s\"", p.Title)
			chain.Add(strings.Split(p.Content, " "))
		}(storyID)
	}
	wg.Wait()
	return chain, nil
}
