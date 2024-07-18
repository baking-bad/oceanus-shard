package main

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"

	"oceanus-shard/component"
	"oceanus-shard/constants"
	"oceanus-shard/msg"
	"oceanus-shard/query"
	"oceanus-shard/system"
)

func main() {
	w, err := cardinal.NewWorld(
		cardinal.WithDisableSignatureVerification(),
		cardinal.WithTickChannel(time.Tick(constants.TickRate)),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	MustInitWorld(w)

	Must(w.StartGame())
}

// MustInitWorld registers all components, messages, queries, and systems. This initialization happens in a helper
// function so that this can be used directly in tests.
func MustInitWorld(w *cardinal.World) {
	// Register components
	// NOTE: You must register your components here for it to be accessible.
	Must(
		cardinal.RegisterComponent[component.Player](w),
		cardinal.RegisterComponent[component.Health](w),
		cardinal.RegisterComponent[component.Farming](w),
		cardinal.RegisterComponent[component.Building](w),
		cardinal.RegisterComponent[component.PlayerResources](w),
		cardinal.RegisterComponent[component.ShipwreckResources](w),
		cardinal.RegisterComponent[component.TileMap](w),
		cardinal.RegisterComponent[component.Position](w),
		cardinal.RegisterComponent[component.Ship](w),
	)

	// Register messages (user action)
	// NOTE: You must register your transactions here for it to be executed.
	Must(
		cardinal.RegisterMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](w, "create-player"),
		cardinal.RegisterMessage[msg.CreateBuildingMsg, msg.CreateBuildingResult](w, "create-building"),
		cardinal.RegisterMessage[msg.DeleteBuildingMsg, msg.DeleteBuildingResult](w, "delete-building"),
		cardinal.RegisterMessage[msg.CreateShipMsg, msg.CreateShipResult](w, "create-ship"),
		cardinal.RegisterMessage[msg.RelocateBuildingMsg, msg.RelocateBuildingResult](w, "relocate-building"),
	)

	// Register queries
	// NOTE: You must register your queries here for it to be accessible.
	Must(
		cardinal.RegisterQuery[
			query.MapStateRequest,
			query.MapStateResponse,
		](w, "player-map", query.PlayerMap),
		cardinal.RegisterQuery[
			query.BuildingsInfoRequest,
			[]query.BuildingInfoResponse,
		](w, "buildings-info", query.AllBuildings),
		cardinal.RegisterQuery[
			query.PlayerResourcesRequest,
			query.PlayerResourcesResponse,
		](w, "player-resources", query.PlayerResources),
		cardinal.RegisterQuery[
			query.GlobalMapRequest,
			[]query.GlobalMapResponse,
		](w, "global-map", query.GlobalMap),
	)

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	Must(cardinal.RegisterSystems(w,
		system.FarmingSystem,
		system.PlayerSpawnerSystem,
		system.CreateBuildingSystem,
		system.DeleteBuildingSystem,
		system.CreateShipSystem,
		system.ShipsSpawnerSystem,
		system.RelocateBuildingSystem,
	))
}

func Must(err ...error) {
	e := errors.Join(err...)
	if e != nil {
		log.Fatal().Err(e).Msg("")
	}
}
