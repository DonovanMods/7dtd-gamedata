package gamexml_test

import (
	_ "embed"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/donovanmods/7dtd-gamedata/gamexml"
	"github.com/donovanmods/7dtd-gamedata/xmltools"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/entityclasses_in.xml
var inputEntityClassesXML string

//go:embed testdata/entityclasses_out.xml
var outputEntityClassesXML string

// TestBlockXML tests the XML encoding and decoding of a Block.
func TestEntityClassesXML(t *testing.T) {
	// Decode the XML data into a Block value.
	var entityclasses gamexml.EntityClasses

	if err := xml.NewDecoder(strings.NewReader(inputEntityClassesXML)).Decode(&entityclasses); err != nil {
		t.Fatal(err)
	}

	// Encode the Block value into XML data.
	var buf strings.Builder
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := enc.Encode(entityclasses); err != nil {
		t.Fatal(err)
	}

	expected := outputEntityClassesXML
	actual := xml.Header + xmltools.RemoveClosingXMLTags(buf.String()) + "\n"

	assert.Equal(t, expected, actual, "Expected XML output to match")
}
