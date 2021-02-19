package match

func MatchSOCKS4(buffer []byte) bool {
	if len(buffer) > 0 {
		return buffer[0] == 4
	}
	return false
}
