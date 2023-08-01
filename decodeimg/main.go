package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	base64ImageData, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error al decodificar la imagen:", err)
		os.Exit(1)
	}

	imageData, err := base64.StdEncoding.DecodeString(string(base64ImageData))
	if err != nil {
		fmt.Println("Error al decodificar la imagen:", err)
		os.Exit(1)
	}
	width := 480
	height := 480
	img := image.NewGray(image.Rect(0, 0, width, height))
	var byteIndex int
	var bitOffset uint
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x%8 == 0 {
				byteIndex = (y*width + x) / 8
				bitOffset = 7
			}
			pixelValue := (imageData[byteIndex] >> bitOffset) & 1
			grayValue := uint8(pixelValue * 255)
			img.SetGray(x, y, color.Gray{Y: grayValue})
			bitOffset--
		}
	}

	// Guardar la imagen en formato PNG
	file, err := os.Create("imagen_decodificada.png")
	if err != nil {
		fmt.Println("Error al guardar la imagen:", err)
		os.Exit(1)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error al guardar la imagen:", err)
		os.Exit(1)
	}

	fmt.Println("Imagen decodificada y guardada correctamente en 'imagen_decodificada.png'.")
}
