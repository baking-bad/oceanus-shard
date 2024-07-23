package system

import (
	"math"
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

			playerEntityID, playerResources, _ := QueryPlayerComponent[comp.PlayerResources](
				world,
				playerComponent.Nickname,
				filter.Component[comp.Player](),
				filter.Component[comp.PlayerResources](),
			)

			_, playerBuildings, err := QueryAllPlayerComponents[comp.Building](
				world,
				playerComponent.Nickname,
				filter.Component[comp.Player](),
				filter.Component[comp.Building](),
			)

			totalCapacity := 0
			for _, building := range playerBuildings {
				totalCapacity += building.StorageCapacity
			}

			if err != nil {
				return true
			}

			for i := range playerResources.Resources {
				if playerResources.Resources[i].Type == farmingComponent.Type {
					tickFarmedAmount := farmingComponent.Speed * constants.TickRate.Seconds() / time.Minute.Seconds()
					if playerResources.Resources[i].Amount <= float64(totalCapacity) {
						playerResources.Resources[i].Amount = math.Min(
							playerResources.Resources[i].Amount+tickFarmedAmount,
							float64(totalCapacity),
						)
					}
					if err := cardinal.SetComponent[comp.PlayerResources](world, playerEntityID, playerResources); err != nil {
						return true
					}
				}
			}

			return true
		})
}
