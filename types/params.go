package types

type MaxResults int

func (m *MaxResults) Valid() bool {
	if m == nil {
		return false
	}

	return *m > 0
}

type HistoryParams struct {
	MaxResults *MaxResults
}
