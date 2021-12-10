package scenario

import "testing"

type Then struct {
	description string
	fn          func(t *testing.T)
}
