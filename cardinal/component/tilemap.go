package component

type Tile struct {
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
		{Tile: CoastlineTile, Building: nil},
		{Tile: CoastlineTile, Building: nil},
		{Tile: CoastlineTile, Building: nil},
		{Tile: CoastlineTile, Building: nil},
		{Tile: CoastlineTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: &mainBuilding},
		{Tile: GenericTile, Building: nil},
		{Tile: StoneTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: GenericTile, Building: nil},
		{Tile: WoodTile, Building: nil},
	}

	return &tiles
}
