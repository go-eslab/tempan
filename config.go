package hotspot

// Config is a configuration of HotSpot.
type Config struct {
	// A floorplan of a multiprocessor system.
	Floorplan string
	// A native configuration file of HotSpot (hotspot.config).
	Configuration string
}
