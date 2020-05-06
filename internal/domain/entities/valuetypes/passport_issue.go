package valuetypes

import "time"

type PassportIssue struct {
	organisations string
	date          time.Time
	code          string
}
