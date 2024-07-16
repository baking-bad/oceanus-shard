package system

import (
	"time"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "oceanus-shard/component"
	"oceanus-shard/constants"
)

func EffectsSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Building](), filter.Component[comp.Effect]())).
		Each(
			world,
			func(id types.EntityID) bool {
				playerComponent, _ := cardinal.GetComponent[comp.Player](world, id)
				effectComponent, _ := cardinal.GetComponent[comp.Effect](world, id)
				buildingComponent, _ := cardinal.GetComponent[comp.Building](world, id)

				if effectComponent.BuildingTimeStartedAt == 0 {
					return true
				}

				playerMapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
					world,
					playerComponent.Nickname,
					filter.Component[comp.Player](),
					filter.Component[comp.PlayerResources](),
				)

				playerResourcesEntityID, playerResources, _ := QueryPlayerComponent[comp.PlayerResources](
					world,
					playerComponent.Nickname,
					filter.Component[comp.Player](),
					filter.Component[comp.PlayerResources](),
				)

				previousEffectAmount := effectComponent.Amount
				endBuildingTime := effectComponent.BuildingTimeStartedAt +
					uint64(time.Duration(effectComponent.BuildingTimeSeconds)*time.Second/time.Millisecond)

				if world.Timestamp() < endBuildingTime {
					return true
				}

				effectComponent.Amount = min(
					constants.ShipyardEffectRaftCapacity,
					effectComponent.Amount+1,
				)
				effectComponent.BuildingTimeStartedAt = 0
				buildingComponent.Effect = effectComponent
				tile := &(*playerMap.Tiles)[buildingComponent.TileID]
				tile.Building = buildingComponent

				for _, playerResourcesEffect := range playerResources.Effects {
					if playerResourcesEffect.Type != effectComponent.Type {
						continue
					}

					totalEffectsAmount := playerResourcesEffect.Amount + effectComponent.Amount - previousEffectAmount
					playerResources.Effects = UpdateEffectsAmount(playerResources.Effects, effectComponent.Type, totalEffectsAmount)
					if err := cardinal.SetComponent(world, playerResourcesEntityID, playerResources); err != nil {
						return true
					}
				}

				if err := cardinal.SetComponent(world, id, effectComponent); err != nil {
					return true
				}
				if err := cardinal.SetComponent(world, id, buildingComponent); err != nil {
					return true
				}
				if err := cardinal.SetComponent(world, playerMapEntityID, playerMap); err != nil {
					return true
				}

				return true
			},
		)
}

func UpdateEffectsAmount(effects []comp.Effect, effectType comp.EffectType, newAmount int) []comp.Effect {
	for i := range effects {
		if effects[i].Type == effectType {
			effects[i].Amount = newAmount
			return effects
		}
	}
	return effects
}
