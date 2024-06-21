package component

type BuildingType string
type EffectsType string

const (
	Main           BuildingType = "Main"
	Woodcutter     BuildingType = "Woodcutter"
	Quarry         BuildingType = "Quarry"
	FishermanHut   BuildingType = "FishermanHut"
	Shipyard       BuildingType = "Shipyard"
	Warehouse      BuildingType = "Warehouse"
	UnitLimitHouse BuildingType = "UnitLimitHouse"
)

const (
	Ship EffectsType = "Ship"
)

type Building struct {
	Level           int          `json:"level"`
	Type            BuildingType `json:"type"`
	FarmingResource ResourceType `json:"farmingResource,omitempty"`
	FarmingSpeed    float32      `json:"farmingSpeed,omitempty"`
	Effect          EffectsType  `json:"effect,omitempty"`
	EffectAmount    int          `json:"effectAmount,omitempty"`
	UnitLimit       int          `json:"unitLimit"`
	StorageCapacity int          `json:"storageCapacity"`
}

func (Building) Name() string {
	return "Building"
}

func GetBuilding(buildingType BuildingType) Building {
	switch buildingType {
	case Main:
		return Building{
			Level:           1,
			Type:            Main,
			UnitLimit:       5,
			StorageCapacity: 200,
		}
	case Quarry:
		return Building{
			Level:           1,
			Type:            Quarry,
			FarmingResource: Stone,
			FarmingSpeed:    10.0,
			UnitLimit:       0,
			StorageCapacity: 200,
		}
	case FishermanHut:
		return Building{
			Level:           1,
			Type:            FishermanHut,
			FarmingResource: Fish,
			FarmingSpeed:    8.0,
			UnitLimit:       0,
			StorageCapacity: 150,
		}
	case Shipyard:
		return Building{
			Level:           1,
			Type:            Shipyard,
			UnitLimit:       100,
			StorageCapacity: 300,
		}
	default:
		return Building{
			Level:           0,
			Type:            buildingType,
			UnitLimit:       0,
			StorageCapacity: 0,
		}
	}
}
