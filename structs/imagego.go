package structs

type NewRequest struct {
	Count       int
	Prompt      string
	Seed        int
	ModelName   string
	AspectRatio string
	AuthCode    string
}
