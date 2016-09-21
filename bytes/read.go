package bytes

func Trim0(bs []byte) (r []byte) {

	i, j := 0, 0
	for {
		if len(bs) > i {
			if bs[i] == 0 {
				j++
				if j == 3 {
					r = bs[:i-j+1]
					return
				}
			} else {
				j = 0
			}

			i++
		} else {
			break
		}
	}
	r = bs[:i-j]
	return
}
