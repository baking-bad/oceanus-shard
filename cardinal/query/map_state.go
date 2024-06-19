package query

import (
	"fmt"
	comp "oceanus-shard/component"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
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
	var playerMap *comp.TileMap
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.TileMap]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			if player.Nickname == req.Nickname {
				playerMap, err = cardinal.GetComponent[comp.TileMap](world, id)
				if err != nil {
					return false
				}
				return false
			}

			return true
		})
	if searchErr != nil {
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if playerMap == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &MapStateResponse{Tiles: playerMap.Tiles, Width: playerMap.Width, Height: playerMap.Height}, nil
}
