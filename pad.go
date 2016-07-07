package leftpad

// PadLeft takes the inputs string src and returns a new string with len length
// padded by the string padding
func PadLeft(src, padding string, length int) string {
	for {
		src = padding + src
		if len(src) > length-1 {
			return src[0:length]
		}
	}
}
