/*
Copyright © 2024 Donovan C. Young <dyoung522@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package modinfo

import (
	XML "encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/donovanmods/7dtd-gamedata/xmltools"
)

type xmlValue struct {
	Value  string `xml:"value,attr"`
	Compat string `xml:"compat,attr,omitempty"`
}

type modInfoMeta struct {
	path string
}

type ModInfo struct {
	XMLName     XML.Name `xml:"xml"`
	Name        xmlValue `xml:"Name"`
	DisplayName xmlValue `xml:"DisplayName"`
	Description xmlValue `xml:"Description"`
	Author      xmlValue `xml:"Author"`
	Version     xmlValue `xml:"Version"`
	Website     xmlValue `xml:"Website,omitempty"`
	meta        modInfoMeta
}

func (M *ModInfo) GetValue(item string) string {
	switch {
	case strings.EqualFold(item, "name"):
		return M.Name.Value
	case strings.EqualFold(item, "displayname"):
		return M.DisplayName.Value
	case strings.EqualFold(item, "description"):
		return M.Description.Value
	case strings.EqualFold(item, "author"):
		return M.Author.Value
	case strings.EqualFold(item, "version"):
		return M.Version.Value
	case strings.EqualFold(item, "compat"):
		return M.Version.Compat
	case strings.EqualFold(item, "website"):
		return M.Website.Value
	default:
		return ""
	}
}

func (M *ModInfo) SetValue(item string, value string) error {
	switch {
	case strings.EqualFold(item, "name"):
		M.Name.Value = value
	case strings.EqualFold(item, "displayname"):
		M.DisplayName.Value = value
	case strings.EqualFold(item, "description"):
		M.DisplayName.Value = value
	case strings.EqualFold(item, "author"):
		M.Author.Value = value
	case strings.EqualFold(item, "version"):
		M.Version.Value = value
	case strings.EqualFold(item, "compat"):
		M.Version.Compat = value
	case strings.EqualFold(item, "website"):
		M.Website.Value = value
	default:
		return fmt.Errorf("unknown item: %s", item)
	}
	return nil
}

func (M *ModInfo) SetName(value string) {
	_ = M.SetValue("name", value)
}

func (M *ModInfo) SetDisplayName(value string) {
	_ = M.SetValue("displayname", value)
}

func (M *ModInfo) SetDescription(value string) {
	_ = M.SetValue("description", value)
}

func (M *ModInfo) SetAuthor(value string) {
	_ = M.SetValue("author", value)
}

func (M *ModInfo) SetVersion(version string, compatibility string) {
	_ = M.SetValue("version", version)
	_ = M.SetValue("compat", compatibility)
}

func (M *ModInfo) SetWebsite(value string) {
	_ = M.SetValue("website", value)
}

func (M *ModInfo) SetPath(path string) {
	M.meta.path = path
}

func (M *ModInfo) Path() string {
	return M.meta.path
}

func (M *ModInfo) Dir() string {
	return filepath.Dir(M.meta.path)
}

func (M *ModInfo) Filename() string {
	return filepath.Base(M.meta.path)
}

func (M *ModInfo) XML() (string, error) {
	xml, err := XML.MarshalIndent(M, "", "  ")
	if err != nil {
		return "", err
	}

	xml = []byte(xmltools.RemoveClosingXMLTags(string(xml)))

	return XML.Header + string(xml), nil
}

func NewModInfo(modletName string) *ModInfo {
	return &ModInfo{
		Name:        xmlValue{Value: modletName},
		DisplayName: xmlValue{Value: "My New Modlet"},
		Description: xmlValue{Value: fmt.Sprintf("This is the description for %s -- please change it", modletName)},
		Author:      xmlValue{Value: "Unknown"},
		Version:     xmlValue{Value: "0.1.0", Compat: "V1"},
		Website:     xmlValue{Value: "https://example.org"},
	}
}
