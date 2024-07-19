package msg

type SailShipwreckMsg struct {
	Player string `json:"player"`
}

type SailShipWreckResult struct {
	Success bool `json:"success"`
}
