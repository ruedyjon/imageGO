package main

import (
	"flag"
	"fmt"
	"strings"

	imagego "github.com/rohitaryal/imageGO/pkg/imageGO"
	"github.com/rohitaryal/imageGO/structs"
)

func main() {
	// Modifying our usage function for better help menu
	flag.Usage = func() {
		fmt.Println("✨ imagego: Generate images using 'Imagen 3' from labs.google")
		flag.PrintDefaults()
		fmt.Println("Author: @rohitaryal")

		fmt.Println("\nAvailable Models:")
		fmt.Println(strings.Join(imagego.Models, ",  "))

		fmt.Println("\nAvailable Aspect Ratios:")
		fmt.Println(strings.Join(imagego.AspectRatios, ",  "))

		fmt.Println()
	}

	count := flag.Int("count", 1, "Number of images to generate")
	prompt := flag.String("prompt", "", "Image description")
	auth := flag.String("auth", "", "Authentication token")
	seed := flag.Int("seed", 0, "Seed for image")
	model := flag.String("model", "IMAGEN_3", "Name of model to use")
	aspect := flag.String("aratio", "IMAGE_ASPECT_RATIO_LANDSCAPE", "Aspect ratio of image")
	name := flag.String("name", "image", "Output file name. Will be assigned index for more than one images.")

	flag.Parse()

	if *prompt == "" {
		fmt.Println("[!] Image prompt is missing")
		flag.Usage()
		return
	}

	if len(*auth) < 10 {
		fmt.Println("[!] Authentication code is either invalid or missing")
		flag.Usage()
		return
	}

	if *count < 1 {
		fmt.Println("[W] Using '1' as default count value")
		*count = 1
	}

	req := structs.NewRequest{
		Count:       *count,
		Prompt:      *prompt,
		AuthCode:    *auth,
		AspectRatio: *aspect,
		Seed:        *seed,
		ModelName:   *model,
	}
	images, err := imagego.Generate(req)
	if err != nil {
		fmt.Println("Failed to generate images: ", err)
		return
	}

	for i, v := range images {
		err := v.Save(fmt.Sprintf("%s-%d.png", *name, i))
		if err != nil {
			fmt.Println("[!] Failed to save one of the image")
		} else {
			fmt.Println("[+] Image saved")
		}
	}
}
