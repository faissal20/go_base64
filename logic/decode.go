package logic

func PrepareString(str string) string{
	result := ""
	for i := 0; i < len(str); i++ {
		if(str[i] >= 65 && str[i] <= 90){
			result = result + string(str[i] - 65)
		}else if(str[i] >= 97 && str[i] <= 122){
			result = result + string(str[i] - 71)
		}else if(str[i] >= 48 && str[i] <= 57){
			result = result + string(str[i] + 4)
		}else if(str[i] == 43){
			result = result + string(62)
		}else if(str[i] == 47){
			result = result + string(63)
		}
	}

	return result
}
func DecodeFromBase64(base64 []byte) string {
	var i int = 0

	var binaries []byte = []byte{}

	for i < len(base64) {
		
	}

	return string(binaries)
}