package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// StartSailShipwreckSystem -
func StartSailShipwreckSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.SailShipwreckMsg, msg.SailShipWreckResult](
		world,
		func(request cardinal.TxData[msg.SailShipwreckMsg]) (msg.SailShipWreckResult, error) {
			mapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			_, targetPlayerPosition, err := QueryPlayerComponent[comp.Position](
				world,
				request.Msg.Player,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			if err != nil {
				return msg.SailShipWreckResult{Success: false}, err
			}

			playerBuildingsEntityIDs, playerBuildings, _ := QueryAllPlayerComponents[comp.Building](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.Building](),
			)

			if playerMap == nil {
				return msg.SailShipWreckResult{Success: false},
					fmt.Errorf("failed to sail shipwreck, this player did not have tilemap")
			}

			player, _ := cardinal.GetComponent[comp.Player](world, mapEntityID)
			playerPosition, _ := cardinal.GetComponent[comp.Position](world, mapEntityID)
			if player.Nickname != request.Tx.PersonaTag {
				return msg.SailShipWreckResult{Success: false}, fmt.Errorf("can't sail another player building")
			}

			var shipyard *comp.Building
			var shipyardEntityID types.EntityID
			for playerBuildingID, playerBuilding := range playerBuildings {
				if playerBuilding.Type != comp.Shipyard {
					continue
				}
				if playerBuilding.Effect.Amount == 0 {
					continue
				}
				if playerBuilding.Effect.Position != playerPosition.Island {
					continue
				}
				shipyard = playerBuilding
				shipyardEntityID = playerBuildingsEntityIDs[playerBuildingID]
			}

			if shipyard == nil {
				return msg.SailShipWreckResult{Success: false}, fmt.Errorf("player didn't have ships on base")
			}

			shipyard.Effect.TargetPosition = &targetPlayerPosition.Shipwreck
			tile := &(*playerMap.Tiles)[shipyard.TileID]
			tile.Building = shipyard

			if err := cardinal.SetComponent(world, shipyardEntityID, shipyard); err != nil {
				return msg.SailShipWreckResult{Success: false}, err
			}

			if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
				return msg.SailShipWreckResult{Success: false}, err
			}

			return msg.SailShipWreckResult{Success: true}, nil
		})
}
