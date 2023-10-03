package main

import (
	"github.com/faissal20/go_base64/logic"
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

	// parameters --encode  --decode --help or -h -e -d
	// handle the parameters
	if(len(os.Args) >= 2){
		if(len(os.Args) > 2){
			switch os.Args[1] {
				case  "-e":
					input = os.Args[2]
				case  "-d":
					input = os.Args[2]
				case "-f":
					input = logic.ReadFile(os.Args[2])
				case  "-h":
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
				fmt.Println(logic.EncodeToBase64(binaries))
			case "--decode", "-d":
				// fmt.Println(logic.convertToBinary(binaries))
			default:
				fmt.Println(logic.EncodeToBase64(binaries))
		}
	}else{
		fmt.Println(logic.EncodeToBase64(binaries))
	}

	// create a file

	filename := logic.CreateTextFileFromBase64(logic.EncodeToBase64(binaries))

	fmt.Println("Done!")
	fmt.Println("File created: " + filename)
}