package naming

import (
	"math/rand"
	"strings"
	"time"

	"github.com/krylphi/travels-naming/internal/util"
)

const (
	yearToken int8 = iota
	adjToken
	seasonToken
	dayTimeToken
	nounToken
)

const minTokens = 3

type randomizer struct {
	random  *rand.Rand
	tokens  map[int8]bool
	flipper *func() bool
}

func newRandomizer(t time.Time) *randomizer {
	tokens := make(map[int8]bool, 5)
	return &randomizer{
		random: rand.New(rand.NewSource(t.UnixNano())),
		tokens: tokens,
	}
}

func epithetFromTime(t time.Time, location string) string {
	now := time.Now()
	random := newRandomizer(now)
	return epithetFunction(random, now, t, location)
}

func epithetFunction(random *randomizer, now time.Time, t time.Time, location string) string {
	year, month, _ := t.Date()
	sb := strings.Builder{}
	if random.flipCoin() {
		if year == now.Year() {
			sb.WriteString(thisYear)
			sb.WriteString(" ")
			random.addToken(yearToken)
		}
		if year+1 == now.Year() {
			sb.WriteString(lastYear)
			sb.WriteString(" ")
			random.addToken(yearToken)
		}
	}

	sb.WriteString(random.getRandomAdjective(false))

	sb.WriteString(random.getRandomSeason(month, false))

	sb.WriteString(random.getRandomDayTime(t, random.checkToken(seasonToken)))
	if !random.checkToken(dayTimeToken) {
		sb.WriteString(random.getRandomWeekDay(t, false))
	}

	sb.WriteString(random.getRandomNoun())

	sb.WriteString(location)

	if random.tokenCount() < minTokens {
		if !random.checkToken(dayTimeToken) {
			sb.WriteString(" at ")
			sb.WriteString(random.getRandomDayTime(t, true))
		}
		if !random.checkToken(seasonToken) {
			sb.WriteString(" in ")
			sb.WriteString(random.getRandomSeason(month, true))
		}

	}
	return util.Title(sb)
}

func (r randomizer) flipCoin() bool {
	if r.flipper != nil {
		flip := *r.flipper
		return flip()
	}
	return r.random.Intn(2) == 1
}

func (r randomizer) getRandomAdjective(noFlip bool) string {
	if noFlip || r.flipCoin() {
		r.addToken(adjToken)
		return util.Concat(adjectives[r.random.Intn(len(adjectives))], " ")
	}
	return ""
}

func (r randomizer) getRandomSeason(month time.Month, noFlip bool) string {
	if noFlip || r.flipCoin() {
		season := seasonMap[month]
		idx := r.random.Intn(len(season))
		r.addToken(seasonToken)
		return util.Concat(season[idx], " ")
	}
	return ""
}

func (r randomizer) getRandomNoun() string {
	idx := r.random.Intn(len(nouns))
	r.addToken(nounToken)
	return nouns[idx]
}

func (r randomizer) getRandomDayTime(t time.Time, noFlip bool) string {
	sb := strings.Builder{}

	if noFlip || r.flipCoin() {
		hours, minutes, _ := t.Clock()
		switch {
		case (hours == 23 && minutes >= 30) || (hours == 0 && minutes <= 30):
			sb.WriteString(midnight)
		case (hours == 11 && minutes >= 30) || (hours == 12 && minutes <= 30):
			sb.WriteString(noon)
		case hours >= 0 && hours <= 2:
			sb.WriteString(night)
		case hours > 2 && hours <= 4:
			sb.WriteString(lateNight)
		case hours > 4 && hours <= 7:
			sb.WriteString(earlyMorning)
		case hours > 7 && hours < 12:
			sb.WriteString(morning)
		case hours >= 12 && hours <= 17:
			sb.WriteString(midday)
		case hours > 17 && hours <= 20:
			sb.WriteString(evening)
		case hours > 20 && hours <= 23:
			sb.WriteString(lateEvening)
		}
		if !noFlip || r.checkToken(seasonToken) {
			sb.WriteString(" ")
		}
		r.addToken(dayTimeToken)
	}
	return sb.String()
}

func (r randomizer) getRandomWeekDay(t time.Time, noFlip bool) string {
	sb := strings.Builder{}

	if noFlip || r.flipCoin() {
		switch wd := t.Weekday(); wd {
		case time.Saturday | time.Sunday:
			if r.flipCoin() {
				sb.WriteString("weekend")
			} else {
				sb.WriteString(wd.String())
			}
		default:
			sb.WriteString(wd.String())
		}
		r.addToken(dayTimeToken)
		sb.WriteString(" ")
	}
	return sb.String()
}

func (r randomizer) checkToken(token int8) bool {
	return r.tokens[token]
}

func (r randomizer) tokenCount() int {
	return len(r.tokens)
}

func (r randomizer) addToken(token int8) {
	r.tokens[token] = true
}
