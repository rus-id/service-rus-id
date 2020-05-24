package snapshots

type ToleranceSnapshot struct {
	FromID    string
	ToID      string
	Accessors []int64
}

func NewTolerance(
	fromID string,
	toID string,
	accessors []int64,
) ToleranceSnapshot {
	return ToleranceSnapshot{
		FromID:    fromID,
		ToID:      toID,
		Accessors: accessors,
	}
}
