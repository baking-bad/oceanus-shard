package component

import (
	"fmt"
	"math/rand/v2"
	"oceanus-shard/constants"
)

type ShipwreckResources struct {
	Resources     *[]Resource `json:"resources"`
	LastSpawnTime uint64      `json:"lastSpawnTime"`
}

func GetShipwreckResourceTypes() []ResourceType {
	return []ResourceType{
		Wood,
		Stone,
		Fish,
	}
}

func randomIntInRange(min, max int) float64 {
	if min >= max {
		panic("min should be lower than max")
	}
	return float64(min) + rand.Float64()*float64(max-min)
}

func GetShipwreckDefaultResourceAmount(resourceType ResourceType) float64 {
	switch resourceType {
	case Wood:
		return randomIntInRange(constants.ShipwreckMinWoodAmount, constants.ShipwreckMaxWoodAmount)
	case Stone:
		return randomIntInRange(constants.ShipwreckMinStoneAmount, constants.ShipwreckMaxStoneAmount)
	case Fish:
		return randomIntInRange(constants.ShipwreckMinFishAmount, constants.ShipwreckMaxFishAmount)
	case Cotton:
	case Iron:
	default:
		panic(fmt.Sprintf("unexpected component.ResourceType: %#v", resourceType))
	}
	return 0
}

func (ShipwreckResources) Name() string {
	return "ShipwreckResources"
}
