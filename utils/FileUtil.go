package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (fd []byte, err error) {
	fi, err := os.Open(path)
	defer fi.Close()
	fd, err = ioutil.ReadAll(fi)
	//str = string(fd)
	return

}
