package http2

import (
	"golang.org/x/net/http2"

	"thetasinner/h2spec/config"
	"thetasinner/h2spec/spec"
)

func ErrorCodes() *spec.TestGroup {
	tg := NewTestGroup("7", "Error Codes")

	// Unknown or unsupported error codes MUST NOT trigger any special
	// behavior. These MAY be treated by an implementation as being
	// equivalent to INTERNAL_ERROR.
	tg.AddTestCase(&spec.TestCase{
		Desc:        "Sends a GOAWAY frame with unknown error code",
		Requirement: "The endpoint MUST NOT trigger any special behavior.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			err := conn.Handshake()
			if err != nil {
				return err
			}

			conn.WriteGoAway(0, 0xff, []byte{})

			return spec.VerifyConnectionClose(conn)
		},
	})

	// Unknown or unsupported error codes MUST NOT trigger any special
	// behavior. These MAY be treated by an implementation as being
	// equivalent to INTERNAL_ERROR.
	tg.AddTestCase(&spec.TestCase{
		Desc:        "Sends a RST_STREAM frame with unknown error code",
		Requirement: "The endpoint MUST NOT trigger any special behavior.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			var streamID uint32 = 1

			err := conn.Handshake()
			if err != nil {
				return err
			}

			headers := spec.CommonHeaders(c)
			headers[0].Value = "POST"

			hp := http2.HeadersFrameParam{
				StreamID:      streamID,
				EndStream:     false,
				EndHeaders:    true,
				BlockFragment: conn.EncodeHeaders(headers),
			}

			conn.WriteHeaders(hp)
			conn.WriteRSTStream(streamID, 0xff)

			data := [8]byte{}
			conn.WritePing(false, data)

			return spec.VerifyPingFrameWithAck(conn, data)
		},
	})

	return tg
}
