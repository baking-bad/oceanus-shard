package system

import (
	"fmt"
	comp "oceanus-shard/component"
	"reflect"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

func QueryComponent[T types.Component](
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

func QueryAllComponents[T types.Component](
	world cardinal.WorldContext,
	targetNickname string,
	components ...filter.ComponentWrapper,
) (types.EntityID, []*T, error) {
	var entityID types.EntityID
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
				entityID = id
				targetComponent, _ := cardinal.GetComponent[T](world, id)
				targetComponents = append(targetComponents, targetComponent)
			}

			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	return entityID, targetComponents, err
}
