package phonenumber

func IsValid(phoneNumber string) bool {
	if len(phoneNumber) != 11 {
		return false
	}
	return true
}
