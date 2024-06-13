package system

import (
	"fmt"
	comp "oceanus-shard/component"
	"pkg.world.dev/world-engine/cardinal/persona/component"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"
)

// queryTargetPlayer queries for the target player's entity ID and health component.
func queryTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Health, error) {
	var playerID types.EntityID
	var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health]())).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if playerHealth == nil {
		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerHealth, err
}

func querySignerComponentByPersona(world cardinal.WorldContext, targetPersonaName string) (types.EntityID, *component.SignerComponent, error) {
	var signerEntityID types.EntityID
	var signerEntity *component.SignerComponent
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[component.SignerComponent]())).Each(world,
		func(id types.EntityID) bool {
			signerEntity, err = cardinal.GetComponent[component.SignerComponent](world, id)

			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if signerEntity.PersonaTag == targetPersonaName {
				signerEntityID = id
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
	if signerEntityID == 0 {
		return 0, nil, fmt.Errorf("signer component %q does not exist", targetPersonaName)
	}

	return signerEntityID, signerEntity, err
}
