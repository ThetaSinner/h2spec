package client

import "thetasinner/h2spec/spec"

func HTTPFrames() *spec.ClientTestGroup {
	tg := NewTestGroup("4", "HTTP Frames")

	tg.AddTestGroup(FrameFormat())
	tg.AddTestGroup(FrameSize())
	tg.AddTestGroup(HeaderCompressionAndDecompression())

	return tg
}
