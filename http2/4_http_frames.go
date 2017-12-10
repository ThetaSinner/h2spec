package http2

import "thetasinner/h2spec/spec"

func HTTPFrames() *spec.TestGroup {
	tg := NewTestGroup("4", "HTTP Frames")

	tg.AddTestGroup(FrameFormat())
	tg.AddTestGroup(FrameSize())
	tg.AddTestGroup(HeaderCompressionAndDecompression())

	return tg
}
