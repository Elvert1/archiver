package main

import "testing"

func TestRunHuffmanCompression(t *testing.T) {
	text := "Hello World!"

	// процесс сжватия Huffman 
	hexCompressedText, decompressedText := RunHuffmanCompression(text)

	// Проверка чтобы оригинал совпадал с разархивированным файлом
	if decompressedText != text {
		t.Errorf("Разархивированная строка отличаетя от оригинала. Оригинал: %s, Результат: %s", text, decompressedText)
	}

	// Проверка чтобы длина архива была меньше оригинала
	if len(hexCompressedText) >= len(text) {
		t.Errorf("Длина архива не меньше длины оригинала. Длина оригинала: %d, Длина сжатого: %d", len(text), len(hexCompressedText))
	}
}
