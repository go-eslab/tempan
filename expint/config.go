package expint

import (
	"encoding/json"
	"errors"
	"os"
)

// Config captures the configuration of a particular problem.
type Config struct {
	// The path to a file specifying the floorplan of the platform to analyze.
	Floorplan string

	// Options related to the HotSpot model.
	HotSpot struct {
		// The path to a native configuration file (hotspot.config).
		Config string
		// A line of parameters overwriting the parameters in the above file.
		Params string
	}

	// A sampling interval to be used for temperature analysis. It is the time
	// between two successive samples of power in power profiles and two
	// successive samples of temperature in temperature profiles. In the
	// package, it is referred to as dt.
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
