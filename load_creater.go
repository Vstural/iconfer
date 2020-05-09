package iconfer

import (
	"os"
)

// ReadOrCreate try to read content of target path
// create if file not exist and return default content and read error
// return empty and write error if create fail
func ReadOrCreate(path, defaultContent string) (string, error) {
	content, err := ReadFileContent(path)
	if err != nil {
		werr := WriteFileContent(path, defaultContent)
		if werr != nil {
			return "", err
		}
		return defaultContent, err
	}
	return content, nil
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
