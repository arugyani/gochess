package util

func Collapse(arr []uint64) (output uint64) {
	for _, v := range arr {
		output |= v
	}

	return
}
