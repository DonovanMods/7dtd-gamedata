package gamexml

type Block struct {
	XMLName  string     `xml:"block"`
	Name     string     `xml:"name,attr"`
	Shapes   string     `xml:"shapes,attr,omitempty"`
	Property []Property `xml:"property"`
	Events   []Event    `xml:"drop"`
	DEO      string     `xml:"dropextendsoff,omitempty"`
	Comments []string   `xml:"-"`
}

var Blocks = []Block{}
