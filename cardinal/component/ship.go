package component

import "oceanus-shard/constants"

type ShipType string

const (
	Raft ShipType = "Raft"
)

func GetAllShipTypes() []ShipType {
	return []ShipType{
		Raft,
	}
}

func GetAllBuildingShipConstants() *map[ShipType]BuildingShipConstants {
	ships := make(map[ShipType]BuildingShipConstants)
	for _, shipType := range GetAllShipTypes() {
		if shipType == Raft {
			ships[shipType] = BuildingShipConstants{
				Type:         shipType,
				BuildingType: Shipyard,
				MaxAmount:    constants.ShipyardShipRaftCapacity,
				BuildResources: []Resource{
					{Type: Wood, Amount: constants.ShipRaftResourceWood},
					{Type: Fish, Amount: constants.ShipRaftResourceFish},
				},
				BuildingTimeSeconds: constants.ShipRaftBuildSeconds,
			}
		}
	}
	return &ships
}

func GetMaxShips(playerBuildings []*Building, shipType ShipType) int {
	shipConstants := (*GetAllBuildingShipConstants())[shipType]
	maxShipsAmountByOneBuilding := shipConstants.MaxAmount
	buildingsAmountCanBuildShip := 0
	for _, playerBuilding := range playerBuildings {
		if playerBuilding.Type == shipConstants.BuildingType {
			buildingsAmountCanBuildShip++
		}
	}

	return maxShipsAmountByOneBuilding * buildingsAmountCanBuildShip
}

type Ship struct {
	Type      ShipType   `json:"type"`
	PositionX float64    `json:"positionX"`
	PositionY float64    `json:"positionY"`
	Resources []Resource `json:"resources,omitempty"`
}

func (Ship) Name() string {
	return "Ship"
}
