package component

type PlayerResources struct {
	Resources []Resource `json:"resources"`
	Effects   []Effect   `json:"effects"`
	// Farming   []Farming  `json:"farming"`
}

func (PlayerResources) Name() string {
	return "PlayerResources"
}
