package component

type ResourceType string

const (
	Wood   ResourceType = "Wood"
	Stone  ResourceType = "Stone"
	Fish   ResourceType = "Fish"
	Iron   ResourceType = "Iron"
	Cotton ResourceType = "Cotton"
)

type Resource struct {
	Type   ResourceType `json:"type"`
	Amount int          `json:"amount"`
}

func (Resource) Name() string {
	return "Resource"
}
