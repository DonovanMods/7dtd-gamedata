package gamexml

type Tags string // a list of CSV tags

type Property struct {
	XMLName  string     `xml:"property"`
	Name     string     `xml:"name,attr,omitempty"`
	Value    string     `xml:"value,attr,omitempty"`
	Class    string     `xml:"class,attr,omitempty"`
	Data     string     `xml:"data,attr,omitempty"`
	Property []Property `xml:"property,omitempty"`
}
