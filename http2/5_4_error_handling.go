package http2

import "thetasinner/h2spec/spec"

func ErrorHandling() *spec.TestGroup {
	tg := NewTestGroup("5.4", "Error Handling")

	tg.AddTestGroup(ConnectionErrorHandling())

	return tg
}
