package durationslice

import (
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	ds, err := Process(time.Duration((75*time.Minute)+(65*time.Second)), "h m s")
	if err != nil {
		t.Errorf(err.Error())
	}
	ev := 3
	if len(ds) != ev {
		t.Errorf("Expected %d time values but found %d", ev, len(ds))
	} else {
		eh := int64(1)
		em := int64(16)
		es := int64(5)
		if ds[0] != eh {
			t.Errorf("Expected %d hours but got %d", eh, ds[0])
		}
		if ds[1] != em {
			t.Errorf("Expected %d minutes but got %d", em, ds[1])
		}
		if ds[2] != es {
			t.Errorf("Expected %d seconds but got %d", es, ds[1])
		}
	}
}
