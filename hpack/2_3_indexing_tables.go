package hpack

import "thetasinner/h2spec/spec"

func IndexingTables() *spec.TestGroup {
	tg := NewTestGroup("2.3", "Indexing Tables")

	tg.AddTestGroup(IndexAddressSpace())

	return tg
}
