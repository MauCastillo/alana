package intervals

// Interval represents interval enum.
type Interval struct {
	Value string
}

var (
	// Seconds one Seconds
	Seconds = &Interval{Value:"1s"}
	// Minute One Minute
	Minute = &Interval{Value:"1m"}
	// ThreeMinutes Three Minutes
	ThreeMinutes = &Interval{Value:"3m"}
	// FiveMinutes Five Minutes
	FiveMinutes = &Interval{Value:"5m"}
	// FifteenMinutes Fifteen Minutes
	FifteenMinutes = &Interval{Value:"15m"}
	// ThirtyMinutes Thirty Minutes
	ThirtyMinutes = &Interval{Value:"30m"}
	// Hour One Hour
	Hour = &Interval{Value:"1h"}
	// TwoHours Two Hours
	TwoHours = &Interval{Value:"2h"}
	// FourHours Four Hours
	FourHours = &Interval{Value:"4h"}
	// SixHours Six Hours
	SixHours = &Interval{Value:"6h"}
	// EightHours Eight Hours
	EightHours = &Interval{Value:"8h"}
	// TwelveHours Twelve Hours
	TwelveHours = &Interval{Value:"12h"}
	// Day One Day
	Day = &Interval{Value:"1d"}
	// ThreeDays Three Days
	ThreeDays = &Interval{Value:"3d"}
	// Week One Week
	Week = &Interval{Value:"1w"}
	// Month One Month
	Month = &Interval{Value:"1M"}
)
