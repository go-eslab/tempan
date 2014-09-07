package expint

import (
	"encoding/json"
	"errors"
	"os"
)

// Config is a particular configuration of the algorithm.
type Config struct {
	Floorplan string
	HotSpot   struct {
		Config string
		Params string
	}
	TimeStep float64
}

func (c *Config) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)

	if err = dec.Decode(c); err != nil {
		return err
	}

	if err = c.validate(); err != nil {
		return err
	}

	return nil
}

func (c *Config) validate() error {
	if c.TimeStep <= 0 {
		return errors.New("the time step is invalid")
	}

	return nil
}
