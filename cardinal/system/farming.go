package system

import (
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "oceanus-shard/component"
	"oceanus-shard/constants"
)

func FarmingSystem(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Building](), filter.Component[comp.Farming]())).
		Each(world, func(id types.EntityID) bool {
			playerComponent, _ := cardinal.GetComponent[comp.Player](world, id)
			farmingComponent, _ := cardinal.GetComponent[comp.Farming](world, id)

			playerEntityID, playerResources, err := QueryPlayerResources(world, playerComponent.Nickname)
			if err != nil {
				return true
			}

			for i := range playerResources.Resources {
				if playerResources.Resources[i].Type == farmingComponent.Type {
					playerResources.Resources[i].Amount += farmingComponent.Speed * float32(constants.TickRate.Seconds()) / 60

					if err := cardinal.SetComponent[comp.PlayerResources](world, playerEntityID, playerResources); err != nil {
						return true
					}
				}
			}

			return true
		})
}
