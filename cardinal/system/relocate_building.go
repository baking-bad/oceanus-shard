package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// RelocateBuildingSystem relocates buildings.
func RelocateBuildingSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.RelocateBuildingMsg, msg.RelocateBuildingResult](
		world,
		func(request cardinal.TxData[msg.RelocateBuildingMsg]) (msg.RelocateBuildingResult, error) {
			mapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			playerBuildingsEntityIDs, playerBuildings, _ := QueryAllPlayerComponents[comp.Building](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.Building](),
			)

			if playerMap == nil {
				return msg.RelocateBuildingResult{Success: false},
					fmt.Errorf("failed to relocate building, this player did not have tilemap")
			}

			player, _ := cardinal.GetComponent[comp.Player](world, mapEntityID)
			if player.Nickname != request.Tx.PersonaTag {
				return msg.RelocateBuildingResult{Success: false}, fmt.Errorf("can't relocate another player building")
			}

			tiles := *playerMap.Tiles

			if request.Msg.TileIndexFrom < 0 || request.Msg.TileIndexFrom >= len(tiles) ||
				request.Msg.TileIndexTo < 0 || request.Msg.TileIndexTo >= len(tiles) {
				return msg.RelocateBuildingResult{Success: false}, fmt.Errorf("index of tiles out of range")
			}

			tileFrom := &(*playerMap.Tiles)[request.Msg.TileIndexFrom]
			if tileFrom.Building == nil {
				return msg.RelocateBuildingResult{Success: false},
					fmt.Errorf("failed to relocate building, source tile didn't have buildings")
			}

			tileTo := &(*playerMap.Tiles)[request.Msg.TileIndexTo]
			if tileTo.Building != nil {
				return msg.RelocateBuildingResult{Success: false},
					fmt.Errorf("failed to relocate building, target tile already have buildings")
			}

			if tileTo.Tile != comp.BuildingConfigs[tileFrom.Building.Type].TileType {
				return msg.RelocateBuildingResult{Success: false},
					fmt.Errorf("failed to relocate building, target tile type and building tiletype mismatch")
			}

			fromBuilding, fromBuildingIndex, err := FindBuildingByTileID(playerBuildings, tileFrom.Building.TileID)
			if err != nil {
				return msg.RelocateBuildingResult{Success: false}, err
			}

			fromBuilding.TileID = request.Msg.TileIndexTo
			tileFrom.Building = nil
			tileTo.Building = fromBuilding

			if err := cardinal.SetComponent(
				world,
				playerBuildingsEntityIDs[fromBuildingIndex],
				fromBuilding); err != nil {
				return msg.RelocateBuildingResult{Success: false}, err
			}

			if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
				return msg.RelocateBuildingResult{Success: false}, err
			}

			return msg.RelocateBuildingResult{Success: true}, nil
		})
}
