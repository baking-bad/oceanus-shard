package system

import (
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "oceanus-shard/component"
	"oceanus-shard/constants"
)

// ShipwreckResourcesSpawner -
func ShipwreckResourcesSpawner(world cardinal.WorldContext) error {
	return cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.ShipwreckResources]())).
		Each(world, func(id types.EntityID) bool {
			shipwreckResources, _ := cardinal.GetComponent[comp.ShipwreckResources](world, id)
			if (shipwreckResources.LastSpawnTime +
				uint64(constants.ShipwreckResourcesRespawnInterval.Milliseconds())) > world.Timestamp() {
				return true
			}
			if shipwreckResources.Resources != nil {
				return true
			}

			shipWreckResources := make([]comp.Resource, len(comp.GetShipwreckResourceTypes()))
			for i, resourceType := range comp.GetShipwreckResourceTypes() {
				shipWreckResources[i] = comp.Resource{
					Type:   resourceType,
					Amount: comp.GetShipwreckDefaultResourceAmount(resourceType),
				}
			}
			shipwreckResources.Resources = &shipWreckResources
			_ = cardinal.SetComponent(world, id, shipwreckResources)
			return true
		})
}
