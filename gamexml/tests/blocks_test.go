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

//go:embed testdata/blocks_in.xml
var inputBlocksXML string

//go:embed testdata/blocks_out.xml
var outputBlocksXML string

// TestBlockXML tests the XML encoding and decoding of a Block.
func TestBlocksXML(t *testing.T) {
	// Decode the XML data into a Block value.
	var blocks gamexml.Blocks

	if err := xml.NewDecoder(strings.NewReader(inputBlocksXML)).Decode(&blocks); err != nil {
		t.Fatal(err)
	}

	// Encode the Block value into XML data.
	var buf strings.Builder
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := enc.Encode(blocks); err != nil {
		t.Fatal(err)
	}

	expected := outputBlocksXML
	actual := xml.Header + xmltools.RemoveClosingXMLTags(buf.String()) + "\n"

	assert.Equal(t, expected, actual, "Expected XML output to match")
}
