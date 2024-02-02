package round

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

type TestCase struct {
	Input    string `json:"input"`
	Up       string `json:"up"`
	Down     string `json:"down"`
	Ceiling  string `json:"ceiling"`
	Floor    string `json:"floor"`
	HalfUp   string `json:"halfUp"`
	HalfDown string `json:"halfDown"`
	HalfEven string `json:"halfEven"`
}

func TestRound2(t *testing.T) {
	require.Equal(t, -0.14, Round(-0.14, 2, UP))
	require.Equal(t, 0.1425, Round(0.1425, 4, DOWN))
}

func TestRound(t *testing.T) {
	dat, err := os.ReadFile("testcase.json")
	require.NoError(t, err)
	testcases := make([]TestCase, 0)
	err = json.Unmarshal(dat, &testcases)
	require.NoError(t, err)
	require.Equal(t, -0.9, Round(-0.95, 1, HALF_DOWN))
	for _, tc := range testcases {
		input, _ := strconv.ParseFloat(tc.Input, 64)
		up, _ := strconv.ParseFloat(tc.Up, 64)
		down, _ := strconv.ParseFloat(tc.Down, 64)
		ceiling, _ := strconv.ParseFloat(tc.Ceiling, 64)
		floor, _ := strconv.ParseFloat(tc.Floor, 64)
		halfUp, _ := strconv.ParseFloat(tc.HalfUp, 64)
		halfDown, _ := strconv.ParseFloat(tc.HalfDown, 64)
		halfEven, _ := strconv.ParseFloat(tc.HalfEven, 64)
		require.NoError(t, err, "fail to parse "+tc.Input)
		require.Equal(t, up, Round(input, 1, UP), tc.Input)
		require.Equal(t, down, Round(input, 1, DOWN), tc.Input)
		require.Equal(t, ceiling, Round(input, 1, CEILING), tc.Input)
		require.Equal(t, floor, Round(input, 1, FLOOR), tc.Input)
		require.Equal(t, halfUp, Round(input, 1, HALF_UP), tc.Input)
		require.Equal(t, halfDown, Round(input, 1, HALF_DOWN), tc.Input)
		require.Equal(t, halfEven, Round(input, 1, HALF_EVEN), tc.Input)
	}
}
