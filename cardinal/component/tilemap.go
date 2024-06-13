package component

type TileType string
type BuildingType string

const MapWidth = 3
const MapHeight = 3

const (
	Generic TileType = "Generic"
	Wood    TileType = "Wood"
	Water   TileType = "Water"
	Stone   TileType = "Stone"
)

const (
	None           BuildingType = "None"
	Main           BuildingType = "Main"
	Woodcutter     BuildingType = "Woodcutter"
	Quarry         BuildingType = "Quarry"
	FishermanHut   BuildingType = "FishermanHut"
	Shipyard       BuildingType = "Shipyard"
	Warehouse      BuildingType = "Warehouse"
	UnitLimitHouse BuildingType = "UnitLimitHouse"
)

type Tile struct {
	Tile     TileType     `json:"tile"`
	Building BuildingType `json:"building"`
}

func GetDefaultTiles() *map[int]Tile {
	tiles := map[int]Tile{
		0: {Tile: Generic, Building: Warehouse},
		1: {Tile: Stone, Building: Quarry},
		2: {Tile: Wood, Building: FishermanHut},
		3: {Tile: Water, Building: UnitLimitHouse},
		4: {Tile: Generic, Building: Shipyard},
		5: {Tile: Generic, Building: None},
		6: {Tile: Water, Building: UnitLimitHouse},
		7: {Tile: Generic, Building: Shipyard},
		8: {Tile: Generic, Building: None},
	}

	return &tiles
}

type TileMap struct {
	Tiles  *map[int]Tile `json:"tiles"`
	Width  int           `json:"width"`
	Height int           `json:"height"`
}

func (TileMap) Name() string {
	return "TileMap"
}
