package hpack

import (
	"thetasinner/h2spec/config"
	"thetasinner/h2spec/spec"
	"golang.org/x/net/http2"
)
 
func MaximumTableSize() *spec.TestGroup {
	tg := NewTestGroup("4.2", "Maximum Table Size")

	// A change in the maximum size of the dynamic table is signaled
	// via a dynamic table size update (see Section 6.3). This dynamic
	// table size update MUST occur at the beginning of the first
	// header block following the change to the dynamic table size.
	// In HTTP/2, this follows a settings acknowledgment (see Section
	// 6.5.3 of [HTTP2]).
	tg.AddTestCase(&spec.TestCase{
		Desc:        "Sends reduced MAXIMUM_TABLE_SIZE setting",
		Requirement: "Server must acknowledge then send dynamic table size update as the first instruction in the next header block it sends",
		Run: func(c *config.Config, conn *spec.Conn) error {

			err := conn.Handshake()
			if err != nil {
				return err
			}

			setting := http2.Setting{
				ID:  http2.SettingHeaderTableSize,
				Val: conn.MaxFrameSize() - 1,
			}

			conn.WriteSettings(setting)
			return spec.VerifySettingsFrameWithAck(conn)
		},
	})

	return tg
}
