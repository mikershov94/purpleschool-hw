package email

func HashIsValid(hash string, target string) bool {
	if hash == target {
		return true
	}

	return false
}
