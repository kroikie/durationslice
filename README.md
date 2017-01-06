# Duration Slice
Duration Slice is a Go library that generates a slice representation
of a duration in units specified.

The builtin Duration provides the ability to retrieve a duration
in different units eg: hours, minutes, seconds etc. However partial
representations are not available, Duration Slice provides this partial
representation.

If you had a Duration of 120 seconds and wanted to represent it in
minutes and seconds you would have to do the arithmetic yourself. With
Duration Slice you can provide the duration and a string representing
what units you want the duration divided into and it will return a slice
with the corresponding unit values.

## Example
If you had a duration that was 48 hours and 10 seconds long and you wanted
to represent that as days, minutes and seconds. You could use the following:

    ds, err := durationslice.Process(time.Duration((48*time.Hour)+(10*time.Second)), "d min s")
	if err != nil {
	    print(err.Error())
		return
	}
	fmt.Printf("%@", ds)
`output: [2 0 10]`

The resulting slice would give the values corresponding to the units
specified in the unit string.