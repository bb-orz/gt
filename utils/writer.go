package utils

import (
	"io"
	"os"
	"path"
)

func CreateFile(fileName string) (io.Writer, error) {
	var err error
	var file io.Writer

	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			dir := path.Dir(fileName)
			dirInfo, err := os.Stat(dir)
			if err != nil {
				if os.IsNotExist(err) {
					err := os.MkdirAll(dir, os.ModePerm)
					if err != nil {
						return nil, err
					}
					dirInfo, err = os.Stat(dir)
				}
			}
			if dirInfo.IsDir() {
				file, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
				if err != nil {
					return nil, err
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
