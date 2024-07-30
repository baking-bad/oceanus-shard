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

// SailShip -
func SailShip(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Building](), filter.Component[comp.Effect]())).
		Each(world, func(id types.EntityID) bool {
			effectComponent, _ := cardinal.GetComponent[comp.Effect](world, id)
			if effectComponent.Amount == 0 || effectComponent.TargetPosition == nil {
				return true
			}

			buildingComponent, _ := cardinal.GetComponent[comp.Building](world, id)
			raftTravelDistancePerTick :=
				constants.RaftTravelSpeedPerMinute / time.Minute.Seconds() * constants.TickRate.Seconds()

			buildingComponent.Effect.Position = findPointAtDistance(
				buildingComponent.Effect.Position,
				*buildingComponent.Effect.TargetPosition,
				raftTravelDistancePerTick,
			)

			_, playerPosition, _ := QueryPlayerComponent[comp.Position](
				world,
				buildingComponent.Effect.Player,
				filter.Component[comp.Player](),
				filter.Component[comp.TileMap](),
				filter.Component[comp.Position](),
			)

			if buildingComponent.Effect.Position == playerPosition.Island {
				buildingComponent.Effect.TargetPosition = nil
				buildingComponent.Effect.SendingPosition = nil
				_ = unloadShip(world, buildingComponent.Effect)
			} else if buildingComponent.Effect.Position == *buildingComponent.Effect.TargetPosition {
				previousTargetPosition := buildingComponent.Effect.TargetPosition
				buildingComponent.Effect.TargetPosition = &playerPosition.Island
				buildingComponent.Effect.SendingPosition = previousTargetPosition
				_ = lootShipwreck(world, buildingComponent.Effect)
			}

			err := updateEffect(world, id, buildingComponent.Effect)
			if err != nil {
				return true
			}
			return true
		})
}

func findPointAtDistance(start, end [2]float64, distance float64) [2]float64 {
	x1, y1 := start[0], start[1]
	x2, y2 := end[0], end[1]

	dx := x2 - x1
	dy := y2 - y1

	length := math.Sqrt(dx*dx + dy*dy)

	if distance > length {
		return end
	}

	unitDx := dx / length
	unitDy := dy / length

	newX := x1 + unitDx*distance
	newY := y1 + unitDy*distance

	return [2]float64{newX, newY}
}

func lootShipwreck(world cardinal.WorldContext, effect *comp.Effect) error {
	playerPositionsIDs, playerPositions, _ := QueryAllComponents[comp.Position](
		world,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
		filter.Component[comp.Position](),
	)

	var shipwreckResources *comp.ShipwreckResources
	var playerMapEntityID types.EntityID
	for i, position := range playerPositions {
		if position.Shipwreck == effect.Position {
			playerMapEntityID = playerPositionsIDs[i]
			shipwreckResources, _ = cardinal.GetComponent[comp.ShipwreckResources](world, playerMapEntityID)
			break
		}
	}

	effect.LootResources = shipwreckResources.Resources
	shipwreckResources.Resources = nil
	shipwreckResources.LastSpawnTime = world.Timestamp()
	if err := cardinal.SetComponent(world, playerMapEntityID, shipwreckResources); err != nil {
		return err
	}

	return nil
}

func unloadShip(world cardinal.WorldContext, effect *comp.Effect) error {
	if effect.LootResources == nil {
		return nil
	}
	playerResourcesID, playerResources, _ := QueryPlayerComponent[comp.PlayerResources](
		world,
		effect.Player,
		filter.Component[comp.Player](),
		filter.Component[comp.TileMap](),
		filter.Component[comp.PlayerResources](),
	)

	var err error
	for _, lootResource := range *effect.LootResources {
		playerResource, _ := GetResourceByType(playerResources, lootResource.Type)
		playerResource.Amount += lootResource.Amount
		SetResourceByType(playerResources, *playerResource)
		err = cardinal.SetComponent(world, playerResourcesID, playerResources)
	}
	effect.LootResources = nil
	return err
}
