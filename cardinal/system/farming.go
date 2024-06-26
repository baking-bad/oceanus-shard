package system

import (
	"time"

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

			playerEntityID, playerResources, err := QueryComponent[comp.PlayerResources](
				world,
				playerComponent.Nickname,
				filter.Component[comp.Player](),
				filter.Component[comp.PlayerResources](),
			)

			if err != nil {
				return true
			}

			for i := range playerResources.Resources {
				if playerResources.Resources[i].Type == farmingComponent.Type {
					playerResources.Resources[i].Amount +=
						farmingComponent.Speed * float32(constants.TickRate.Seconds()) / float32(time.Minute.Seconds())

					if err := cardinal.SetComponent[comp.PlayerResources](world, playerEntityID, playerResources); err != nil {
						return true
					}
				}
			}

			return true
		})
}
