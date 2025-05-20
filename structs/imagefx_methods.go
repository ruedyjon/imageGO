package structs

import (
	"encoding/base64"
	"fmt"
	"imageGO/internal/util"
	"os"
)

// Simple logger for request
func (req Request) Log() {
	fmt.Println("----------- REQUEST ----------")

	fmt.Println("[+] Prompt: ", req.UserInput.Prompts)
	fmt.Println("[+] Total Requested Images: ", req.UserInput.CandidatesCount)
	fmt.Println("[+] Model Chosen: ", req.ModelInput.ModelNameType)
}

// Simple logger for response
func (res Response) Log() {
	generatedImages := res.ImagePanels[0].GeneratedImages

	fmt.Println("----------- RESPONSE ----------")

	fmt.Println("[+] Total Images: ", len(generatedImages))
	fmt.Print("[+] Model Used: ")

	if len(generatedImages) > 0 {
		fmt.Println(generatedImages[0].ModelNameType)
		fmt.Println("[+] Prompt: ", generatedImages[0].Prompt)
		fmt.Println("[+] Status: Success")
	} else {
		fmt.Println("UNKNOWN")
		fmt.Println("[+] Status: Failed")
	}
}

// Reciever for a generated image
func (g GeneratedImage) Save(fileName string) error {
	decoded, err := base64.StdEncoding.DecodeString(g.EncodedImage)
	util.CheckForFailure(err)

	return os.WriteFile(fileName, decoded, 0664) // Normal permisison
}
