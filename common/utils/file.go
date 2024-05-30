package utils

import (
	"io"
	"os"

	"apple/common/log"
)

var File = &_File{}

type _File struct {
}

func (_File) ReadWithIOUtil(path string) []byte {
	//start := time.Now()
	fi, err := os.Open(path)
	defer func() {
		if fi != nil {
			fi.Close()
		}
	}()
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	fd, err := io.ReadAll(fi)
	if err != nil {
		return nil
	}
	return fd
}
