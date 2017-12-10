package client

import (
	"thetasinner/h2spec/config"
	"thetasinner/h2spec/spec"
)

func StartingHTTP2() *spec.ClientTestGroup {
	tg := NewTestGroup("1", "Starting HTTP/2")

	tg.AddTestCase(&spec.ClientTestCase{
		Desc:        "Sends a server connection preface",
		Requirement: "The endpoint MUST accept server connection preface.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			err := conn.Handshake()
			return err
		},
	})

	return tg
}
