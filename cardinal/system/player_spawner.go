package system

import (
	"fmt"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"pkg.world.dev/world-engine/cardinal"

	comp "oceanus-shard/component"
	"oceanus-shard/msg"
)

const (
	InitialHP = 100
)

// PlayerSpawnerSystem spawns players based on `CreatePlayer` transactions.
// This provides an example of a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](
		world,
		func(create cardinal.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerResult, error) {
			var playerExist = false
			var err error
			err = cardinal.NewSearch().Entity(
				filter.Contains(filter.Component[comp.Player]())).
				Each(world, func(id types.EntityID) bool {
					var player *comp.Player
					player, err = cardinal.GetComponent[comp.Player](world, id)
					if err != nil {
						return false
					}

					// Terminates the search if the player is found
					if player.Nickname == create.Msg.Nickname {
						playerExist = true
						return false
					}

					// Continue searching if the player is not the target player
					return true
				})
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
			}

			if playerExist {
				return msg.CreatePlayerResult{Success: false},
					fmt.Errorf("error creating player, player with nickname %s already exist", create.Msg.Nickname)
			}

			id, err := cardinal.Create(world,
				comp.Player{Nickname: create.Msg.Nickname},
				comp.TileMap{Tiles: comp.GetDefaultTiles()},
			)
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_player",
				"id":    id,
			})
			if err != nil {
				return msg.CreatePlayerResult{}, err
			}
			return msg.CreatePlayerResult{Success: true}, nil
		})
}
