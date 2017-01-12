package durationslice

import (
	"errors"
	"strings"
	"time"
)

var (
	UNITS = []string{"y", "mon", "w", "d", "h", "m", "s", "milli"}
)

const (
	YEAR_IN_NANO  = time.Duration(time.Hour * 8760)
	MONTH_IN_NANO = time.Duration(time.Hour * 730) // used 730 instead of 730.001
	WEEK_IN_NANO  = time.Duration(time.Hour * 168)
	DAY_IN_NANO   = time.Duration(time.Hour * 24)
)

// Process takes a duration and a string defining time units.
// It returns a slice with values corresponding to the units defined in the string.
func Process(d time.Duration, s string) ([]int64, error) {
	u := strings.Fields(s)
	_, err := checkUnits(u)
	if err != nil {
		return nil, err
	}

	a := make([]int64, len(u))
	for i := 0; i < len(u); i++ {
		v := u[i]
		uv, b := extractUnit(d, v)
		a[i] = uv
		d = time.Duration(b)
	}
	return a, nil
}

func extractUnit(d time.Duration, u string) (int64, int64) {
	switch u {
	case "y":
		y := d.Nanoseconds() / YEAR_IN_NANO.Nanoseconds()
		b := d.Nanoseconds() - (y * YEAR_IN_NANO.Nanoseconds())
		return y, b
	case "mon":
		m := d.Nanoseconds() / MONTH_IN_NANO.Nanoseconds()
		b := d.Nanoseconds() - (m * MONTH_IN_NANO.Nanoseconds())
		return m, b
	case "w":
		w := d.Nanoseconds() / WEEK_IN_NANO.Nanoseconds()
		b := d.Nanoseconds() - (w * WEEK_IN_NANO.Nanoseconds())
		return w, b
	case "d":
		day := d.Nanoseconds() / DAY_IN_NANO.Nanoseconds()
		b := d.Nanoseconds() - (day * DAY_IN_NANO.Nanoseconds())
		return day, b
	case "h":
		h := d.Nanoseconds() / time.Hour.Nanoseconds()
		b := d.Nanoseconds() - (h * time.Hour.Nanoseconds())
		return h, b
	case "m":
		min := d.Nanoseconds() / time.Minute.Nanoseconds()
		b := d.Nanoseconds() - (min * time.Minute.Nanoseconds())
		return min, b
	case "s":
		s := d.Nanoseconds() / time.Second.Nanoseconds()
		b := d.Nanoseconds() - (s * time.Second.Nanoseconds())
		return s, b
	case "milli":
		milli := d.Nanoseconds() / time.Millisecond.Nanoseconds()
		b := d.Nanoseconds() - (milli * time.Millisecond.Nanoseconds())
		return milli, b
	}
	return 0, 0
}

func checkUnits(u []string) (bool, error) {
	lp := -1
	for _, v := range u {
		f := false
		for j, w := range UNITS {
			if v == w {
				if j > lp {
					lp = j
					f = true
					break
				} else {
					return false, errors.New("Wrong order of time units, must be desending.")
				}
			}
		}
		if !f {
			return false, errors.New("Invalid time unit used.")
		}
	}
	return true, nil
}
