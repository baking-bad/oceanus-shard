package msg

type RelocateBuildingMsg struct {
	TileIndexFrom int `json:"tileIndexFrom"`
	TileIndexTo   int `json:"tileIndexTo"`
}

type RelocateBuildingResult struct {
	Success bool `json:"success"`
}
