package component

type Farming struct {
	Type     ResourceType `json:"type"`
	Speed    float32      `json:"speed"`
	Capacity int          `json:"capacity,omitempty"`
}

func (Farming) Name() string {
	return "Farming"
}
