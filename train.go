package chatfreely

import (
	"github.com/mb-14/gomarkov"
	"github.com/writeas/go-writeas/v3"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

const writeasURL = "https://write.as"

func fetchBlogPosts(alias, instance string) ([]writeas.Post, error) {
	// Normalize and correctly parse instance
	if instance == "" {
		instance = writeasURL
	}
	if !strings.HasPrefix(instance, "http") {
		instance = "https://" + instance
	}
	host, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	host.Path = "/api"

	// Set up WriteFreely API client
	log.Printf("Fetching blog posts from '%s' on %s (via %s)...", alias, host.Host, host.String())
	var c *writeas.Client
	if instance == writeasURL {
		c = writeas.NewClient()
		// Write.as requires an application key to get around rate limiting
		c.SetApplicationKey(os.Getenv("WRITEAS_APP_KEY"))
	} else {
		c = writeas.NewClientWith(writeas.Config{
			URL: host.String(),
		})
	}

	var posts *[]writeas.Post
	var allPosts []writeas.Post
	i := 1
	for i == 1 || len(*posts) != 0 {
		log.Printf("Page %d...", i)
		posts, err = c.GetCollectionPosts(alias, i)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, *posts...)
		i++
	}
	return allPosts, err
}

// BuildModel creates a model with the given order for the given collection alias.
func BuildModel(alias, instance string, order int) (*gomarkov.Chain, error) {
	posts, err := fetchBlogPosts(alias, instance)
	if err != nil {
		return nil, err
	}
	chain := gomarkov.NewChain(order)
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
