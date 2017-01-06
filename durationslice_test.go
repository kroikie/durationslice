package durationslice

import (
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	ds, err := Process(time.Duration((48*time.Hour)+(10*time.Second)), "h s")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(ds) != 2 {
		t.Errorf("Expected 3 time values but found %d", len(ds))
	}
	if ds[0] != 48 {
		t.Errorf("Expected 48 hours but got %d", ds[0])
	}
	//if ds[1] != 0 {
	//	t.Errorf("Expected 0 minutes but got %d", ds[1])
	//}
	if ds[1] != 10 {
		t.Errorf("Expected 10 seconds but got %d", ds[1])
	}
}
