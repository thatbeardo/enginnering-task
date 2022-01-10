package inject

import "encoding/json"

// Reset function restores the default value of injected variables
// Used during testing
func Reset() {
	Marshal = defaultMarshal
}

// Marshal allows injection of definition that
// marshaller at runtime. Handy during testing
var Marshal = defaultMarshal
var defaultMarshal = json.Marshal
