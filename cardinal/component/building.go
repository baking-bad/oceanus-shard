package component

import (
	"fmt"
	"oceanus-shard/constants"
)

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

type BuildingShip struct {
	Type                  ShipType `json:"shipType"`
	BuildingTimeSeconds   int      `json:"buildingTimeSeconds"`
	BuildingTimeStartedAt uint64   `json:"buildingTimeStartedAt"`
}

type Building struct {
	TileID          int           `json:"tileID"`
	Level           int           `json:"level"`
	Type            BuildingType  `json:"type"`
	Farming         *Farming      `json:"farming,omitempty"`
	BuildingShip    *BuildingShip `json:"buildingShip,omitempty"`
	UnitLimit       int           `json:"unitLimit,omitempty"`
	StorageCapacity int           `json:"storageCapacity,omitempty"`
}

func (Building) Name() string {
	return "Building"
}

type BuildingShipConstants struct {
	Type                ShipType     `json:"type"`
	BuildingType        BuildingType `json:"buildingType"`
	MaxAmount           int          `json:"maxAmount"`
	BuildResources      []Resource   `json:"buildResources"`
	BuildingTimeSeconds int          `json:"buildingTimeSeconds"`
}

type BuildingConstants struct {
	TileID          int
	UnitLimit       int
	StorageCapacity int
	Resources       []Resource
	Farming         *Farming
	ShipsInfo       *map[ShipType]BuildingShipConstants
	TileType        TileType
}

var BuildingConfigs = map[BuildingType]BuildingConstants{
	Main: {
		TileID:          constants.MainTileID,
		Resources:       []Resource{},
		UnitLimit:       constants.MainUnitLimit,
		StorageCapacity: constants.MainStorageCapacity,
		TileType:        GenericTile,
	},
	Woodcutter: {
		Resources: []Resource{},
		Farming: &Farming{
			Type:  Wood,
			Speed: constants.WoodcutterFarmingSpeed,
		},
		TileType: WoodTile,
	},
	Quarry: {
		Resources: []Resource{
			{Type: Wood, Amount: constants.QuarryResourcesWoodAmount},
		},
		Farming: &Farming{
			Type:  Stone,
			Speed: constants.QuarryFarmingStoneSpeed,
		},
		TileType: StoneTile,
	},
	FishermanHut: {
		Resources: []Resource{
			{Type: Wood, Amount: constants.FishermanHutResourcesWoodAmount},
			{Type: Stone, Amount: constants.FishermanHutResourcesStoneAmount},
		},
		Farming: &Farming{
			Type:  Fish,
			Speed: constants.FishermanHutFarmingFishSpeed,
		},
		TileType: CoastlineTile,
	},
	Shipyard: {
		Resources: []Resource{
			{Type: Wood, Amount: constants.ShipyardResourcesWoodAmount},
			{Type: Stone, Amount: constants.ShipyardResourcesStoneAmount},
			{Type: Fish, Amount: constants.ShipyardResourcesFishAmount},
		},
		ShipsInfo: GetAllBuildingShipConstants(),
		TileType:  CoastlineTile,
	},
	Warehouse: {
		Resources: []Resource{
			{Type: Wood, Amount: constants.WarehouseResourcesWoodAmount},
			{Type: Fish, Amount: constants.WarehouseResourcesFishAmount},
		},
		TileType:        GenericTile,
		StorageCapacity: constants.WarehouseStorageCapacity,
	},
	// todo: no build!
	UnitLimitHouse: {
		Resources: []Resource{
			{Type: Wood, Amount: constants.UnitLimitHouseResourcesWoodAmount},
			{Type: Stone, Amount: constants.UnitLimitHouseResourcesStoneAmount},
		},
		TileType:  GenericTile,
		UnitLimit: constants.UnitLimitHouseUnitLimit,
	},
}

func GetBuilding(buildingType BuildingType) (Building, error) {
	config := BuildingConfigs[buildingType]

	switch buildingType {
	case Main:
		return Building{
			TileID:          config.TileID,
			Level:           1,
			Type:            buildingType,
			UnitLimit:       config.UnitLimit,
			StorageCapacity: config.StorageCapacity,
		}, nil
	case Woodcutter:
		return Building{
			Level:   1,
			Type:    buildingType,
			Farming: config.Farming,
		}, nil
	case Quarry:
		return Building{
			Level:   1,
			Type:    buildingType,
			Farming: config.Farming,
		}, nil
	case FishermanHut:
		return Building{
			Level:   1,
			Type:    buildingType,
			Farming: config.Farming,
		}, nil
	case Shipyard:
		return Building{
			Level: 1,
			Type:  buildingType,
			// Ship:  config.Ship,
		}, nil
	case Warehouse:
		return Building{
			Level:           1,
			Type:            buildingType,
			StorageCapacity: config.StorageCapacity,
		}, nil
	case UnitLimitHouse:
		return Building{
			Level:     1,
			Type:      buildingType,
			UnitLimit: config.UnitLimit,
		}, nil
	default:
		return Building{
			Level: 0,
			Type:  buildingType,
		}, fmt.Errorf("%s is invalid building type", buildingType)
	}
}
