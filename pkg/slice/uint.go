package slice

func DoseExist(list []uint, value uint) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func MapUnit64ToUint(l []uint64) []uint {
	r := make([]uint, len(l))

	for _, i := range l {
		r[i] = uint(l[i])
	}
	return r
}
