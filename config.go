package hotspot

// Config is a configuration of the HotSpot tool.
type Config struct {
	// A floorplan of an electronic system.
	Floorplan string
	// A native configuration file of the HotSpot tool (hotspot.config).
	Configuration string
	// A line of parameters overwriting the parameters in the native
	// configuration file. The names and values of the parameters should be
	// given in pairs separated by whitespaces. An example of a valid line of
	// parameters is "t_chip 0.00015 k_chip 100.0".
	Parameters string
}
