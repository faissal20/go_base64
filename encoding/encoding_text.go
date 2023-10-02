package encoding

import (
	
)

func convertToBinary(binaries []byte) (string) {
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