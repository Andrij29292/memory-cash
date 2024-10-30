package memorycache

type keys string

func (k keys) In(memory map[keys]any) bool {
	isKeyInKeys := false
	for m_k := range memory {
		if k == m_k {
			isKeyInKeys = true
		}
	}

	return isKeyInKeys
}
