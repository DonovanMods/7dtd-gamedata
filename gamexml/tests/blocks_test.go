package gamexml_blocks_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/donovanmods/7dtd-gamedata/gamexml"
	"github.com/donovanmods/7dtd-gamedata/xmltools"
	"github.com/stretchr/testify/assert"
)

// TestBlockXML tests the XML encoding and decoding of a Block.
func TestBlockXML(t *testing.T) {
	// XML data for a Block.
	const blockXML = `<block name="corrugatedMetalShapes" shapes="All">
  <property name="DescriptionKey" value="scrapIronGroupDesc"/>
  <property name="Material" value="Mmetal_shapes"/>
  <property name="Shape" value="New"/>
  <property name="Texture" value="194"/>
  <property name="UiBackgroundTexture" value="194"/>
  <property name="EconomicValue" value="15"/>
  <property name="EconomicBundleSize" value="20"/>
  <property class="RepairItems">
    <property name="resourceScrapIron" value="12"/>
  </property>
  <property class="UpgradeBlock">
    <property name="ToBlock" value="cobblestoneShapes"/>
    <property name="Item" value="resourceCobblestones"/>
    <property name="ItemCount" value="10"/>
    <property name="UpgradeHitCount" value="4"/>
  </property>
  <property name="UpgradeSound" value="place_block_metal"/>
  <property name="SortOrder1" value="S060"/>
  <property name="Group" value="Building,advBuilding"/>
  <property name="FilterTags" value="MC_Shapes"/>
  <drop event="Harvest" name="resourceScrapIron" count="10" tag="allHarvest"/>
  <drop event="Destroy" count="0"/>
  <drop event="Fall" name="scrapMetalPile" count="1" prob="0.6" stick_chance="1"/>
</block>`

	// Decode the XML data into a Block value.
	var block gamexml.Block

	if err := xml.NewDecoder(strings.NewReader(blockXML)).Decode(&block); err != nil {
		t.Fatal(err)
	}

	// Encode the Block value into XML data.
	var buf strings.Builder
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")
	if err := enc.Encode(block); err != nil {
		t.Fatal(err)
	}

	received := xmltools.RemoveClosingXMLTagsCompact(buf.String())

	assert.Equal(t, blockXML, received, "Expected XML output to match")
}
