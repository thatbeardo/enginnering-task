package inject

import "encoding/json"

func Reset() {
	Marshal = defaultMarshal
}

var Marshal = defaultMarshal
var defaultMarshal = json.Marshal
