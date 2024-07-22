package system

import (
	"errors"
	"fmt"
	comp "oceanus-shard/component"
	"reflect"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

func QueryPlayerComponent[T types.Component](
	world cardinal.WorldContext,
	targetNickname string,
	components ...filter.ComponentWrapper,
) (types.EntityID, *T, error) {
	var entityID types.EntityID
	var targetComponent *T
	var err error

	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(components...)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				entityID = id
				targetComponent, err = cardinal.GetComponent[T](world, id)
				if err != nil {
					return false
				}
				return false
			}

			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if targetComponent == nil {
		return 0, nil, fmt.Errorf("component %s on %q does not exist", reflect.TypeOf(targetComponent).Name(), targetNickname)
	}

	return entityID, targetComponent, err
}

func QueryAllPlayerComponents[T types.Component](
	world cardinal.WorldContext,
	targetNickname string,
	components ...filter.ComponentWrapper,
) ([]types.EntityID, []*T, error) {
	var entityIDs []types.EntityID
	var targetComponents []*T
	var err error

	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(components...)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			if player.Nickname == targetNickname {
				entityIDs = append(entityIDs, id)
				targetComponent, _ := cardinal.GetComponent[T](world, id)
				targetComponents = append(targetComponents, targetComponent)
			}

			return true
		})
	if searchErr != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}
	return entityIDs, targetComponents, err
}

func QueryAllComponents[T types.Component](
	world cardinal.WorldContext,
	components ...filter.ComponentWrapper,
) ([]types.EntityID, []*T, error) {
	var entityIDs []types.EntityID
	var targetComponents []*T
	var err error

	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(components...)).Each(world,
		func(id types.EntityID) bool {
			if err != nil {
				return false
			}

			entityIDs = append(entityIDs, id)
			targetComponent, _ := cardinal.GetComponent[T](world, id)
			targetComponents = append(targetComponents, targetComponent)
			return true
		})
	if searchErr != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}
	return entityIDs, targetComponents, err
}

func FindBuildingByTileID(buildings []*comp.Building, tileID int) (*comp.Building, int, error) {
	for index, building := range buildings {
		if building.TileID == tileID {
			return building, index, nil
		}
	}
	return nil, -1, errors.New("building with given tileID not found")
}

func SubtractResources(world cardinal.WorldContext, resources []comp.Resource, personaTag string) error {
	playerResourcesEntityID, playerResources, _ := QueryPlayerComponent[comp.PlayerResources](
		world,
		personaTag,
		filter.Component[comp.Player](),
		filter.Component[comp.PlayerResources](),
	)

	for _, resource := range resources {
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

func GetTotalPlayersAmount(world cardinal.WorldContext) (int, error) {
	totalPlayers := 0
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.TileMap]())).Each(world,
		func(_ types.EntityID) bool {
			totalPlayers++
			return true
		})
	if searchErr != nil {
		return 0, searchErr
	}
	return totalPlayers, searchErr
}

func updateEffect(world cardinal.WorldContext, buildingEntityID types.EntityID, effect *comp.Effect) error {

	building, _ := cardinal.GetComponent[comp.Building](world, buildingEntityID)
	player, _ := cardinal.GetComponent[comp.Player](world, buildingEntityID)
	mapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
		world,
		player.Nickname,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
	)

	building.Effect = effect
	tile := &(*playerMap.Tiles)[building.TileID]
	tile.Building = building
	if err := cardinal.SetComponent(world, buildingEntityID, effect); err != nil {
		return err
	}
	if err := cardinal.SetComponent(world, buildingEntityID, building); err != nil {
		return err
	}
	if err := cardinal.SetComponent(world, mapEntityID, playerMap); err != nil {
		return err
	}
	return nil
}
