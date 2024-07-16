package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// CreateEffectSystem spawns effects.
func CreateEffectSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreateEffectMsg, msg.CreateEffectResult](world,
		func(request cardinal.TxData[msg.CreateEffectMsg]) (msg.CreateEffectResult, error) {
			playerMapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			playerBuildingsEntityIDs, playerBuildingsWithEffect, _ := QueryAllPlayerComponents[comp.Building](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.Building](),
				filter.Component[comp.Effect](),
			)

			if playerMap == nil {
				return msg.CreateEffectResult{Success: false},
					fmt.Errorf("failed to create effect, this player did not have tilemap")
			}

			tileMapPlayer, _ := cardinal.GetComponent[comp.Player](world, playerMapEntityID)
			if tileMapPlayer.Nickname != request.Tx.PersonaTag {
				return msg.CreateEffectResult{Success: false}, fmt.Errorf("can't use another player map")
			}

			if request.Msg.TileIndex < 0 || request.Msg.TileIndex >= len(*playerMap.Tiles) {
				return msg.CreateEffectResult{Success: false}, fmt.Errorf("index of tiles out of range")
			}

			tile := &(*playerMap.Tiles)[request.Msg.TileIndex]
			if tile.Building == nil {
				return msg.CreateEffectResult{Success: false},
					fmt.Errorf("failed to create effect, this tile didn't have buildings")
			}

			if tile.Building.Effect == nil {
				return msg.CreateEffectResult{Success: false},
					fmt.Errorf("failed to create effect, this building didn't have effect")
			}

			building, i, err := FindBuildingByTileID(playerBuildingsWithEffect, tile.Building.TileID)
			if err != nil {
				return msg.CreateEffectResult{Success: false}, err
			}

			effectComponent, _ := cardinal.GetComponent[comp.Effect](world, playerBuildingsEntityIDs[i])
			effectComponent.BuildingTimeStartedAt = world.Timestamp()
			resourcesForEffect := comp.BuildingConfigs[building.Type].Effect.Resources
			if err := SubtractResources(world, resourcesForEffect, request.Tx.PersonaTag); err != nil {
				return msg.CreateEffectResult{Success: false}, err
			}
			if err := cardinal.SetComponent(world, playerBuildingsEntityIDs[i], effectComponent); err != nil {
				return msg.CreateEffectResult{Success: false}, fmt.Errorf("failed to create effect: %w", err)
			}

			tile.Building.Effect = effectComponent
			if err := cardinal.SetComponent(world, playerMapEntityID, playerMap); err != nil {
				return msg.CreateEffectResult{Success: false}, fmt.Errorf("failed to create effect: %w", err)
			}

			return msg.CreateEffectResult{Success: true}, nil
		})
}
