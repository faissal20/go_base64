package main

import (
	"github.com/faissal20/go_base64/logic"
	"fmt"
	"os"
)

func ConvertToBinary(integer byte) ([8]byte) {
	result := [8]byte{}
	
	counter := 7
	for integer > 0 {
		result[counter] = integer % 2
		integer = integer / 2
		counter--
	}
	
	return result

}

func ConvertToBinaryBase64(integer byte) ([6]byte){
	result := [6]byte{}
	
	counter := 5
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
		fmt.Println("Usage: base64 [OPTION]... [FILE|STRING]");
	}

	
	// get the bytes of the string
	bytes := []byte(input)
	
	

	if (len(os.Args) >= 2){
		switch os.Args[1] {
			case "--encode", "-e":
				binaries := []byte{}


				for i := 0; i < len(bytes); i++ {
					
					binary := logic.ConvertToBinary(bytes[i])
					// append the binary to the array
					for j := 0; j < 8; j++ {
						binaries = append(binaries, binary[j])
					}

				}

				fmt.Println(logic.EncodeToBase64(binaries))
				filename := logic.CreateTextFileFromBase64(logic.EncodeToBase64(binaries))

				fmt.Println("Done!")
				fmt.Println("File created: " + filename)

			case "--decode", "-d":

				stringPrepared := logic.PrepareString(input)
			
				binaries := []byte{}

				for i := 0; i < len(stringPrepared); i++ {

					binary := logic.ConvertToBinaryBase64(stringPrepared[i])
					// append the binary to the array
					for j := 0; j < 6; j++ {
						binaries = append(binaries, binary[j])
					}

				}

				fmt.Println(logic.DecodeFromBase64(binaries)) 

			default:
				fmt.Println("still not implemented")
		}
	}else{
		fmt.Println("Usage: base64 [OPTION]... [FILE]")
	}

	
}