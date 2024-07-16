package component

type Position struct {
	Island    [2]float64 `json:"island"`
	Shipwreck [2]float64 `json:"shipwreck"`
}

func (Position) Name() string {
	return "Position"
}
