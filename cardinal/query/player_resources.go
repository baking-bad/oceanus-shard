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
	Effects   []comp.Effect   `json:"effects"`
	Farming   []comp.Farming  `json:"farming"`
}

func PlayerResources(world cardinal.WorldContext, req *PlayerResourcesRequest) (*PlayerResourcesResponse, error) {
	_, playerResources, _ := system.QueryComponent[comp.PlayerResources](
		world,
		req.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.PlayerResources](),
	)

	_, farmingComponents, err := system.QueryAllComponents[comp.Farming](
		world,
		req.Nickname,
		filter.Component[comp.Building](),
		filter.Component[comp.Farming](),
	)

	aggregatedFarmingMap := make(map[comp.ResourceType]float32)
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

	if playerResources == nil {
		return nil, fmt.Errorf("error querying player %s resources", req.Nickname)
	}

	effectsMap := make(map[comp.EffectType]int)
	capacity := 0
	for _, building := range *GetAllBuildings() {
		if building.Effect != nil {
			effectsMap[building.Effect.Type] += building.Effect.Amount
		}

		capacity += building.StorageCapacity
	}

	resourcesResponse := make([]comp.Resource, 0, len(playerResources.Resources))
	for _, resource := range playerResources.Resources {
		resource.Capacity = capacity
		resourcesResponse = append(resourcesResponse, resource)
	}

	effectsResponse := make([]comp.Effect, 0, len(playerResources.Effects))
	for _, effect := range playerResources.Effects {
		effect.Capacity = effectsMap[effect.Type]
		effectsResponse = append(effectsResponse, effect)
	}

	return &PlayerResourcesResponse{
		Resources: resourcesResponse,
		Effects:   effectsResponse,
		Farming:   aggregatedfarmingSlice,
	}, err
}
