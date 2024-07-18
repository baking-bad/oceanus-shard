package system

import (
	"time"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "oceanus-shard/component"
)

// ShipsSpawnerSystem -
func ShipsSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Building]())).
		Each(
			world,
			func(id types.EntityID) bool {
				buildingComponent, _ := cardinal.GetComponent[comp.Building](world, id)

				if buildingComponent.BuildingShip == nil {
					return true
				}

				playerComponent, _ := cardinal.GetComponent[comp.Player](world, id)
				playerMapEntityID, playerMap, _ := QueryPlayerComponent[comp.TileMap](
					world,
					playerComponent.Nickname,
					filter.Component[comp.Player](),
					filter.Component[comp.TileMap](),
				)
				playerMapPosition, _ := cardinal.GetComponent[comp.Position](world, playerMapEntityID)

				endBuildingTime := buildingComponent.BuildingShip.BuildingTimeStartedAt +
					uint64(time.Duration(buildingComponent.BuildingShip.BuildingTimeSeconds)*time.Second/time.Millisecond)

				if world.Timestamp() < endBuildingTime {
					return true
				}

				_, err := cardinal.Create(world,
					playerComponent,
					comp.Ship{
						Type:      buildingComponent.BuildingShip.Type,
						PositionX: playerMapPosition.Island[0],
						PositionY: playerMapPosition.Island[1],
					},
				)

				if err != nil {
					return true
				}

				buildingComponent.BuildingShip = nil
				tile := &(*playerMap.Tiles)[buildingComponent.TileID]
				tile.Building = buildingComponent

				if err := cardinal.SetComponent(world, id, buildingComponent); err != nil {
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
