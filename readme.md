# iconfer 

`go get github.com/Vstural/iconfer`

read target path file, if read fail, create a example config and return error

usage 
```go
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

var path = "./hello.json"
func TestLoadConfig(t *testing.T) {

	filecontent, err := iconfer.ReadOrCreate(writePath, string(res))
	if err ...
	fmt.Println(filecontent)
}
```