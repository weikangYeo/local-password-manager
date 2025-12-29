package utils

import (
	"os"
	"path/filepath"
)

func toAppStoragePath(path string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".local-pwd-manager", path)
}

func IsFileExists(path string) bool {
	_, err := os.Stat(toAppStoragePath(path))
	return !os.IsNotExist(err)
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(toAppStoragePath(path))
}

func WriteFile(path string, data []byte) error {
	fullPath := toAppStoragePath(path)
	if err := os.MkdirAll(filepath.Dir(fullPath), 0700); err != nil {
		return err
	}
	return os.WriteFile(toAppStoragePath(path), data, 0600)
}

// func AppendFile(path string, data []byte) error {
// 	return os.AppendFile(toAppStoragePath(path), data, 0644)
// }
