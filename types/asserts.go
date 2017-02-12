package types

func AssertEqual(typ Type, types ...Type) bool {
	for _, t := range types {
		if !typ.Equals(t) {
			return false
		}
	}
	return true
}

func AssertMatch(typ Type, types ...Type) bool {
	for _, t := range types {
		if !typ.Match(t) {
			return false
		}
	}
	return true
}

func AssertSliceEqual(atypes []Type, btypes []Type) bool {
	if len(atypes) != len(btypes) {
		return false
	}

	for i, t := range atypes {
		if !t.Equals(btypes[i]) {
			return false
		}
	}
	return true
}

func AssertSliceMatch(atypes []Type, btypes []Type) bool {
	if len(atypes) != len(btypes) {
		return false
	}

	for i, t := range atypes {
		if !t.Match(btypes[i]) {
			return false
		}
	}
	return true
}
