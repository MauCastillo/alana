package intervals

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinute(t *testing.T) {
	c := require.New(t)
	c.Equal("1m", Minute.Value)
}

func TestThreeMinutes(t *testing.T) {
	c := require.New(t)
	c.Equal("3m", ThreeMinutes.Value)
}

func TestFiveMinutes(t *testing.T) {
	c := require.New(t)
	c.Equal("5m", FiveMinutes.Value)
}

func TestFifteenMinutes(t *testing.T) {
	c := require.New(t)
	c.Equal("15m", FifteenMinutes.Value)
}

func TestThirtyMinutes(t *testing.T) {
	c := require.New(t)
	c.Equal("30m", ThirtyMinutes.Value)
}

func TestHour(t *testing.T) {
	c := require.New(t)
	c.Equal("1h", Hour.Value)
}

func TestTwoHours(t *testing.T) {
	c := require.New(t)
	c.Equal("2h", TwoHours.Value)
}

func TestFourHours(t *testing.T) {
	c := require.New(t)
	c.Equal("4h", FourHours.Value)
}

func TestSixHours(t *testing.T) {
	c := require.New(t)
	c.Equal("6h", SixHours.Value)
}

func TestEightHours(t *testing.T) {
	c := require.New(t)
	c.Equal("8h", EightHours.Value)
}

func TestTwelveHours(t *testing.T) {
	c := require.New(t)
	c.Equal("12h", TwelveHours.Value)
}

func TestDay(t *testing.T) {
	c := require.New(t)
	c.Equal("1d", Day.Value)
}

func TestThreeDays(t *testing.T) {
	c := require.New(t)
	c.Equal("3d", ThreeDays.Value)
}

func TestWeek(t *testing.T) {
	c := require.New(t)
	c.Equal("1w", Week.Value)
}

func TestMonth(t *testing.T) {
	c := require.New(t)
	c.Equal("1M", Month.Value)
}

func TestSeconds(t *testing.T) {
	c := require.New(t)
	c.Equal("1s", Seconds.Value)
}