package structs

// ================================================== REQUEST
// This is the request sent to server
type Request struct {
	UserInput     UserInput     `json:"userInput"`
	ClientContext ClientContext `json:"clientContext"`
	ModelInput    ModelInput    `json:"modelInput"`
	AspectRatio   string        `json:"aspectRatio"`
}

type UserInput struct {
	// Number of images requested
	CandidatesCount int `json:"candidatesCount"`
	// User's description about image(s)
	Prompts []string `json:"prompts"`
	// Seed value
	Seed int `json:"seed"`
}

type ClientContext struct {
	// Un-required field kept for sake of completion
	SessionId string `json:"sessionId"`
	// 'IMAGE_FX' for this repo
	Tool string `json:"tool"`
}

type ModelInput struct {
	// Imagen model name, to be precise
	ModelNameType string `json:"modelNameType"`
}

// =============================================== RESPONSE
// This is the response obtained from server
type Response struct {
	ImagePanels []ImagePanels `json:"imagePanels"`
}

type ImagePanels struct {
	// Your original prompt/image description
	Prompt string `json:"prompt"`
	// Array of generated image(s)
	GeneratedImages []GeneratedImage `json:"generatedImages"`
}

type GeneratedImage struct {
	// Actual image in base64 format
	EncodedImage string `json:"encodedImage"`
	// Seed value
	Seed              int    `json:"seed"`
	MediaGenerationId string `json:"mediaGenerationId"`
	IsMaskEditedImage bool   `json:"isMaskEditedImage"`
	Prompt            string `json:"prompt"`
	// Name of model used
	ModelNameType          string `json:"modelNameType"`
	WorkFlowId             string `json:"workflowId"`
	FingerprintLogRecordId string `json:"fingerprintLogRecordId"`
}
