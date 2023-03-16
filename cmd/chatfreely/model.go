package main

import (
	"encoding/json"
	"github.com/mb-14/gomarkov"
	"os"
)

const modelFile = "model.json"

func loadModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := os.ReadFile(modelFile)
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func saveModel(chain *gomarkov.Chain) error {
	jsonObj, err := json.Marshal(chain)
	if err != nil {
		return err
	}
	return os.WriteFile(modelFile, jsonObj, 0644)
}
