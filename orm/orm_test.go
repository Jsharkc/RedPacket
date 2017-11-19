package orm

import (
	"testing"
)

func TestInitOrm(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Connecting mysql err:", err)
		}
	}()

	InitOrm()
}
