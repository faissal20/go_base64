package main

import 
(
	"fmt"
	"os"
	"github.com/faissal20/go_base64/encoding/enconding_text"
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

	// parameters --encode  --decode --help or -h -e -d
	// handle the parameters
	if(len(os.Args) >= 2){
		if(len(os.Args) > 2){
			switch os.Args[1] {
				case "--encode", "-e":
					input = os.Args[2]
				case "--decode", "-d":
					input = os.Args[2]
				case "--help", "-h":
					fmt.Println("Usage: base64 [OPTION]... [FILE]")
			}
		}else{
			input = os.Args[1]
		}
	}else{
		fmt.Println("Usage: base64 [OPTION]... [FILE]")
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

	if (len(os.Args) >= 2){
		switch os.Args[1] {
			case "--encode", "-e":
				fmt.Println(enconding_text.ConvertToBase64(binaries))
			case "--decode", "-d":
				fmt.Println(enconding_text.ConvertFromBase64(binaries))
			default:
				fmt.Println(enconding_text.ConvertToBase64(binaries))
		}
	}else{
		fmt.Println(enconding_text.ConvertToBase64(binaries))
	}


}