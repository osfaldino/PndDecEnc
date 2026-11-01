
English

A small experimental tool written in Go. It reads a PNG image, converts pixels to bytes, obfuscates them with an XOR key, and saves everything into a custom .bin file. To get your image back, just run the decode mode with the same key.

How it works

1. Header: The program writes the image width and height as int32 fields at the very beginning of the .bin file.
2. Pixel Data: It iterates over every pixel, converts RGBA channels to 8-bit values, applies ^ key (XOR), and writes them as uint32 values.
3. Decryption: Reads the dimensions, processes the stream of 32-bit integers, reverses the XOR operation, and packs the pixels back into a clean PNG.

Prerequisites

1. Go 1.16 or higher
2. An input image named exactly photo.png in the working directory

Usage

1. Run the program:bash
go run main.go
Используйте код с осторожностью.
2. Choose action: Enter 1 for encryption or 2 for decryption.
3. Enter the key: Type any numeric key (e.g., 42, 1337).
! When encrypting, it looks for photo.png and creates enc.bin.
! When decrypting, it looks for enc.bin and creates restored_photo.png.
