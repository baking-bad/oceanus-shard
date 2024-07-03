package system

import (
	"errors"
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

			if err := SubtractResourcesToBuild(world, building, request.Tx.PersonaTag); err != nil {
				return msg.CreateBuildingResult{Success: false}, err
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

func SubtractResourcesToBuild(world cardinal.WorldContext, building comp.Building, personaTag string) error {
	playerResourcesEntityID, playerResources, _ := QueryComponent[comp.PlayerResources](
		world,
		personaTag,
		filter.Component[comp.Player](),
		filter.Component[comp.PlayerResources](),
	)

	resourcesToBuild := comp.BuildingConfigs[building.Type].Resources
	for _, resource := range resourcesToBuild {
		var playerResource *comp.Resource
		var err error
		if playerResource, err = GetResourceByType(playerResources, resource.Type); err != nil {
			return fmt.Errorf("can't get player resource %s: %w", resource.Type, err)
		}
		if playerResource.Amount < resource.Amount {
			return fmt.Errorf("not enough resource %s", resource.Type)
		}

		playerResource.Amount -= resource.Amount
		SetResourceByType(playerResources, *playerResource)

		if err := cardinal.SetComponent(world, playerResourcesEntityID, playerResources); err != nil {
			return fmt.Errorf("failed to update player resource: %w", err)
		}
	}

	return nil
}

func GetResourceByType(playerResources *comp.PlayerResources, resourceType comp.ResourceType) (*comp.Resource, error) {
	for i := range playerResources.Resources {
		if playerResources.Resources[i].Type == resourceType {
			return &playerResources.Resources[i], nil
		}
	}
	return nil, errors.New("resource not found")
}

func SetResourceByType(playerResources *comp.PlayerResources, newResource comp.Resource) {
	for i := range playerResources.Resources {
		if playerResources.Resources[i].Type == newResource.Type {
			playerResources.Resources[i] = newResource
		}
	}
}
