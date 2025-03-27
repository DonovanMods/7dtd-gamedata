package gamexml

type DropEvent struct {
	XMLName      string `xml:"drop"`
	Event        string `xml:"event,attr"` // the event type (e.g. "Harvest")
	Name         string `xml:"name,attr,omitempty"`
	Count        string `xml:"count,attr,omitempty"`
	Probability  string `xml:"prob,attr,omitempty"`
	Tag          Tags   `xml:"tag,attr,omitempty"`
	StickChance  string `xml:"stick_chance,attr,omitempty"`
	ToolCategory string `xml:"tool_category,attr,omitempty"`
}

type Block struct {
	XMLName  string      `xml:"block"`
	Name     string      `xml:"name,attr"`
	Shapes   string      `xml:"shapes,attr,omitempty"`
	Property []Property  `xml:"property,omitempty"`
	Events   []DropEvent `xml:"drop,omitempty"`
	DEO      string      `xml:"dropextendsoff,omitempty"`
}

type Blocks struct {
	XMLName string  `xml:"blocks"`
	Block   []Block `xml:"block"`
	DDKey   string  `xml:"defaultDescriptionKey,attr,omitempty"`
}
