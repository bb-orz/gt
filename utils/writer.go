package utils

import (
	"os"
	"path"
)

func CreateFile(fileName string) (*os.File, error) {
	var err error
	var file *os.File

	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			dir := path.Dir(fileName)
			_, err := os.Stat(dir)
			if err != nil {
				if os.IsNotExist(err) {
					err := os.MkdirAll(dir, os.ModePerm)
					if err != nil {
						return nil, err
					}
				}
			}

		}
	}

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}
