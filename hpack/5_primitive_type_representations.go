package hpack

import "thetasinner/h2spec/spec"

func PrimitiveTypeRepresentations() *spec.TestGroup {
	tg := NewTestGroup("5", "Primitive Type Representations")

	tg.AddTestGroup(StringLiteralRepresentation())

	return tg
}
