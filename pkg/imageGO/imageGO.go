package imagego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"imageGO/internal/netops"
	"imageGO/structs"
	"net/http"
	"os"
	"slices"
)

var (
	reqUrl = "https://aisandbox-pa.googleapis.com/v1:runImageFx"
	// List of available models that are 100% expected to work.
	// Tested with burpsuite intruder :)
	ApectRatios = []string{
		"IMAGE_ASPECT_RATIO_SQUARE",
		"IMAGE_ASPECT_RATIO_PORTRAIT",
		"IMAGE_ASPECT_RATIO_LANDSCAPE",
		"IMAGE_ASPECT_RATIO_UNSPECIFIED",
		"IMAGE_ASPECT_RATIO_LANDSCAPE_FOUR_THREE",
		"IMAGE_ASPECT_RATIO_PORTRAIT_THREE_FOUR",
	}

	// List of available models that are 100% expected to work.
	// Tested with burpsuite intruder :)
	Models = []string{
		"IMAGEN_2",
		"IMAGEN_3",
		"IMAGEN_3_1",
		"IMAGEN_3_5",
		"IMAGEN_2_LANDSCAPE",
		"IMAGEN_3_PORTRAIT",
		"IMAGEN_3_LANDSCAPE",
		"IMAGEN_3_PORTRAIT_THREE_FOUR",
		"IMAGEN_3_LANDSCAPE_FOUR_THREE",
		"IMAGE_MODEL_NAME_UNSPECIFIED",
	}
)

// Creates a generate request to imagefx server
func Generate(req structs.NewRequest) ([]structs.GeneratedImage, error) {
	if len(req.AuthCode) < 10 {
		fmt.Println("[!] Invalid authentication code")
		os.Exit(1)
	}

	// Some defaults
	if req.AspectRatio == "" {
		req.AspectRatio = "IMAGE_ASPECT_RATIO_LANDSCAPE"
	}

	if req.ModelName == "" {
		req.ModelName = "IMAGEN_3_1"
	}

	if !slices.Contains(Models, req.ModelName) {
		fmt.Println("[!] Invalid model selected")
		os.Exit(1)
	}

	if !slices.Contains(ApectRatios, req.AspectRatio) {
		fmt.Println("[!] Invalid aspect ratio selected")
		os.Exit(1)
	}

	// Create request in proper schema
	reqBody := structs.Request{
		UserInput: structs.UserInput{
			CandidatesCount: req.Count,
			Prompts:         []string{req.Prompt},
			Seed:            req.Seed,
		},
		ClientContext: structs.ClientContext{
			SessionId: ";1747730436522",
			Tool:      "IMAGE_FX",
		},
		ModelInput: structs.ModelInput{
			ModelNameType: req.ModelName,
		},
		AspectRatio: req.AspectRatio,
	}

	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to stringify POST body: %w", err)
	}

	r, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}

	// Dont forget these hehe
	r.Header.Set("Authorization", req.AuthCode)
	r.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	r.Header.Set("Accept", "*/*")

	// Make a request:w http.ResponseWriter, r *http.Request
	res, err := netops.Fetch(r)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	// Parse to Response type
	var parsedRes structs.Response
	err = json.Unmarshal([]byte(res), &parsedRes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse to JSON: %w", err)
	}

	return parsedRes.ImagePanels[0].GeneratedImages, nil
}
