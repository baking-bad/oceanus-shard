package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// CreateShipSystem -
func CreateShipSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreateShipMsg, msg.CreateShipResult](world,
		func(request cardinal.TxData[msg.CreateShipMsg]) (msg.CreateShipResult, error) {
			playerMapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
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
				return msg.CreateShipResult{Success: false},
					fmt.Errorf("failed to create ship, this player did not have tilemap")
			}

			tileMapPlayer, _ := cardinal.GetComponent[comp.Player](world, playerMapEntityID)
			if tileMapPlayer.Nickname != request.Tx.PersonaTag {
				return msg.CreateShipResult{Success: false}, fmt.Errorf("can't use another player map")
			}

			if request.Msg.TileIndex < 0 || request.Msg.TileIndex >= len(*playerMap.Tiles) {
				return msg.CreateShipResult{Success: false}, fmt.Errorf("index of tiles out of range")
			}

			tile := &(*playerMap.Tiles)[request.Msg.TileIndex]
			if tile.Building == nil {
				return msg.CreateShipResult{Success: false},
					fmt.Errorf("failed to create ship, this tile didn't have buildings")
			}

			targetBuilding, buildingIndex, err := FindBuildingByTileID(playerBuildings, tile.Building.TileID)
			if err != nil {
				return msg.CreateShipResult{Success: false}, err
			}

			shipType := comp.ShipType(request.Msg.ShipType)

			if comp.BuildingConfigs[targetBuilding.Type].ShipsInfo == nil {
				return msg.CreateShipResult{Success: false},
					fmt.Errorf("this building can't build ships")
			}

			if (*comp.BuildingConfigs[targetBuilding.Type].ShipsInfo)[shipType].BuildingType != targetBuilding.Type {
				return msg.CreateShipResult{Success: false},
					fmt.Errorf("this building can't build ships")
			}

			maxShipAmount := comp.GetMaxShips(playerBuildings, shipType)

			_, playerShips, _ := QueryAllPlayerComponents[comp.Ship](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.Ship](),
			)

			currentShipsAmount := 0
			for _, ship := range playerShips {
				if ship.Type == shipType {
					currentShipsAmount++
				}
			}

			if maxShipAmount >= currentShipsAmount {
				return msg.CreateShipResult{Success: false},
					fmt.Errorf("can't create more than %d ships", maxShipAmount)
			}

			targetBuilding.BuildingShip = &comp.BuildingShip{
				Type:                  shipType,
				BuildingTimeSeconds:   (*comp.GetAllBuildingShipConstants())[shipType].BuildingTimeSeconds,
				BuildingTimeStartedAt: world.Timestamp(),
			}

			tile.Building = targetBuilding
			if err := cardinal.SetComponent(world, playerBuildingsEntityIDs[buildingIndex], targetBuilding); err != nil {
				return msg.CreateShipResult{Success: false}, fmt.Errorf("failed to create ship: %w", err)
			}

			resourcesForShipBuilding := (*comp.GetAllBuildingShipConstants())[shipType].BuildResources
			if err := SubtractResources(world, resourcesForShipBuilding, request.Tx.PersonaTag); err != nil {
				return msg.CreateShipResult{Success: false}, err
			}

			if err := cardinal.SetComponent(world, playerMapEntityID, playerMap); err != nil {
				return msg.CreateShipResult{Success: false}, fmt.Errorf("failed to create ship: %w", err)
			}

			return msg.CreateShipResult{Success: true}, nil
		})
}
