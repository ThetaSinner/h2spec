package http2

import "thetasinner/h2spec/spec"

func StreamPriority() *spec.TestGroup {
	tg := NewTestGroup("5.3", "Stream Priority")

	tg.AddTestGroup(StreamDependencies())

	return tg
}
