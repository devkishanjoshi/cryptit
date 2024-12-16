package encrypt

func Nimbus(str string) string {
	encryptedStr := ""

	for _, value := range str {
		asciiCode := int(value)
		asciiCharacter := string(asciiCode + 3)
		encryptedStr += asciiCharacter
	}

	return encryptedStr
}
