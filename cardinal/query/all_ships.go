package query

import (
	comp "oceanus-shard/component"
	"oceanus-shard/system"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
)

type AllShipsRequest struct {
	Nickname string
}

// AllShips -
func AllShips(world cardinal.WorldContext, req *AllShipsRequest) (*map[string]comp.Effect, error) {
	_, effects, err := system.QueryAllComponents[comp.Effect](
		world,
		filter.Component[comp.Player](),
		filter.Component[comp.Building](),
		filter.Component[comp.Effect](),
	)

	shipsMap := make(map[string]comp.Effect)
	for _, effect := range effects {
		if effect.Amount > 0 {
			shipsMap[effect.ID] = *effect
		}
	}

	return &shipsMap, err
}
