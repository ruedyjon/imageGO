# imageGO
Re-implementation of lab.google's unofficial imageFX API in golang

## How to get authorization token?
1. Visit [imageFX](https://labs.google/fx/tools/image-fx) page
2. Open dev-tools by pressing <kbd>Ctrl</kbd> + <kbd>Shift</kbd> + <kbd>I</kbd> and paste the following code to extract your token

```javascript
let script = document.querySelector("#__NEXT_DATA__");
let obj = JSON.parse(script.textContent);
let authToken = obj["props"]["pageProps"]["session"]["access_token"];

window.prompt("Copy the auth token: ", authToken);
```

## Usage
### Using as module
```go
package main

import (
	"fmt"
	imagego "github.com/rohitaryal/imageGO/pkg/imageGO"
	"github.com/rohitaryal/imageGO/structs"
)

func main() {
	// Create a new request with parameters
	request := structs.NewRequest{
        // Optional
		Count:       1, // Number of images
		AuthCode:    "--YOUR_AUTH_CODE_HERE--",   
		Prompt:      "A mountain landscape at sunset",
		AspectRatio: "IMAGE_ASPECT_RATIO_LANDSCAPE", // Optional
		ModelName:   "IMAGEN_3",                     // Optional
		Seed:        42,                             // Optional
	}
	
	// Generate images using the API
	images, _ := imagego.Generate(request)
	
	// Save the first generated image as "test.png"
    // But make sure you check length first before indexing
	image[0].Save("test")
}
```

### Using as CLI
Download latest binary from release and run the following:
```bash
chmod +x imagego
./imagego -prompt "a cat" -auth "YOUR_AUTH_CODE_HERE"
```

## Build

```bash
git clone https://github.com/rohitaryal/imageGO
cd imageGO
go build -o imagego ./cmd/imageGO
```

## More usage
```yaml
✨ imagego: Generate images using 'Imagen 3' from labs.google
  -aratio string
    	Aspect ratio of image (default "IMAGE_ASPECT_RATIO_LANDSCAPE")
  -auth string
    	Authentication token
  -count int
    	Number of images to generate (default 1)
  -model string
    	Name of model to use (default "IMAGEN_3")
  -name string
    	Output file name. Will be assigned index for more than one images. (default "image")
  -prompt string
    	Image description
  -seed int
    	Seed for image
Author: @rohitaryal

Available Models:
IMAGEN_2,  IMAGEN_3,  IMAGEN_3_1,  IMAGEN_3_5,  IMAGEN_2_LANDSCAPE,  IMAGEN_3_PORTRAIT,  IMAGEN_3_LANDSCAPE,  IMAGEN_3_PORTRAIT_THREE_FOUR,  IMAGEN_3_LANDSCAPE_FOUR_THREE,  IMAGE_MODEL_NAME_UNSPECIFIED

Available Aspect Ratios:
IMAGE_ASPECT_RATIO_SQUARE,  IMAGE_ASPECT_RATIO_PORTRAIT,  IMAGE_ASPECT_RATIO_LANDSCAPE,  IMAGE_ASPECT_RATIO_UNSPECIFIED,  IMAGE_ASPECT_RATIO_LANDSCAPE_FOUR_THREE,  IMAGE_ASPECT_RATIO_PORTRAIT_THREE_FOUR
```