package component

type Tile struct {
	ID       int          `json:"id"`
	Tile     TileType     `json:"tile"`
	Building BuildingType `json:"building"`
}

type TileMap struct {
	Tiles  *[]Tile `json:"tiles"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
}

func (TileMap) Name() string {
	return "TileMap"
}

const MapWidth = 9
const MapHeight = 5

type TileType string
type BuildingType string

const (
	Generic   TileType = "Generic"
	Wood      TileType = "Wood"
	Stone     TileType = "Stone"
	Water     TileType = "Water"
	Coastline TileType = "Coastline"
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

func GetDefaultTiles() *[]Tile {
	tiles := []Tile{
		{ID: Tile0, Tile: Coastline, Building: None},
		{ID: Tile1, Tile: Coastline, Building: None},
		{ID: Tile2, Tile: Coastline, Building: None},
		{ID: Tile3, Tile: Coastline, Building: None},
		{ID: Tile4, Tile: Coastline, Building: None},
		{ID: Tile5, Tile: Generic, Building: None},
		{ID: Tile6, Tile: Generic, Building: None},
		{ID: Tile7, Tile: Generic, Building: None},
		{ID: Tile8, Tile: Generic, Building: None},
		{ID: Tile9, Tile: Generic, Building: None},
		{ID: Tile10, Tile: Generic, Building: None},
		{ID: Tile11, Tile: Generic, Building: None},
		{ID: Tile12, Tile: Generic, Building: None},
		{ID: Tile13, Tile: Generic, Building: None},
		{ID: Tile14, Tile: Generic, Building: None},
		{ID: Tile15, Tile: Generic, Building: None},
		{ID: Tile16, Tile: Generic, Building: None},
		{ID: Tile17, Tile: Generic, Building: None},
		{ID: Tile18, Tile: Generic, Building: None},
		{ID: Tile19, Tile: Generic, Building: None},
		{ID: Tile20, Tile: Generic, Building: None},
		{ID: Tile21, Tile: Generic, Building: None},
		{ID: Tile22, Tile: Generic, Building: Main},
		{ID: Tile23, Tile: Generic, Building: None},
		{ID: Tile24, Tile: Stone, Building: None},
		{ID: Tile25, Tile: Generic, Building: None},
		{ID: Tile26, Tile: Generic, Building: None},
		{ID: Tile27, Tile: Generic, Building: None},
		{ID: Tile28, Tile: Generic, Building: None},
		{ID: Tile29, Tile: Generic, Building: None},
		{ID: Tile30, Tile: Generic, Building: None},
		{ID: Tile31, Tile: Generic, Building: None},
		{ID: Tile32, Tile: Generic, Building: None},
		{ID: Tile33, Tile: Generic, Building: None},
		{ID: Tile34, Tile: Generic, Building: None},
		{ID: Tile35, Tile: Generic, Building: None},
		{ID: Tile36, Tile: Generic, Building: None},
		{ID: Tile37, Tile: Generic, Building: None},
		{ID: Tile38, Tile: Generic, Building: None},
		{ID: Tile39, Tile: Generic, Building: None},
		{ID: Tile40, Tile: Generic, Building: None},
		{ID: Tile41, Tile: Generic, Building: None},
		{ID: Tile42, Tile: Generic, Building: None},
		{ID: Tile43, Tile: Generic, Building: None},
		{ID: Tile44, Tile: Wood, Building: None},
	}

	return &tiles
}

const (
	Tile0  = 0
	Tile1  = 1
	Tile2  = 2
	Tile3  = 3
	Tile4  = 4
	Tile5  = 5
	Tile6  = 6
	Tile7  = 7
	Tile8  = 8
	Tile9  = 9
	Tile10 = 10
	Tile11 = 11
	Tile12 = 12
	Tile13 = 13
	Tile14 = 14
	Tile15 = 15
	Tile16 = 16
	Tile17 = 17
	Tile18 = 18
	Tile19 = 19
	Tile20 = 20
	Tile21 = 21
	Tile22 = 22
	Tile23 = 23
	Tile24 = 24
	Tile25 = 25
	Tile26 = 26
	Tile27 = 27
	Tile28 = 28
	Tile29 = 29
	Tile30 = 30
	Tile31 = 31
	Tile32 = 32
	Tile33 = 33
	Tile34 = 34
	Tile35 = 35
	Tile36 = 36
	Tile37 = 37
	Tile38 = 38
	Tile39 = 39
	Tile40 = 40
	Tile41 = 41
	Tile42 = 42
	Tile43 = 43
	Tile44 = 44
)
