package query

import (
	"fmt"
	comp "oceanus-shard/component"
	"oceanus-shard/system"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
)

type MapStateRequest struct {
	Nickname string
}

type MapStateResponse struct {
	Tiles    *[]comp.Tile `json:"tiles"`
	Width    int          `json:"width"`
	Height   int          `json:"height"`
	Position [2]float64   `json:"position"`
}

func PlayerMap(world cardinal.WorldContext, req *MapStateRequest) (*MapStateResponse, error) {
	playerMapEntityID, playerMap, err := system.QueryPlayerComponent[comp.TileMap](
		world,
		req.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
	)
	if playerMap == nil {
		return nil, fmt.Errorf("error querying players %s map", req.Nickname)
	}
	position, _ := cardinal.GetComponent[comp.Position](world, playerMapEntityID)
	return &MapStateResponse{
		Tiles:    playerMap.Tiles,
		Width:    playerMap.Width,
		Height:   playerMap.Height,
		Position: position.Island,
	}, err
}
