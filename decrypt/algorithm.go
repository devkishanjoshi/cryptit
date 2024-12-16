package decrypt

func Nimbus(str string) string {
	decryptedStr := ""

	for _, value := range str {
		asciiCode := int(value)
		asciiCharacter := string(asciiCode - 3)
		decryptedStr += asciiCharacter
	}

	return decryptedStr
}
