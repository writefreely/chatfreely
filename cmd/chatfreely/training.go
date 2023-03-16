package main

import (
	"github.com/mb-14/gomarkov"
	"github.com/writeas/go-writeas/v3"
	"log"
	"os"
	"strings"
	"sync"
)

func fetchBlogPosts(alias string) ([]writeas.Post, error) {
	log.Printf("Fetching blog posts from '%s'...", alias)
	c := writeas.NewClient()
	c.SetApplicationKey(os.Getenv("APP_KEY"))
	var posts *[]writeas.Post
	var allPosts []writeas.Post
	var err error
	i := 1
	for i == 1 || len(*posts) != 0 {
		posts, err = c.GetCollectionPosts(alias, i)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, *posts...)
		i++
	}
	return allPosts, err
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
