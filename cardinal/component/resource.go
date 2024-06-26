package component

type ResourceType string

const (
	Wood   ResourceType = "Wood"
	Stone  ResourceType = "Stone"
	Fish   ResourceType = "Fish"
	Iron   ResourceType = "Iron"
	Cotton ResourceType = "Cotton"
)

func GetAllResourceTypes() []ResourceType {
	return []ResourceType{
		Wood,
		Stone,
		Fish,
		Iron,
		Cotton,
	}
}

type Resource struct {
	Type   ResourceType `json:"type"`
	Amount float32      `json:"amount"`
}

func (Resource) Name() string {
	return "Resource"
}
