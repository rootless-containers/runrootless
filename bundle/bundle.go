package bundle

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func Transform(newBundle, oldBundle string) error {
	spec, err := readSpec(oldBundle)
	if err != nil {
		return err
	}
	if err := transformSpec(spec, oldBundle); err != nil {
		return err
	}
	return writeSpec(newBundle, spec)
}

func readSpec(bundle string) (*specs.Spec, error) {
	f := filepath.Join(bundle, "config.json")
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var spec specs.Spec
	err = json.Unmarshal(b, &spec)
	return &spec, err
}

func writeSpec(bundle string, spec *specs.Spec) error {
	if err := os.MkdirAll(bundle, 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(spec, "", "\t")
	if err != nil {
		return err
	}
	f := filepath.Join(bundle, "config.json")
	return ioutil.WriteFile(f, data, 0666)
}
