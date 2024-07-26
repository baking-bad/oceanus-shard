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
	LootResources         *[]Resource `json:"lootResources,omitempty"`
	Speed                 float64     `json:"speed"`
	Position              [2]float64  `json:"position,omitempty"`
	TargetPosition        *[2]float64 `json:"targetPosition,omitempty"`
	SendingPosition       *[2]float64 `json:"sendingPosition,omitempty"`
	BuildingTimeSeconds   int         `json:"buildingTimeSeconds,omitempty"`
	BuildingTimeStartedAt uint64      `json:"buildingTimeStartedAt,omitempty"`
}

func (Effect) Name() string {
	return "Effect"
}
