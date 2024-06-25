package component

type Farming struct {
	Type  ResourceType `json:"type"`
	Speed float32      `json:"speed"`
}

func (Farming) Name() string {
	return "Farming"
}
