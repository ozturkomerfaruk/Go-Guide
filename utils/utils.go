package utils

var Article = "The mountain gave birth to a mouse"

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f,l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return full, length
}