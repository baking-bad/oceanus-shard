package query

import (
	comp "oceanus-shard/component"
	"oceanus-shard/system"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
)

type MapStateRequest struct {
	Nickname string
}

type MapStateResponse struct {
	Tiles  *[]comp.Tile `json:"tiles"`
	Width  int          `json:"width"`
	Height int          `json:"height"`
}

func PlayerMap(world cardinal.WorldContext, req *MapStateRequest) (*MapStateResponse, error) {
	_, playerMap, err := system.QueryComponent[comp.TileMap](
		world,
		req.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
	)
	return &MapStateResponse{
		Tiles:  playerMap.Tiles,
		Width:  playerMap.Width,
		Height: playerMap.Height,
	}, err
}
