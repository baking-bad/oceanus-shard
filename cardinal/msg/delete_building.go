package msg

type DeleteBuildingMsg struct {
	TileIndex int `json:"tileID"`
}

type DeleteBuildingResult struct {
	Success bool `json:"success"`
}
