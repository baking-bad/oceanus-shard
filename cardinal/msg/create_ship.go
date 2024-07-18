package msg

type CreateShipMsg struct {
	ShipType  string `json:"shipType"`
	TileIndex int    `json:"tileIndex"`
}

type CreateShipResult struct {
	Success bool `json:"success"`
}
