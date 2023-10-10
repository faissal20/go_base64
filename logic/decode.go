package logic


func PrepareString(str string) string {

	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			result = result + string(str[i]-65)
		} else if str[i] >= 97 && str[i] <= 122 {
			result = result + string(str[i]-71)
		} else if str[i] >= 48 && str[i] <= 57 {
			result = result + string(str[i]+4)
		} else if str[i] == 43 {
			result = result + string(62)
		} else if str[i] == 47 {
			result = result + string(63)
		}
	}

	return result
}
func DecodeFromBase64(base64 []byte) string {

	var i int = 0
	str := ""

	for i < len(base64) {
		if (i + 7) > len(base64) {
			break
		}

		oneString := base64[i : i+8]
		
		// convert the binary to decimal
		var decimal int = 0
		y := 0

		for j := len(oneString) - 1; j >= 0; j-- {
			if oneString[j] == 1 {
				decimal = decimal + (1 << uint(y))
			}
			y++
		}
		
		
		str = str + string(decimal)

		i = i + 8
	}

	return str
}