package bytes

func Trim0(bs []byte) (r []byte) {

	i, j := 0, 0
	for {
		if len(bs) > i {
			if bs[i] == 0 {
				j++
			}
			if j == 3 {
				return bs[:i-j+1]
			}
			i++
		} else {
			break
		}
	}
	return bs[:i-j+1]
}
