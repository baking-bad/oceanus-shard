package query

import (
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
}

func PlayerResources(world cardinal.WorldContext, req *PlayerResourcesRequest) (*PlayerResourcesResponse, error) {
	_, playerResources, err := system.QueryComponent[comp.PlayerResources](
		world,
		req.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.PlayerResources](),
	)
	return &PlayerResourcesResponse{
		Resources: playerResources.Resources,
		Effects:   playerResources.Effects,
	}, err
}
