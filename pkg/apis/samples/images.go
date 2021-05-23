package samples

import (
	"fmt"
	"sort"
)

type SinkImages struct {
	HTTP string
}

// Validate returns an error if any image is not set.
func (i SinkImages) Validate() error {
	var unset []string
	for _, f := range []struct {
		v, name string
	}{
		{i.HTTP, "sink-http"},
	} {
		if f.v == "" {
			unset = append(unset, f.name)
		}
	}
	if len(unset) > 0 {
		sort.Strings(unset)
		return fmt.Errorf("found unset image flags: %s", unset)
	}
	return nil
}
