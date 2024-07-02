package query

import (
	comp "oceanus-shard/component"

	"pkg.world.dev/world-engine/cardinal"
)

type BuildingsInfoRequest struct {
}

type BuildingInfoResponse struct {
	Building        comp.BuildingType `json:"building"`
	TileType        comp.TileType     `json:"tileType"`
	Resources       []comp.Resource   `json:"resources"`
	Effect          *comp.Effect      `json:"effect,omitempty"`
	Farming         *comp.Farming     `json:"farming,omitempty"`
	UnitLimit       int               `json:"unitLimit,omitempty"`
	StorageCapacity int               `json:"storageCapacity,omitempty"`
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
			Effect:          buildingConf.Effect,
			Farming:         buildingConf.Farming,
			UnitLimit:       buildingConf.UnitLimit,
			StorageCapacity: buildingConf.StorageCapacity,
		}
	}

	return &buildings
}
