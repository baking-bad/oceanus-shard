package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// DeleteBuildingSystem removes buildings.
func DeleteBuildingSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.DeleteBuildingMsg, msg.DeleteBuildingResult](
		world,
		func(request cardinal.TxData[msg.DeleteBuildingMsg]) (msg.DeleteBuildingResult, error) {
			mapEntityID, playerMap, _ := QueryComponent[comp.TileMap](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			buildingsEntityIDs, playerBuildings, _ := QueryAllComponents[comp.Building](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.Building](),
			)

			if playerMap == nil {
				return msg.DeleteBuildingResult{Success: false},
					fmt.Errorf("failed to delete building, this player did not have tilemap")
			}

			tileMapPlayer, _ := cardinal.GetComponent[comp.Player](world, mapEntityID)
			if tileMapPlayer.Nickname != request.Tx.PersonaTag {
				return msg.DeleteBuildingResult{Success: false}, fmt.Errorf("can't delete another player building")
			}

			tiles := *playerMap.Tiles

			if request.Msg.TileIndex < 0 || request.Msg.TileIndex >= len(tiles) {
				return msg.DeleteBuildingResult{Success: false}, fmt.Errorf("index of tiles out of range")
			}

			tile := &(*playerMap.Tiles)[request.Msg.TileIndex]
			if tile.Building == nil {
				return msg.DeleteBuildingResult{Success: false},
					fmt.Errorf("failed to delete building, this tile didn't have buildings")
			}
			for i, building := range playerBuildings {
				if building.TileID == tile.Building.TileID {
					if err := cardinal.Remove(world, buildingsEntityIDs[i]); err != nil {
						return msg.DeleteBuildingResult{Success: false}, fmt.Errorf("failed to delete building entity: %w", err)
					}
				}
			}
			tile.Building = nil

			if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
				return msg.DeleteBuildingResult{Success: false}, fmt.Errorf("failed to delete building from tile: %w", err)
			}

			return msg.DeleteBuildingResult{Success: true}, nil
		})
}
