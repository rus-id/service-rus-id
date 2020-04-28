package valuetypes

const (
	ActiveUser UserStatus = iota + 1
	BlockedUser
)

type UserStatus int

func (us UserStatus) String() string {
	switch us {
	case ActiveUser:
		return "active user"
	case BlockedUser:
		return "blocked user"
	}

	return ""
}
