package types

type UnitAndValueInt64 struct {
	Unit  string `json:"unit"`
	Value int64  `json:"value"`
}

type UnitAndValueFloat64 struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}
