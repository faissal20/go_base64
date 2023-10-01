package main

import 
(
	"fmt"
	"os"
)

func convertToBinary(integer byte) ([8]byte) {
	result := [8]byte{}
	
	counter := 7
	for integer > 0 {
		result[counter] = integer % 2
		integer = integer / 2
		counter--
	}
	
	return result

}
	


func main() {
	
	// get the input
	var input string
	
	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		fmt.Println("Please enter a string to encode")
		return
	}
	// get the bytes of the string
	bytes := []byte(input)
	
	binaries := []byte{}


	for i := 0; i < len(bytes); i++ {
		
		binary := convertToBinary(bytes[i])
		// append the binary to the array
		for j := 0; j < 8; j++ {
			binaries = append(binaries, binary[j])
		}

	}

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

	fmt.Println(str)

}