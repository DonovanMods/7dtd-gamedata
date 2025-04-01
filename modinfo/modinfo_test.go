package modinfo_test

import (
	XML "encoding/xml"
	"testing"

	"github.com/donovanmods/7dtd-gamedata/modinfo"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	modletName := "TestModlet"
	modInfo := modinfo.NewModInfo(modletName)

	assert.Equal(t, modletName, modInfo.Name.Value, "Expected Name.Value to be %s", modletName)
	assert.Equal(t, "My New Modlet", modInfo.DisplayName.Value, "Expected DisplayName.Value to be 'My New Modlet'")
	expectedDescription := "This is the description for TestModlet -- please change it"
	assert.Equal(t, expectedDescription, modInfo.Description.Value, "Expected Description.Value to be '%s'", expectedDescription)
	assert.Equal(t, "Unknown", modInfo.Author.Value, "Expected Author.Value to be 'Unknown'")
	assert.Equal(t, "0.1.0", modInfo.Version.Value, "Expected Version.Value to be '0.1.0'")
	assert.Equal(t, "V1", modInfo.Version.Compat, "Expected Version.Compat to be 'V1'")
	assert.Equal(t, "https://example.org", modInfo.Website.Value, "Expected Website.Value to be 'https://example.org'")
}

func TestXML(t *testing.T) {
	modletName := "TestModlet"
	modInfo := modinfo.NewModInfo(modletName)

	xmlOutput, err := modInfo.XML()
	assert.NoError(t, err, "Expected no error while generating XML")

	expectedXML := XML.Header + `<xml>
  <Name value="TestModlet" />
  <DisplayName value="My New Modlet" />
  <Description value="This is the description for TestModlet -- please change it" />
  <Author value="Unknown" />
  <Version value="0.1.0" compat="V1" />
  <Website value="https://example.org" />
</xml>`

	assert.Equal(t, expectedXML, xmlOutput, "Expected XML output to match")
}

func TestSetters(t *testing.T) {
	modletName := "TestModlet"
	modInfo := modinfo.NewModInfo(modletName)

	// Test SetDescription
	description := "This is a test description"
	modInfo.SetDescription(description)
	assert.Equal(t, description, modInfo.Description.Value, "Expected Description.Value to be '%s'", description)

	// Test SetAuthor
	author := "Test Author"
	modInfo.SetAuthor(author)
	assert.Equal(t, author, modInfo.Author.Value, "Expected Author.Value to be '%s'", author)

	// Test SetVersion
	version := "1.0.0"
	compat := "V2"
	modInfo.SetVersion(version, compat)
	assert.Equal(t, version, modInfo.Version.Value, "Expected Version.Value to be '%s'", version)
	assert.Equal(t, compat, modInfo.Version.Compat, "Expected Version.Compat to be '%s'", compat)

	// Test SetVersion without Compat
	modInfo.SetVersion(version, "")
	assert.Equal(t, version, modInfo.Version.Value, "Expected Version.Value to be '%s'", version)
	assert.Empty(t, modInfo.Version.Compat, "Expected Version.Compat to be empty")

	// Test SetWebsite
	website := "https://testwebsite.com"
	modInfo.SetWebsite(website)
	assert.Equal(t, website, modInfo.Website.Value, "Expected Website.Value to be '%s'", website)
}

func TestPath(t *testing.T) {
	modletName := "TestModlet"
	modInfo := modinfo.NewModInfo(modletName)

	path := "/path/to/modlet.xml"
	modInfo.SetPath(path)
	assert.Equal(t, path, modInfo.Path(), "Expected Path to be '%s'", path)
	assert.Equal(t, "modlet.xml", modInfo.Filename(), "Expected Filename to be 'modlet.xml'")
	assert.Equal(t, "/path/to", modInfo.Dir(), "Expected Dir to be '/path/to'")
}
