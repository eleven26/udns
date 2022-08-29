package main

import (
	"errors"
	"os"
	"path/filepath"

	fs "github.com/eleven26/go-filesystem"
)

func DefaultConfigPath() (string, error) {
	for _, p := range paths() {
		if exists, _ := fs.Exists(filepath.Join(p, File)); exists {
			return filepath.Join(p, File), nil
		}
	}

	return "", errors.New("")
}

func paths() []string {
	var res []string
	res = append(res, "/etc")

	home, err := os.UserHomeDir()
	if err == nil {
		res = append(res, home)
	}

	return res
}
