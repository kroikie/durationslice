你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Duration Slice ![Build Status](https://travis-ci.org/kroikie/durationslice.svg?branch=master)
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
If you had a duration that was 75 minutes and 65 seconds long and you wanted
to represent that as hours, minutes and seconds. You could use the following:

    ds, err := durationslice.Process(time.Duration((75*time.Minute)+(65*time.Second)), "h m s")
	if err != nil {
	    print(err.Error())
		return
	}
	fmt.Printf("%@", ds)
`output: [1 16 5]`

The resulting slice would give the values corresponding to the units
specified in the unit string.
