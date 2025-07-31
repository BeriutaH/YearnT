package factory

import "encoding/json"

func EmptyGroup() []byte {
	group, _ := json.Marshal([]string{})
	return group
}
