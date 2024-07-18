package component

type PlayerResources struct {
	Resources []Resource `json:"resources"`
	Ships     []Ship     `json:"ships"`
}

func (PlayerResources) Name() string {
	return "PlayerResources"
}
