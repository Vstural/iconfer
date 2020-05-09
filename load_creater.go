package iconfer

import (
	"os"
)

func ReadOrCreate(path, defaultContent string) (string, error) {
	content, err := ReadFileContent(path)
	if err != nil {
		_ = WriteFileContent(path, defaultContent)
		return "", err
	}
	return content, nil
}

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
