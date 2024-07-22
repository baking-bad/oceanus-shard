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
