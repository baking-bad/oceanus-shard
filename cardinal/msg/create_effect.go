package msg

type CreateEffectMsg struct {
	EffectType string `json:"effectType"`
	TileIndex  int    `json:"tileIndex"`
}

type CreateEffectResult struct {
	Success bool `json:"success"`
}
