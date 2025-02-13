package rpc

import "fmt"

const FetchProgressId = `fetch-progress`

func AppendID(base string, parts ...string) string {
	var out = base
	for _, part := range parts {
		out += fmt.Sprintf("-%s", part)
	}
	return out
}
