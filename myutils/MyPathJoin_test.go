package myutils

import "testing"

func TestMyPathJoin(t *testing.T) {
	joinedPath := MyPathJoin("a", "b", "c")
	t.Logf("Joined path: %s", joinedPath)
}
