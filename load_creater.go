package iconfer

import (
	"encoding/json"
	"os"
)

// ReadOrCreate try to read content of target path
// create if file not exist and return default content and read error
// return empty and write error if create fail
func ReadOrCreate(path, defaultContent string) (string, bool, error) {
	content, err := ReadFileContent(path)
	if err != nil {
		werr := WriteFileContent(path, defaultContent)
		if werr != nil {
			return "", true, err
		}
		return defaultContent, true, nil
	}
	return content, false, nil
}

// ReadFileContent try to read target path file content
func ReadFileContent(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var data []byte
	buf := make([]byte, 2000)
	for {
		l, err := f.Read(buf)
		if err != nil {
			break
		}
		data = append(data, buf[:l]...)
	}
	return string(data), nil
}

// WriteFileContent write content into target file, create if not exist
func WriteFileContent(path string, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

type ConfigReader[T interface{}] struct {
	ConfigPath string
	Config     T
}

func NewConf[T interface{}](configPath string, exampleConfig *T) (*ConfigReader[T], bool, error) {

	res := &ConfigReader[T]{
		ConfigPath: configPath,
		//Config:     exampleConfig,
	}
	var exampleConfigStr string
	var err error
	if exampleConfig != nil {
		res.Config = *exampleConfig
		// read immediately
		exampleConfigByte, err := json.Marshal(res.Config)
		if err != nil {
			return nil, false, err
		}
		exampleConfigStr = string(exampleConfigByte)
	} else {
		exampleConfigByte, _ := json.Marshal(new(T))
		exampleConfigStr = string(exampleConfigByte)
	}
	configContent, craeted, err := ReadOrCreate(res.ConfigPath, exampleConfigStr)

	if err != nil {
		return nil, craeted, err
	}

	err = json.Unmarshal([]byte(configContent), &res.Config)
	if err != nil {
		return nil, craeted, err
	}

	return res, craeted, nil
}

// watch
