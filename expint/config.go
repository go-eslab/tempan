package expint

import (
	"encoding/json"
	"os"
)

// Config is a particular configuration of the algorithm.
type Config struct {
	Floorplan string
	HotSpot   struct {
		Config string
		Params string
	}
}

func (c *Config) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(c)

	if err != nil {
		return err
	}

	return nil
}
