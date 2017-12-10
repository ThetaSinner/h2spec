package client

import "thetasinner/h2spec/spec"

func ErrorHandling() *spec.ClientTestGroup {
	tg := NewTestGroup("5.4", "Error Handling")

	tg.AddTestGroup(ConnectionErrorHandling())

	return tg
}
