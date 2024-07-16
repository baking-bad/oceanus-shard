package query

import (
	comp "oceanus-shard/component"
	"oceanus-shard/system"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
)

type GlobalMapRequest struct{}

type GlobalMapResponse struct {
	Player    string        `json:"player"`
	Island    ResourcePoint `json:"island"`
	Shipwreck ResourcePoint `json:"shipwreck"`
}

type ResourcePoint struct {
	Position  [2]float64      `json:"position"`
	Resources []comp.Resource `json:"resources"`
}

func GlobalMap(world cardinal.WorldContext, _ *GlobalMapRequest) (*[]GlobalMapResponse, error) {
	playerEntityIDs, players, _ := system.QueryAllComponents[comp.Player](
		world,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
		filter.Component[comp.Position](),
	)

	response := make([]GlobalMapResponse, len(playerEntityIDs))

	for i, islandEntityID := range playerEntityIDs {
		playerResources, _ := cardinal.GetComponent[comp.PlayerResources](world, islandEntityID)
		shipwreckResources, _ := cardinal.GetComponent[comp.ShipwreckResources](world, islandEntityID)
		position, _ := cardinal.GetComponent[comp.Position](world, islandEntityID)

		response[i] = GlobalMapResponse{
			Player: players[i].Nickname,
			Island: ResourcePoint{
				Position:  position.Island,
				Resources: playerResources.Resources,
			},
			Shipwreck: ResourcePoint{
				Position:  position.Shipwreck,
				Resources: shipwreckResources.Resources,
			},
		}
	}
	return &response, nil
}
