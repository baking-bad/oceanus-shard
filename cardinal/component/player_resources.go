package component

type PlayerResources struct {
	Resources []Resource `json:"resources"`
	Effects   []Effect   `json:"effects"`
}

func (PlayerResources) Name() string {
	return "PlayerResources"
}
