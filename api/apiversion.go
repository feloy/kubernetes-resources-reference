package api

import "regexp"

type ApiVersion string

func (a ApiVersion) String() string {
	return string(a)
}

func (this ApiVersion) LessThan(that ApiVersion) bool {
	re := regexp.MustCompile("(v\\d+)(alpha|beta|)(\\d*)")
	thisMatches := re.FindStringSubmatch(string(this))
	thatMatches := re.FindStringSubmatch(string(that))

	a := []string{thisMatches[1]}
	if len(thisMatches) > 2 {
		a = []string{thisMatches[2], thisMatches[1], thisMatches[0]}
	}

	b := []string{thatMatches[1]}
	if len(thatMatches) > 2 {
		b = []string{thatMatches[2], thatMatches[1], thatMatches[0]}
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		v1 := ""
		v2 := ""
		if i < len(a) {
			v1 = a[i]
		}
		if i < len(b) {
			v2 = b[i]
		}
		// If the "beta" or "alpha" is missing, then it is ga (empty string comes before non-empty string)
		if len(v1) == 0 || len(v2) == 0 {
			return v1 < v2
		}
		// The string with the higher number comes first (or in the case of alpha/beta, beta comes first)
		if v1 != v2 {
			return v1 > v2
		}
	}

	return false
}
