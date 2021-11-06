package magicimage

import (
	"errors"
	"os"
)

func RenameFolderAndFile(old_path, new_path string) error {
	// check directory or file exists
	if _, err := os.Stat(old_path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	// rename path
	err := os.Rename(old_path, new_path)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFolderAndFile(path string) error {
	// check directory or file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Remove all the directories and files
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
