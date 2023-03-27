package main

import (
	"encoding/json"
	"fmt"
	"github.com/mb-14/gomarkov"
	"os"
)

const modelFile = "model.json"

func loadModel(name string, order int) (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := os.ReadFile(fmt.Sprintf("%s-%d.json", name, order))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Data for '%s' (order %d) not found. Create training data with: chatfreely gen -c %s", name, order, name)
		}
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func saveModel(chain *gomarkov.Chain, name string) error {
	jsonObj, err := json.Marshal(chain)
	if err != nil {
		return err
	}
	return os.WriteFile(fmt.Sprintf("%s-%d.json", name, chain.Order), jsonObj, 0644)
}
