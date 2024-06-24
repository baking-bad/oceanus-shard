package msg

type CreateBuildingMsg struct {
	BuildingType string `json:"buildingType"`
	TileIndex    int    `json:"tileIndex"`
}

type CreateBuildingResult struct {
	Success bool `json:"success"`
}
