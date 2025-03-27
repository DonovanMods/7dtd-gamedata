package gamexml

type Requirement struct {
	XMLName         string `xml:"requirement"`
	Name            string `xml:"name,attr"`
	Cvar            string `xml:"cvar,attr,omitempty"`
	Buff            string `xml:"buff,attr,omitempty"`
	ProgressionName string `xml:"progression_name,attr,omitempty"`
	Operation       string `xml:"operation,attr,omitempty"`
	Value           string `xml:"value,attr,omitempty"`
	Tags            Tags   `xml:"tags,attr,omitempty"`
}

type TriggeredEffect struct {
	XMLName     string        `xml:"triggered_effect"`
	Trigger     string        `xml:"trigger,attr,omitempty"`
	Action      string        `xml:"action,attr,omitempty"`
	Target      string        `xml:"target,attr,omitempty"`
	Buff        string        `xml:"buff,attr,omitempty"`
	Requirement []Requirement `xml:"requirement,omitempty"`
}

type PassiveEffect struct {
	XMLName     string        `xml:"passive_effect"`
	Name        string        `xml:"name,attr"`
	Operation   string        `xml:"operation,attr,omitempty"`
	Value       string        `xml:"value,attr,omitempty"`
	MAT         bool          `xml:"match_all_tags,attr,omitempty"`
	Tags        Tags          `xml:"tags,attr,omitempty"`
	Requirement []Requirement `xml:"requirement,omitempty"`
}

type EffectGroup struct {
	Name            string            `xml:"name,attr,omitempty"`
	Requirement     []Requirement     `xml:"requirement,omitempty"`
	PassiveEffect   []PassiveEffect   `xml:"passive_effect,omitempty"`
	TriggeredEffect []TriggeredEffect `xml:"triggered_effect,omitempty"`
}

type EntityClass struct {
	XMLName     string        `xml:"entity_class"`
	Name        string        `xml:"name,attr"`
	Property    []Property    `xml:"property,omitempty"`
	EffectGroup []EffectGroup `xml:"effect_group,omitempty"`
}

type EntityClasses struct {
	XMLName     string        `xml:"entity_classes"`
	EntityClass []EntityClass `xml:"entity_class"`
}
