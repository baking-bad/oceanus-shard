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
	Type   EffectType `json:"type"`
	Amount int        `json:"amount"`
}

func (Effect) Name() string {
	return "Effect"
}
