package mytumblrhandlers

type RateLimits struct {
	Calls            PerLimits `json:"calls"`
	Posts            PerLimits `json:"posts"`
	Images           PerLimits `json:"images"`
	Follows          PerLimits `json:"follows"`
	Likes            PerLimits `json:"likes"`
	Blogs            PerLimits `json:"blogs"`
	Videos           PerLimits `json:"videos"`
	VideoTimeMinutes PerLimits `json:"videotimeminutes"`
}

type PerLimits struct {
	PerIP          *TimeLimits `json:"IP,omitempty"`
	PerConsumerKey *TimeLimits `json:"consumerkey,omitempty"`
	PerUser        *TimeLimits `json:"user,omitempty"`
}

type TimeLimits struct {
	PerMinute *int `json:"minute,omitempty"`
	PerHour   *int `json:"hour,omitempty"`
	PerDay    *int `json:"day,omitempty"`
}
