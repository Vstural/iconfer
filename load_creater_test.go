package iconfer

import (
	"fmt"
	"testing"
)

type TestConfig struct {
	Path  string `json:"path"`
	Value int    `json:"value"`
	User  string `json:"user"`
}

func TestNewConf(t *testing.T) {
	exampleConfig := TestConfig{
		Path:  "123",
		Value: 456,
		User:  "789",
	}

	res, _, err := NewConf[TestConfig]("./config.json", &exampleConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Config)
}
