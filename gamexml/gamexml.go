package gamexml

const XMLHeader = `<?xml version="1.0" encoding="UTF-8"?>`

type Property struct {
	XMLName  string     `xml:"property"`
	Name     string     `xml:"name,attr,omitempty"`
	Value    string     `xml:"value,attr,omitempty"`
	Class    string     `xml:"class,attr,omitempty"`
	Property []Property `xml:"property,omitempty"`
}

type Event struct {
	XMLName     string // the event type (e.g. "drop")
	Event       string `xml:"event,attr"` // the event type (e.g. "Harvest")
	Name        string `xml:"name,attr,omitempty"`
	Count       string `xml:"count,attr,omitempty"`
	Probability string `xml:"prob,attr,omitempty"`
	Tag         string `xml:"tag,attr,omitempty"`
	StickChance string `xml:"stick_chance,attr,omitempty"`
}
