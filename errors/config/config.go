package config

import (
	"errors"
	"fmt"
	"io/ioutil"
)

const fileHeader = "APPCONF"

var ErrNoConfigFile = errors.New("no config file at specific location /tmp/myHotelApp/config.txt")

type HeaderError struct {
	FaultyHeader string
}

func (e *HeaderError) Error() string {
	return fmt.Sprintf("Bad header, Provided %s, expected: APPCONF", e.FaultyHeader)
}

func Load() (string, error) {
	dat, err := ioutil.ReadFile("/tmp/myHotelApp/config.txt")

	if err != nil {
		return "", ErrNoConfigFile
	}
	conf := string(dat)
	if conf[0:8] != fileHeader {
		return "", &HeaderError{FaultyHeader: conf[0:8]}
	}
	return conf, nil

	// if err != nil {
	// 	return "", err
	// }

	// conf := string(dat)
	// if conf[0:8] != fileHeader {
	// 	return "", fmt.Errorf("the config file do not begin by accepted header")
	// }
	// return conf, nil
}
