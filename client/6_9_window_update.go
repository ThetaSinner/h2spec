package client

import (
	"golang.org/x/net/http2"

	"thetasinner/h2spec/config"
	"thetasinner/h2spec/spec"
)

func WindowUpdate() *spec.ClientTestGroup {
	tg := NewTestGroup("6.9", "WINDOW_UPDATE")

	// A receiver MUST treat the receipt of a WINDOW_UPDATE frame with
	// an flow-control window increment of 0 as a stream error
	// (Section 5.4.2) of type PROTOCOL_ERROR; errors on the connection
	// flow-control window MUST be treated as a connection error
	// (Section 5.4.1).
	tg.AddTestCase(&spec.ClientTestCase{
		Desc:        "Sends a WINDOW_UPDATE frame with a flow control window increment of 0",
		Requirement: "The endpoint MUST treat this as a connection error of type PROTOCOL_ERROR.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			err := conn.Handshake()
			if err != nil {
				return err
			}

			conn.WriteWindowUpdate(0, 0)

			return spec.VerifyConnectionError(conn, http2.ErrCodeProtocol)
		},
	})

	// A receiver MUST treat the receipt of a WINDOW_UPDATE frame with
	// an flow-control window increment of 0 as a stream error
	// (Section 5.4.2) of type PROTOCOL_ERROR; errors on the connection
	// flow-control window MUST be treated as a connection error
	// (Section 5.4.1).
	tg.AddTestCase(&spec.ClientTestCase{
		Desc:        "Sends a WINDOW_UPDATE frame with a flow control window increment of 0 on a stream",
		Requirement: "The endpoint MUST treat this as a connection error of type PROTOCOL_ERROR.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			err := conn.Handshake()
			if err != nil {
				return err
			}

			req, err := conn.ReadRequest()
			if err != nil {
				return err
			}

			headers := spec.CommonRespHeaders(c)
			hp := http2.HeadersFrameParam{
				StreamID:      req.StreamID,
				EndStream:     false,
				EndHeaders:    true,
				BlockFragment: conn.EncodeHeaders(headers),
			}
			conn.WriteHeaders(hp)

			conn.WriteWindowUpdate(req.StreamID, 0)

			return spec.VerifyStreamError(conn, http2.ErrCodeProtocol)
		},
	})

	// A WINDOW_UPDATE frame with a length other than 4 octets MUST
	// be treated as a connection error (Section 5.4.1) of type
	// FRAME_SIZE_ERROR.
	tg.AddTestCase(&spec.ClientTestCase{
		Desc:        "Sends a WINDOW_UPDATE frame with a length other than 4 octets",
		Requirement: "The endpoint MUST treat this as a connection error of type FRAME_SIZE_ERROR.",
		Run: func(c *config.Config, conn *spec.Conn) error {
			err := conn.Handshake()
			if err != nil {
				return err
			}

			// WINDOW_UPDATE frame:
			// length: 3, flags: 0x0, stream_id: 0
			conn.Send([]byte("\x00\x00\x03\x08\x00\x00\x00\x00\x00"))
			conn.Send([]byte("\x00\x00\x01"))

			return spec.VerifyConnectionError(conn, http2.ErrCodeFrameSize)
		},
	})

	tg.AddTestGroup(TheFlowControlWindow())

	return tg
}
