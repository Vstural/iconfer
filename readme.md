# iconfer 

read target path file, if read fail, create a example config and return error

usage

```go
func TestCreateFile(t *testing.T) {
	var testPath = "./test"
	f, err := os.Create(testPath)
	if err != nil {
		t.Error("create file fail:", err)
	}
	if f == nil {
		t.Error("got nil file desc")
	}
	f.Close()

	// clear test
	err = os.Remove(testPath)
	if err != nil {
		t.Error(err)
	}
}
```