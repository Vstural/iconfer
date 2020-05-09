package test

import (
	"encoding/json"
	"fmt"
	"github.com/Vstural/iconfer"
	"os"
	"testing"
)

type ExampleConfig struct {
	Name       string            `json:"name"`
	Password   string            `json:"password"`
	Labels     map[string]string `json:"labels"`
	SubConfigs []SubConfig       `json:"sub_configs"`
}

type SubConfig struct {
	Event string `json:"event"`
	ID    int    `json:"id"`
}

var example = ExampleConfig{
	Name:     "Jack",
	Password: "123456",
	Labels: map[string]string{
		"Hunman": "true",
		"Age":    "29",
	},
	SubConfigs: []SubConfig{
		{
			Event: "Nope",
			ID:    1,
		},
		{
			Event: "Yes",
			ID:    2,
		},
	},
}

func TestLoadConfig(t *testing.T) {
	res, err := json.Marshal(example)
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	writePath := fmt.Sprintf("%s\\%s", path, `example_config.json`)
	filecontent, err := iconfer.ReadOrCreate(writePath, string(res))
	if err != nil {
		panic(err)
	}
	fmt.Println(filecontent)
}

func TestGenerateExampleConfig(t *testing.T) {

	res, err := json.Marshal(example)
	if err != nil {
		panic(err)
	}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	writePath := fmt.Sprintf("%s\\%s", path, `example_config.json`)
	fmt.Println(writePath)
	f, err := os.OpenFile(writePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(res)
	if err != nil {
		panic(err)
	}
}
