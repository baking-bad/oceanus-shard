package component

type TileType string
type BuildingType string

const (
	Generic TileType = "Generic"
	Wood             = "Wood"
	Water            = "Water"
	Stone            = "Stone"
)

const (
	None           BuildingType = "None"
	Main                        = "Main"
	Woodcutter                  = "Woodcutter"
	Quarry                      = "Quarry"
	FishermanHut                = "FishermanHut"
	Shipyard                    = "Shipyard"
	Warehouse                   = "Warehouse"
	UnitLimitHouse              = "UnitLimitHouse"
)

type Tile struct {
	Tile     TileType     `json:"tile"`
	Building BuildingType `json:"building"`
	X        int          `json:"x"`
	Y        int          `json:"y"`
}

func GetDefaultTiles() *[]Tile {
	tiles := make([]Tile, 8)

	tiles[0] = Tile{Tile: Generic, Building: Warehouse, X: 0, Y: 0}
	tiles[1] = Tile{Tile: Stone, Building: Quarry, X: 1, Y: 0}
	tiles[2] = Tile{Tile: Wood, Building: FishermanHut, X: 2, Y: 0}
	tiles[3] = Tile{Tile: Water, Building: UnitLimitHouse, X: 3, Y: 0}
	tiles[4] = Tile{Tile: Generic, Building: Shipyard, X: 4, Y: 0}
	tiles[5] = Tile{Tile: Generic, Building: None, X: 5, Y: 0}
	tiles[6] = Tile{Tile: Generic, Building: Woodcutter, X: 6, Y: 0}
	tiles[7] = Tile{Tile: Generic, Building: Main, X: 7, Y: 0}

	return &tiles
}

type TileMap struct {
	Tiles *[]Tile `json:"tiles"`
}

func (TileMap) Name() string {
	return "TileMap"
}
