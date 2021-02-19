package match

func MatchSOCKS5(buffer []byte) bool {
	if len(buffer) > 0 {
		return buffer[0] == 5
	}
	return false
}
