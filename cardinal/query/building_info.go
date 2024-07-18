package query

import (
	comp "oceanus-shard/component"

	"pkg.world.dev/world-engine/cardinal"
)

type BuildingsInfoRequest struct {
}

type BuildingInfoResponse struct {
	Building        comp.BuildingType            `json:"building"`
	TileType        comp.TileType                `json:"tileType"`
	Resources       []comp.Resource              `json:"resources"`
	ShipsInfo       []comp.BuildingShipConstants `json:"shipsInfo,omitempty"`
	Farming         *comp.Farming                `json:"farming,omitempty"`
	UnitLimit       int                          `json:"unitLimit,omitempty"`
	StorageCapacity int                          `json:"storageCapacity,omitempty"`
}

func AllBuildings(_ cardinal.WorldContext, _ *BuildingsInfoRequest) (*[]BuildingInfoResponse, error) {
	return GetAllBuildings(), nil
}

func GetAllBuildings() *[]BuildingInfoResponse {
	buildings := make([]BuildingInfoResponse, len(comp.GetAllBuildingTypes()))
	for i, buildingType := range comp.GetAllBuildingTypes() {
		buildingConf := comp.BuildingConfigs[buildingType]

		buildings[i] = BuildingInfoResponse{
			Building:        buildingType,
			TileType:        buildingConf.TileType,
			Resources:       buildingConf.Resources,
			ShipsInfo:       mapToSlice(buildingConf.ShipsInfo),
			Farming:         buildingConf.Farming,
			UnitLimit:       buildingConf.UnitLimit,
			StorageCapacity: buildingConf.StorageCapacity,
		}
	}

	return &buildings
}

func mapToSlice(m *map[comp.ShipType]comp.BuildingShipConstants) []comp.BuildingShipConstants {
	if m == nil {
		return nil
	}
	result := make([]comp.BuildingShipConstants, 0, len(*m))
	for _, value := range *m {
		result = append(result, value)
	}
	return result
}
