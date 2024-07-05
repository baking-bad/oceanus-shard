package component

type EffectType string

const (
	Raft EffectType = "Raft"
)

func GetAllEffectTypes() []EffectType {
	return []EffectType{
		Raft,
	}
}

type Effect struct {
	Type                  EffectType `json:"type"`
	Amount                int        `json:"amount"`
	Capacity              int        `json:"capacity"`
	Resources             []Resource `json:"resources,omitempty"`
	BuildingTimeSeconds   int        `json:"buildingTimeSeconds,omitempty"`
	BuildingTimeStartedAt uint64     `json:"buildingTimeStartedAt,omitempty"`
}

func (Effect) Name() string {
	return "Effect"
}
