package helpers

func IsEmpty(str string) bool { return len(str) == 0 }

func IsEmail(str string) bool {
	return len(str) > 3 && str[len(str)-4:] == ".com"
}