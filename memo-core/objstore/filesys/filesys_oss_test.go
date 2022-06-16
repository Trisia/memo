package filesys

import (
	"os"
	"testing"
)

func TestFileSystemBaseOSS_Get(t *testing.T) {
	oss := NewFileSystemBaseOSS()
	file, err := os.Open("filesys_oss.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	hashName, _, err := oss.Put("77", "filesys_oss.go", file)
	if err != nil {
		t.Fatal(err)
	}
	_, err = oss.Get("77", hashName, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileSystemBaseOSS_Del(t *testing.T) {
	oss := NewFileSystemBaseOSS()
	file, err := os.Open("filesys_oss.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	hashName, _, err := oss.Put("77", "filesys_oss.go", file)
	if err != nil {
		t.Fatal(err)
	}
	_, err = oss.Get("77", hashName, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}

	err = oss.Del("77", hashName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileSystemBaseOSS_DelDoc(t *testing.T) {
	oss := NewFileSystemBaseOSS()
	file, err := os.Open("filesys_oss.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	hashName, _, err := oss.Put("77", "filesys_oss.go", file)
	if err != nil {
		t.Fatal(err)
	}
	_, err = oss.Get("77", hashName, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}

	err = oss.DelDoc("77")
	if err != nil {
		t.Fatal(err)
	}
}
