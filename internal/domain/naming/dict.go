package naming

// those dictionaries are static, however it would be better to load them with program from config files/storage

import "time"

const ( // daytime
	earlyMorning = "early morning"
	morning      = "morning"

	midday = "daytime"
	noon   = "noon"

	evening     = "evening"
	lateEvening = "late evening"

	night     = "night"
	lateNight = "nighttime"
	midnight  = "midnight"
)

const ( // season
	spring = "spring"
	summer = "summer"
	autumn = "autumn"
	winter = "winter"
)

var ( // months
	seasonMap = map[time.Month][]string{
		time.January:   {"January", winter},
		time.February:  {"February", winter},
		time.March:     {"March", spring},
		time.April:     {"April", spring},
		time.May:       {"May", spring},
		time.June:      {"June", summer},
		time.July:      {"July", summer},
		time.August:    {"August", summer},
		time.September: {"September", autumn},
		time.October:   {"October", autumn},
		time.November:  {"November", autumn},
		time.December:  {"December", winter},
	}
)

const (
	thisYear = "this year's"
	lastYear = "last year's"
)

var adjectives = []string{
	"majestic",
	"unforgettable",
	"wonderful",
	"charming",
	"marvelous",
	"magnificent",
	"fascinating",
	"captivating",
	"charming",
	"delightful",
	"enchanting",
	"beautiful",
	"captivating",
	"enjoyable",
	"entertaining",
	"lovely",
	"pleasing",
}

var nouns = []string{
	"occasion at ",
	"event in ",
	"journey to ",
	"trip to ",
	"tour to ",
	"experience from ",
}
