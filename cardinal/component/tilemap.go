package component

type Tile struct {
	ID       int       `json:"id"`
	Tile     TileType  `json:"tile"`
	Building *Building `json:"building"`
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

const (
	GenericTile   TileType = "Generic"
	WoodTile      TileType = "Wood"
	StoneTile     TileType = "Stone"
	WaterTile     TileType = "Water"
	CoastlineTile TileType = "Coastline"
)

func GetDefaultTiles() *[]Tile {
	var mainBuilding, _ = GetBuilding(Main)
	tiles := []Tile{
		{ID: Tile0, Tile: CoastlineTile, Building: nil},
		{ID: Tile1, Tile: CoastlineTile, Building: nil},
		{ID: Tile2, Tile: CoastlineTile, Building: nil},
		{ID: Tile3, Tile: CoastlineTile, Building: nil},
		{ID: Tile4, Tile: CoastlineTile, Building: nil},
		{ID: Tile5, Tile: GenericTile, Building: nil},
		{ID: Tile6, Tile: GenericTile, Building: nil},
		{ID: Tile7, Tile: GenericTile, Building: nil},
		{ID: Tile8, Tile: GenericTile, Building: nil},
		{ID: Tile9, Tile: GenericTile, Building: nil},
		{ID: Tile10, Tile: GenericTile, Building: nil},
		{ID: Tile11, Tile: GenericTile, Building: nil},
		{ID: Tile12, Tile: GenericTile, Building: nil},
		{ID: Tile13, Tile: GenericTile, Building: nil},
		{ID: Tile14, Tile: GenericTile, Building: nil},
		{ID: Tile15, Tile: GenericTile, Building: nil},
		{ID: Tile16, Tile: GenericTile, Building: nil},
		{ID: Tile17, Tile: GenericTile, Building: nil},
		{ID: Tile18, Tile: GenericTile, Building: nil},
		{ID: Tile19, Tile: GenericTile, Building: nil},
		{ID: Tile20, Tile: GenericTile, Building: nil},
		{ID: Tile21, Tile: GenericTile, Building: nil},
		{ID: Tile22, Tile: GenericTile, Building: &mainBuilding},
		{ID: Tile23, Tile: GenericTile, Building: nil},
		{ID: Tile24, Tile: StoneTile, Building: nil},
		{ID: Tile25, Tile: GenericTile, Building: nil},
		{ID: Tile26, Tile: GenericTile, Building: nil},
		{ID: Tile27, Tile: GenericTile, Building: nil},
		{ID: Tile28, Tile: GenericTile, Building: nil},
		{ID: Tile29, Tile: GenericTile, Building: nil},
		{ID: Tile30, Tile: GenericTile, Building: nil},
		{ID: Tile31, Tile: GenericTile, Building: nil},
		{ID: Tile32, Tile: GenericTile, Building: nil},
		{ID: Tile33, Tile: GenericTile, Building: nil},
		{ID: Tile34, Tile: GenericTile, Building: nil},
		{ID: Tile35, Tile: GenericTile, Building: nil},
		{ID: Tile36, Tile: GenericTile, Building: nil},
		{ID: Tile37, Tile: GenericTile, Building: nil},
		{ID: Tile38, Tile: GenericTile, Building: nil},
		{ID: Tile39, Tile: GenericTile, Building: nil},
		{ID: Tile40, Tile: GenericTile, Building: nil},
		{ID: Tile41, Tile: GenericTile, Building: nil},
		{ID: Tile42, Tile: GenericTile, Building: nil},
		{ID: Tile43, Tile: GenericTile, Building: nil},
		{ID: Tile44, Tile: WoodTile, Building: nil},
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
