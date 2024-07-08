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

			if tiles[request.Msg.TileIndex].Tile != comp.BuildingConfigs[building.Type].TileType {
				return msg.CreateBuildingResult{Success: false},
					fmt.Errorf("failed to create building, this building doesn't fit this tiletype")
			}

			tile := &(*playerMap.Tiles)[request.Msg.TileIndex]
			if tile.Building == nil {
				building.TileID = request.Msg.TileIndex
				tile.Building = &building
			} else {
				return msg.CreateBuildingResult{Success: false},
					fmt.Errorf("failed to create building, this tile already have another building")
			}

			resourcesForBuild := comp.BuildingConfigs[building.Type].Resources
			if err := SubtractResources(world, resourcesForBuild, request.Tx.PersonaTag); err != nil {
				return msg.CreateBuildingResult{Success: false}, err
			}

			player, _ := cardinal.GetComponent[comp.Player](world, mapEntityID)

			if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
				return msg.CreateBuildingResult{Success: false}, fmt.Errorf("failed to create building: %w", err)
			}

			buildingEntityID, _ := cardinal.Create(world,
				player,
				building,
			)

			if building.Farming != nil {
				farmingComponent := &comp.Farming{
					Type:  building.Farming.Type,
					Speed: building.Farming.Speed,
				}
				_ = cardinal.AddComponentTo[comp.Farming](world, buildingEntityID)
				_ = cardinal.SetComponent(world, buildingEntityID, farmingComponent)
			}

			if building.Effect != nil {
				_ = cardinal.AddComponentTo[comp.Effect](world, buildingEntityID)
				_ = cardinal.SetComponent(world, buildingEntityID, building.Effect)
			}

			return msg.CreateBuildingResult{Success: true}, nil
		})
}
