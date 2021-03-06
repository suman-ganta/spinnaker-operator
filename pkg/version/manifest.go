package version

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	operatorHomePath = "OPERATOR_HOME"
	manifestFile     = "MANIFEST"
)

var manifest = make(map[string]string)

func read() error {
	// read MANIFEST file, that contains Operator Version Information
	// absolute path: $(OPERATOR_HOME)/MANIFEST
	path := os.Getenv(operatorHomePath)
	body, err := ioutil.ReadFile(path + "/" + manifestFile)

	if err != nil {
		return err
	}

	raws := strings.Split(string(body), "\n")

	for _, raw := range raws {
		if raw != "" {
			values := strings.Split(raw, "=")
			if len(values) == 2 {
				manifest[values[0]] = values[1]
			}
		}
	}
	return nil
}

func GetManifestValue(key string) (string, error) {

	// initialize manifest lazily
	if len(manifest) == 0 {
		err := read()
		if err != nil {
			return "", err
		}
	}

	val, ok := manifest[key]
	if ok {
		return val, nil
	}
	return "", fmt.Errorf("key %v not found in manifest", key)
}
