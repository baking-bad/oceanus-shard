package query

import (
	comp "oceanus-shard/component"

	"pkg.world.dev/world-engine/cardinal"
)

type BuildingsInfoRequest struct {
}

type BuildingInfoResponse struct {
	Building  comp.BuildingType `json:"building"`
	TileType  comp.TileType     `json:"tileType"`
	Resources []comp.Resource   `json:"resources"`
	Effect    *comp.Effect      `json:"effect,omitempty"`
	Farming   *comp.Farming     `json:"farming,omitempty"`
}

func AllBuildings(world cardinal.WorldContext, req *BuildingsInfoRequest) (*[]BuildingInfoResponse, error) {
	var buildings []BuildingInfoResponse
	for _, buildingType := range comp.GetAllBuildingTypes() {
		buildingConf := comp.BuildingConfigs[buildingType]

		buildings = append(buildings, BuildingInfoResponse{
			Building:  buildingType,
			TileType:  buildingConf.TileType,
			Resources: buildingConf.Resources,
			Effect:    buildingConf.Effect,
			Farming:   buildingConf.Farming,
		})
	}
	return &buildings, nil
}
