package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	fmt.Println("1 - Encode, 2 - Decode")
	var ch int8
	_, err := fmt.Scan(&ch)
	if err != nil {
		panic(err)
	}

	if ch == 1 {
		encode()
	} else if ch == 2 {
		decode()
	} else {
		return
	}
}

func encode() {
	fmt.Println("Enter the key") // getting the key
	var key uint32
	_, err := fmt.Scan(&key)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Open("photo.png") // opening the photo
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, err := png.Decode(inputFile) // decoding it
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	outputFile, err := os.Create("enc.bin") // creating encoded file
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	// photo sizw
	binary.Write(writer, binary.LittleEndian, int32(width))
	binary.Write(writer, binary.LittleEndian, int32(height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(bounds.Min.X+x, bounds.Min.Y+y).RGBA()
			// converting to 64 bit
			r32 := uint32(uint8(r >> 8))
			g32 := uint32(uint8(g >> 8))
			b32 := uint32(uint8(b >> 8))
			a32 := uint32(uint8(a >> 8))

			encR := r32 ^ key
			encG := g32 ^ key
			encB := b32 ^ key
			encA := a32 ^ key

			binary.Write(writer, binary.LittleEndian, encR)
			binary.Write(writer, binary.LittleEndian, encG)
			binary.Write(writer, binary.LittleEndian, encB)
			binary.Write(writer, binary.LittleEndian, encA)
		}
	}
	writer.Flush()
	fmt.Println("Photo encrypted succesfully")
}

func decode() {
	fmt.Println("Enter the key") // getting the key
	var key uint32
	_, err := fmt.Scan(&key)
	if err != nil {
		panic(err)
	}

	encodedFile, err := os.Open("enc.bin") // opening enc file
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(encodedFile)
	var width, height int32
	binary.Read(reader, binary.LittleEndian, &width)
	binary.Read(reader, binary.LittleEndian, &height)
	decryptedPhoto := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			var encR, encG, encB, encA uint32

			binary.Read(reader, binary.LittleEndian, &encR)
			binary.Read(reader, binary.LittleEndian, &encG)
			binary.Read(reader, binary.LittleEndian, &encB)
			binary.Read(reader, binary.LittleEndian, &encA)

			r := uint8(encR ^ key)
			g := uint8(encG ^ key)
			b := uint8(encB ^ key)
			a := uint8(encA ^ key)
			decryptedPhoto.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}
	outputFile, err := os.Create("restored_photo.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, decryptedPhoto)
	if err != nil {
		panic(err)
	}
	fmt.Println("Photo decrypted succesfully")
}
