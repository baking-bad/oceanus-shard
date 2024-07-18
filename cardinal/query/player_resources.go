package query

import (
	"fmt"
	comp "oceanus-shard/component"
	"oceanus-shard/system"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
)

type PlayerResourcesRequest struct {
	Nickname string
}

type PlayerResourcesResponse struct {
	Resources []comp.Resource `json:"resources"`
	Ships     []comp.Ship     `json:"ships"`
	Farming   []comp.Farming  `json:"farming"`
}

func PlayerResources(world cardinal.WorldContext, req *PlayerResourcesRequest) (*PlayerResourcesResponse, error) {
	_, playerResources, _ := system.QueryPlayerComponent[comp.PlayerResources](
		world,
		req.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.PlayerResources](),
	)

	_, farmingComponents, _ := system.QueryAllPlayerComponents[comp.Farming](
		world,
		req.Nickname,
		filter.Component[comp.Building](),
		filter.Component[comp.Farming](),
	)

	_, allBuildings, _ := system.QueryAllPlayerComponents[comp.Building](
		world,
		req.Nickname,
		filter.Component[comp.Building](),
		filter.Component[comp.Player](),
	)

	if playerResources == nil {
		return nil, fmt.Errorf("error querying player %s resources", req.Nickname)
	}

	aggregatedFarmingMap := make(map[comp.ResourceType]float64)
	for _, farming := range farmingComponents {
		aggregatedFarmingMap[farming.Type] += farming.Speed
	}

	aggregatedfarmingSlice := make([]comp.Farming, 0, len(aggregatedFarmingMap))
	for resourceType, speed := range aggregatedFarmingMap {
		aggregatedfarmingSlice = append(aggregatedfarmingSlice, comp.Farming{
			Type:  resourceType,
			Speed: speed,
		})
	}

	capacity := 0
	for _, building := range allBuildings {
		capacity += building.StorageCapacity
	}

	resourcesResponse := make([]comp.Resource, 0, len(playerResources.Resources))
	for _, resource := range playerResources.Resources {
		resource.Capacity = capacity
		resourcesResponse = append(resourcesResponse, resource)
	}

	_, allShips, err := system.QueryAllPlayerComponents[comp.Ship](
		world,
		req.Nickname,
		filter.Component[comp.Ship](),
		filter.Component[comp.Player](),
	)

	allShipsResponse := make([]comp.Ship, 0, len(allShips))
	for _, ship := range allShips {
		allShipsResponse = append(allShipsResponse, *ship)
	}

	return &PlayerResourcesResponse{
		Resources: resourcesResponse,
		Ships:     allShipsResponse,
		Farming:   aggregatedfarmingSlice,
	}, err
}
