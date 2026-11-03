
English

A small experimental tool written in Go. It reads a PNG image, converts pixels to bytes, obfuscates them with an XOR key, and saves everything into a custom .bin file. To get your image back, just run the decode mode with the same key.

**How it works**

1. Header: The program writes the image width and height as int32 fields at the very beginning of the .bin file.
2. Pixel Data: It iterates over every pixel, converts RGBA channels to 8-bit values, applies ^ key (XOR), and writes them as uint32 values.
3. Decryption: Reads the dimensions, processes the stream of 32-bit integers, reverses the XOR operation, and packs the pixels back into a clean PNG.

**Prerequisites**

1. Go 1.16 or higher
2. An input image named exactly photo.png in the working directory

**Usage**

1. Run the program:bash
go run main.go
2. Choose action: Enter 1 for encryption or 2 for decryption.
3. Enter the key: Type any numeric key.
   
When encrypting, it looks for photo.png and creates enc.bin.

When decrypting, it looks for enc.bin and creates restored_photo.png.

-------------------------------------
Русский

Небольшая экспериментальная утилита на Go. Она считывает PNG картинку, разбирает пиксели на байты, шифрует их обычным XOR по ключу и собирает всё это в кастомный бинарный файл .bin. Чтобы вернуть картинку обратно, нцжно запустить режим декодирования с тем же самым ключом.

**Как это работает**

1. Заголовок: В самое начало .bin файла программа записывает ширину и высоту картинки в формате int32.
2. Данные пикселей: Код проходится по каждому пикселю, приводит RGBA-каналы к 8-битным значениям, применяет к ним ^ key (XOR) и сохраняет как числа uint32.
3. Расшифровка: Читает размеры из заголовка, обрабатывает поток 32-битных целых чисел, делает обратный XOR и собирает пиксели обратно в чистый PNG.

**Что нужно для запуска**

1. Установленный Go версии 1.16 или выше
2. Исходная картинка строго с именем photo.png в рабочей директории

**Инструкция**

1. Запустите программу:bash
go run main.go

2. Выберите действие: Введите 1 для шифрования или 2 для расшифровки.
3. Введите ключ: Наберите любое число.
   
При шифровании программа ищет photo.png и создает enc.bin.

При расшифровке она берет enc.bin и создает restored_photo.png.
