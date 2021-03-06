package valuetypes

const (
	UserStateActive UserState = iota + 1
	UserStateBlocked
)

type UserState int

func (u UserState) String() string {
	switch u {
	case UserStateActive:
		return "active user"
	case UserStateBlocked:
		return "blocked user"
	}

	return ""
}

func (u UserState) IsValid() bool {
	if u != UserStateActive && u != UserStateBlocked {
		return false
	}

	return true
}
