package main

import (
	"encoding/json"
	"fmt"
	"github.com/mb-14/gomarkov"
	"os"
)

const modelFile = "model.json"

func loadModel(name string) (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := os.ReadFile(name + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Data for '%s' not found. Create training data with: chatfreely gen -c %s", name, name)
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
	return os.WriteFile(name+".json", jsonObj, 0644)
}
