package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

// CreateBuildingSystem spawns buildings.
func CreateBuildingSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreateBuildingMsg, msg.CreateBuildingResult](
		world,
		func(request cardinal.TxData[msg.CreateBuildingMsg]) (msg.CreateBuildingResult, error) {
			mapEntityID, playerMap, _ := QueryComponent[comp.TileMap](
				world,
				request.Tx.PersonaTag,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
			)

			building, _ := comp.GetBuilding(comp.BuildingType(request.Msg.BuildingType))
			tiles := *playerMap.Tiles

			if request.Msg.TileIndex < 0 || request.Msg.TileIndex >= len(tiles) {
				return msg.CreateBuildingResult{Success: false}, fmt.Errorf("index of tiles out of range")
			}

			tile := &(*playerMap.Tiles)[request.Msg.TileIndex]
			if tile.Building == nil {
				tile.Building = &building
			} else {
				return msg.CreateBuildingResult{Success: false},
					fmt.Errorf("failed to create building, this tile already have another building")
			}

			player, _ := cardinal.GetComponent[comp.Player](world, mapEntityID)

			if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
				return msg.CreateBuildingResult{Success: false}, fmt.Errorf("failed to create building: %w", err)
			}

			playerEntityID, _ := cardinal.Create(world,
				player,
				comp.Building{
					Level:           building.Level,
					Type:            building.Type,
					Farming:         building.Farming,
					Effect:          building.Effect,
					UnitLimit:       building.UnitLimit,
					StorageCapacity: building.StorageCapacity,
				},
			)

			if building.Farming != nil {
				farmingComponent := &comp.Farming{
					Type:  building.Farming.Type,
					Speed: building.Farming.Speed,
				}
				_ = cardinal.AddComponentTo[comp.Farming](world, playerEntityID)
				_ = cardinal.SetComponent(world, playerEntityID, farmingComponent)
			}

			if building.Effect != nil {
				effectComponent := &comp.Effect{
					Type:   building.Effect.Type,
					Amount: building.Effect.Amount,
				}
				_ = cardinal.AddComponentTo[comp.Effect](world, playerEntityID)
				_ = cardinal.SetComponent(world, playerEntityID, effectComponent)
			}

			return msg.CreateBuildingResult{Success: true}, nil
		})
}
