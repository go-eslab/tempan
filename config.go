package hotspot

// Config is a configuration of the HotSpot model.
type Config struct {
	// A floorplan of a multiprocessor system.
	Floorplan string
	// A native configuration file of the HotSpot model (hotspot.config).
	Configuration string
}
