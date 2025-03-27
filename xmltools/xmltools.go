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
package xmltools

import "regexp"

func RemoveClosingXMLTags(xml string) string {
	return removeClosingXMLTags(xml, false)
}

func RemoveClosingXMLTagsCompact(xml string) string {
	return removeClosingXMLTags(xml, true)
}

func removeClosingXMLTags(xml string, compact bool) string {
	closingTag := "/>"

	if !compact {
		closingTag = " />"
	}

	re := regexp.MustCompile(`></\b\w+\b>`)

	return re.ReplaceAllString(string(xml), closingTag)
}
