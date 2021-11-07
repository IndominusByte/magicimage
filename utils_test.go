package magicimage

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenameFolderAndFile(t *testing.T) {
	// path not exists
	t.Run("path not exists", func(t *testing.T) {
		result := RenameFolderAndFile("test", "test")
		assert.NotNil(t, result)
	})
	// path with sub not exists
	t.Run("path with sub not exists", func(t *testing.T) {
		result := RenameFolderAndFile("test/testing", "test/testing")
		assert.NotNil(t, result)
	})

	os.MkdirAll("test/testing", os.ModePerm)
	// success rename folder
	t.Run("success rename folder", func(t *testing.T) {
		result := RenameFolderAndFile("test", "test-change")
		assert.Nil(t, result)
		assert.DirExists(t, "test-change")
	})
	// success rename folder with sub
	t.Run("success rename folder with sub", func(t *testing.T) {
		result := RenameFolderAndFile("test-change/testing", "test-change/testing-change")
		assert.Nil(t, result)
		assert.DirExists(t, "test-change/testing-change")
	})

	f, e := os.Create("test-change/filename.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello")

	f, e = os.Create("test-change/testing-change/filename.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello")

	// success rename file
	t.Run("success rename file", func(t *testing.T) {
		result := RenameFolderAndFile("test-change/filename.txt", "test-change/filename-change.txt")
		assert.Nil(t, result)
		assert.FileExists(t, "test-change/filename-change.txt")
	})
	// success rename file with sub
	t.Run("success rename file with sub", func(t *testing.T) {
		result := RenameFolderAndFile("test-change/testing-change/filename.txt", "test-change/testing-change/filename-change.txt")
		assert.Nil(t, result)
		assert.FileExists(t, "test-change/testing-change/filename-change.txt")
	})

	os.RemoveAll("test-change")
}

func TestDeleteFolderAndFile(t *testing.T) {
	// path not exists
	t.Run("path not exists", func(t *testing.T) {
		result := DeleteFolderAndFile("test")
		assert.NotNil(t, result)
	})
	// path with sub not exists
	t.Run("path with sub not exists", func(t *testing.T) {
		result := DeleteFolderAndFile("test/testing")
		assert.NotNil(t, result)
	})

	os.MkdirAll("test/testing", os.ModePerm)
	f, e := os.Create("test/filename.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello")

	f, e = os.Create("test/testing/filename.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	fmt.Fprintln(f, "Hello")

	// delete file
	t.Run("delete file", func(t *testing.T) {
		result := DeleteFolderAndFile("test/testing/filename.txt")
		assert.Nil(t, result)
		assert.NoFileExists(t, "test/testing/filename.txt")
	})
	// delete folder
	t.Run("delete folder", func(t *testing.T) {
		result := DeleteFolderAndFile("test")
		assert.Nil(t, result)
		assert.NoDirExists(t, "test")
	})
}
