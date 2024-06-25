package component

import "fmt"

type BuildingType string

const (
	Main           BuildingType = "Main"
	Woodcutter     BuildingType = "Woodcutter"
	Quarry         BuildingType = "Quarry"
	FishermanHut   BuildingType = "FishermanHut"
	Shipyard       BuildingType = "Shipyard"
	Warehouse      BuildingType = "Warehouse"
	UnitLimitHouse BuildingType = "UnitLimitHouse"
)

func GetAllBuildingTypes() []BuildingType {
	return []BuildingType{
		Main,
		Woodcutter,
		Quarry,
		FishermanHut,
		Shipyard,
		Warehouse,
		UnitLimitHouse,
	}
}

type Building struct {
	Level           int          `json:"level"`
	Type            BuildingType `json:"type"`
	Farming         *Farming     `json:"farming,omitempty"`
	Effect          *Effect      `json:"effect,omitempty"`
	UnitLimit       int          `json:"unitLimit"`
	StorageCapacity int          `json:"storageCapacity"`
}

func (Building) Name() string {
	return "Building"
}

type BuildingConstants struct {
	UnitLimit       int
	StorageCapacity int
	Resources       []Resource
	Farming         *Farming
	Effect          *Effect
	TileType        TileType
}

var BuildingConfigs = map[BuildingType]BuildingConstants{
	Main: {
		Resources:       []Resource{},
		UnitLimit:       5,
		StorageCapacity: 200,
		TileType:        GenericTile,
	},
	Woodcutter: {
		Resources: []Resource{},
		Farming: &Farming{
			Type:  Wood,
			Speed: 5,
		},
		TileType: GenericTile,
	},
	Quarry: {
		Resources: []Resource{},
		Farming: &Farming{
			Type:  Stone,
			Speed: 10.0,
		},
		TileType: GenericTile,
	},
	FishermanHut: {
		Resources: []Resource{
			{Type: Wood, Amount: 50},
			{Type: Stone, Amount: 50},
		},
		Farming: &Farming{
			Type:  Fish,
			Speed: 5,
		},
		TileType: GenericTile,
	},
	Shipyard: {
		Resources: []Resource{
			{Type: Wood, Amount: 100},
			{Type: Stone, Amount: 100},
		},
		Effect: &Effect{
			Type:   Raft,
			Amount: 2,
		},
		TileType: GenericTile,
	},
}

func GetBuilding(buildingType BuildingType) (Building, error) {
	config := BuildingConfigs[buildingType]

	switch buildingType {
	case Main:
		return Building{
			Level:           1,
			Type:            buildingType,
			UnitLimit:       config.UnitLimit,
			StorageCapacity: config.StorageCapacity,
		}, nil
	case Quarry:
		return Building{
			Level:           1,
			Type:            buildingType,
			Farming:         config.Farming,
			UnitLimit:       config.UnitLimit,
			StorageCapacity: config.StorageCapacity,
		}, nil
	case FishermanHut:
		return Building{
			Level:   1,
			Type:    buildingType,
			Farming: config.Farming,
		}, nil
	case Shipyard:
		return Building{
			Level:  1,
			Type:   buildingType,
			Effect: config.Effect,
		}, nil

	case Woodcutter:
		return Building{
			Level:   1,
			Type:    buildingType,
			Farming: config.Farming,
		}, nil
		// todo: refactor
	case UnitLimitHouse:
		return Building{
			Level:           1,
			Type:            buildingType,
			UnitLimit:       config.UnitLimit,
			StorageCapacity: config.StorageCapacity,
		}, nil
		// todo: refactor
	case Warehouse:
		return Building{
			Level:           1,
			Type:            buildingType,
			UnitLimit:       config.UnitLimit,
			StorageCapacity: config.StorageCapacity,
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
