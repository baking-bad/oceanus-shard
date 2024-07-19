package component

import "oceanus-shard/constants"

type EffectType string

const (
	Raft EffectType = "Raft"
)

func GetAllEffectTypes() []EffectType {
	return []EffectType{
		Raft,
	}
}

func GetCapacityByEffectType(effectType EffectType) int {
	switch effectType {
	case Raft:
		return constants.ShipyardEffectRaftCapacity
	default:
		return 0
	}
}

type Effect struct {
	ID                    string      `json:"id"`
	Player                string      `json:"player"`
	Type                  EffectType  `json:"type"`
	Amount                int         `json:"amount"`
	Capacity              int         `json:"capacity"`
	Resources             []Resource  `json:"resources,omitempty"`
	Position              [2]float64  `json:"position,omitempty"`
	TargetPosition        *[2]float64 `json:"targetPosition,omitempty"`
	BuildingTimeSeconds   int         `json:"buildingTimeSeconds,omitempty"`
	BuildingTimeStartedAt uint64      `json:"buildingTimeStartedAt,omitempty"`
}

func (Effect) Name() string {
	return "Effect"
}
