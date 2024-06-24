package component

import "fmt"

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

const (
	InitialBuildingLevel        int     = 1
	MainBuildingUnitLimit       int     = 5
	MainBuildingStorageCapacity int     = 200
	QuarryFarmingSpeed          float32 = 10.0
	QuarryUnitLimit             int     = 0
	QuarryStorageCapacity       int     = 200
	FishermanHutFarmingSpeed    float32 = 8.0
	FishermanHutUnitLimit       int     = 0
	FishermanHutStorageCapacity int     = 150
	ShipyardUnitLimit           int     = 100
	ShipyardStorageCapacity     int     = 300
)

func GetBuilding(buildingType BuildingType) (Building, error) {
	switch buildingType {
	case Main:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Main,
			UnitLimit:       MainBuildingUnitLimit,
			StorageCapacity: MainBuildingStorageCapacity,
		}, nil
	case Quarry:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Quarry,
			FarmingResource: Stone,
			FarmingSpeed:    QuarryFarmingSpeed,
			UnitLimit:       QuarryUnitLimit,
			StorageCapacity: QuarryStorageCapacity,
		}, nil
	case FishermanHut:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            FishermanHut,
			FarmingResource: Fish,
			FarmingSpeed:    FishermanHutFarmingSpeed,
			UnitLimit:       FishermanHutUnitLimit,
			StorageCapacity: FishermanHutStorageCapacity,
		}, nil
	case Shipyard:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Shipyard,
			UnitLimit:       ShipyardUnitLimit,
			StorageCapacity: ShipyardStorageCapacity,
		}, nil

	// todo: refactor
	case Woodcutter:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Shipyard,
			UnitLimit:       ShipyardUnitLimit,
			StorageCapacity: ShipyardStorageCapacity,
		}, nil
		// todo: refactor
	case UnitLimitHouse:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Shipyard,
			UnitLimit:       ShipyardUnitLimit,
			StorageCapacity: ShipyardStorageCapacity,
		}, nil
		// todo: refactor
	case Warehouse:
		return Building{
			Level:           InitialBuildingLevel,
			Type:            Shipyard,
			UnitLimit:       ShipyardUnitLimit,
			StorageCapacity: ShipyardStorageCapacity,
		}, nil
	default:
		return Building{
			Level:           0,
			Type:            buildingType,
			UnitLimit:       0,
			StorageCapacity: 0,
		}, fmt.Errorf("%s is invalid building type", buildingType)
	}
}
