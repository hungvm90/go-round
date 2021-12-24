package round

import (
	"math"
)

type RoundMode int
type Rounder func(value float64, scale uint) float64

const (
	UP RoundMode = iota
	DOWN
	CEILING
	FLOOR
	HALF_UP
	HALF_DOWN
	HALF_EVEN
)

var (
	// Roundings defines the set of Rounders used by Context. Users may add their
	// own, but modification of this map is not safe during any other parallel
	// Context operations.
	Roundings = map[RoundMode]Rounder{
		UP:        roundUp,
		DOWN:      roundDown,
		CEILING:   roundCeiling,
		FLOOR:     roundFloor,
		HALF_UP:   roundHalfUp,
		HALF_DOWN: roundHalfDown,
		HALF_EVEN: roundHalfEven,
	}
)

func Round(value float64, scale uint, mode RoundMode) float64 {
	rounding, ok := Roundings[mode]
	if ok {
		return rounding(value, scale)
	}
	return roundHalfUp(value, scale)
}

func roundUp(value float64, scale uint) float64 {
	sign := 1.0
	if value < 0 {
		sign = -1
	}
	value = math.Abs(value)
	factor := math.Pow(10, float64(scale))
	return sign * math.Ceil(value*factor) / factor
}

func roundDown(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	return math.Trunc(value*factor) / factor
}

func roundCeiling(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	return math.Ceil(value*factor) / factor
}

func roundFloor(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	return math.Floor(value*factor) / factor
}

func roundHalfUp(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	return math.Round(value*factor) / factor
}

func roundHalfDown(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	temp := value * factor
	_, div := math.Modf(temp)
	if math.Abs(div) <= 0.5 {
		temp = math.Trunc(temp)
	}
	if value > 0 {
		return math.Round(temp) / factor
	} else {
		return math.Round(temp) / factor
	}

}

func roundHalfEven(value float64, scale uint) float64 {
	factor := math.Pow(10, float64(scale))
	return math.RoundToEven(value*factor) / factor
}
