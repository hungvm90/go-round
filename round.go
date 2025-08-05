package round

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
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
	result := round(value, scale, math.Ceil)
	if sign > 0 {
		return result
	} else {
		return -result
	}
}

func round(value float64, scale uint, fn func(x float64) float64) float64 {
	factor := new(big.Float).SetPrec(64).SetFloat64(math.Pow10(int(scale)))
	f := new(big.Float).SetPrec(64).SetFloat64(value)
	f.Mul(f, factor)
	t, _ := strconv.ParseFloat(fmt.Sprintf("%f", f), 64)
	t = fn(t)
	f.SetFloat64(t)
	f.Quo(f, factor)
	result, _ := strconv.ParseFloat(fmt.Sprintf("%f", f), 64)
	return result
}

func roundDown(value float64, scale uint) float64 {
	return round(value, scale, math.Trunc)
}

func roundCeiling(value float64, scale uint) float64 {
	return round(value, scale, math.Ceil)
}

func roundFloor(value float64, scale uint) float64 {
	return round(value, scale, math.Floor)
}

func roundHalfUp(value float64, scale uint) float64 {
	return round(value, scale, math.Round)
}

func roundHalfDown(value float64, scale uint) float64 {
	factor := new(big.Float).SetPrec(64).SetFloat64(math.Pow10(int(scale)))
	f := new(big.Float).SetPrec(64).SetFloat64(value)
	f.Mul(f, factor)
	t, _ := strconv.ParseFloat(fmt.Sprintf("%f", f), 64)
	_, div := math.Modf(t)
	if math.Abs(div) <= 0.5 {
		t = math.Trunc(t)
	} else {
		t = math.Round(t)
	}
	f.SetFloat64(t)
	f.Quo(f, factor)
	result, _ := strconv.ParseFloat(fmt.Sprintf("%f", f), 64)
	return result

}

func roundHalfEven(value float64, scale uint) float64 {
	return round(value, scale, math.RoundToEven)
}
