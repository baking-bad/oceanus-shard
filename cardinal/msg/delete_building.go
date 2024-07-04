package msg

type DeleteBuildingMsg struct {
	TileIndex int `json:"tileIndex"`
}

type DeleteBuildingResult struct {
	Success bool `json:"success"`
}
