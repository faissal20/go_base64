package logic

import (
	"os"
)

func EncodeToBase64(binaries []byte) (string)  {
	i := 0
	str := ""
	for  i < len(binaries) {
		padded := false
		if (i + 5) > len(binaries) {
			// add 0s to the end
			for j := 0; j < 6 - (len(binaries) - i); j++ {
				binaries = append(binaries, 0)
			}
			padded = true
		}

		// get the 6 bits
		newBinary := binaries[i:i+6]
		// convert the binary to decimal
		var decimal int = 0
		y := 0

		for j := len(newBinary) - 1 ; j >= 0 ; j-- {
			if newBinary[j] == 1 {
				decimal = decimal + (1 << uint(y))
			}
			y++
		}
		// convert the decimal to ASCII  base64
		if decimal < 26 {
			str = str + string(decimal + 65)
		} else if decimal < 52 {
			str = str + string(decimal + 71)
		} else if decimal < 62 {
			str = str + string(decimal - 4)
		} else if decimal == 62 {
			str = str + "+"
		} else if decimal == 63 {
			str = str + "/"
		}

		// hendele paading
		if padded {
			str = str + "="
		}


		i = i + 6
	}

	return str
}

func CreateTextFileFromBase64(base64 string) (string) {
	
	fileName := "base64.txt"

	os.Create(fileName)

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	// write the base64 to the file
	file.WriteString(base64)

	defer file.Close()
	
	return fileName

}

func ReadFile(fileName string) (string) {
	// file, err := os.Open(fileName)
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(file)
}