package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)


// HuffmanNode представляет узел дерева Хаффмана.
type HuffmanNode struct {
	char     rune
	frequency int
	left     *HuffmanNode
	right    *HuffmanNode
}

// PriorityQueue представляет приоритетную очередь для узлов Хаффмана.
type PriorityQueue []*HuffmanNode

// Функции для реализации интерфейса heap.Interface
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].frequency < pq[j].frequency }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}


func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}


// buildHuffmanTree строит дерево Хаффмана на основе частот символов.
func BuildHuffmanTree(frequencies map[rune]int) *HuffmanNode {
	pq := make(PriorityQueue, 0, len(frequencies))
	for char, freq := range frequencies {
		pq = append(pq, &HuffmanNode{char: char, frequency: freq})
	}
	heap.Init(&pq)

	for len(pq) > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		combined := &HuffmanNode{
			char:     0,
			frequency: left.frequency + right.frequency,
			left:     left,
			right:    right,
		}
		heap.Push(&pq, combined)
	}

	return pq[0]
}


// buildHuffmanCodes строит коды Хаффмана для каждого символа.
func BuildHuffmanCodes(root *HuffmanNode, code string, codes map[rune]string) {
	if root == nil {
		return
	}
	if root.char != 0 {
		codes[root.char] = code
	}
	BuildHuffmanCodes(root.left, code+"0", codes)
	BuildHuffmanCodes(root.right, code+"1", codes)
}

// compress сжимает текст с использованием кодов Хаффмана.
func Compress(text string, codes map[rune]string) string {
	var compressed strings.Builder

	for _, char := range text {
		compressed.WriteString(codes[char])
	}

	return compressed.String()
}


// convertToHexadecimal конвертирует бинарную строку в шестнадцатеричное представление.
func ConvertToHexadecimal(binary string) string {
	var hexBuilder strings.Builder

	for i := 0; i < len(binary); i += 4 {
		end := i + 4
		if end > len(binary) {
			end = len(binary)
		}
		substring := binary[i:end]

		value, _ := strconv.ParseInt(substring, 2, 64)
		hexBuilder.WriteString(fmt.Sprintf("%X", value))
	}

	return hexBuilder.String()
}

// decompress разжимает текст, используя дерево Хаффмана.
func Decompress(compressed string, root *HuffmanNode) string {
	var decompressed strings.Builder
	current := root

	for _, bit := range compressed {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current.char != 0 {
			decompressed.WriteRune(current.char)
			current = root
		}
	}
	return decompressed.String()
}


// Процесс сжатия
func RunHuffmanCompression(input string) (string, string) {
	// Подсчет частот символов
	frequencies := make(map[rune]int)
	for _, char := range input {
		frequencies[char]++
	}

	// Построение дерева Хаффмана
	root := BuildHuffmanTree(frequencies)

	// Построение кодов Хаффмана
	codes := make(map[rune]string)
	BuildHuffmanCodes(root, "", codes)

	// Сжатие текста
	compressedText := Compress(input, codes)

	// Конвертация в шестнадцатеричное представление
	hexCompressedText := ConvertToHexadecimal(compressedText)

	// Разжатие текста
	decompressedText := Decompress(compressedText, root)

	// Возвращаем результаты
	return hexCompressedText, decompressedText
}

func main() {
	input := "Hello World!"

	hexCompressedText, decompressedText := RunHuffmanCompression(input)

	fmt.Println("Original:", input)
	fmt.Println("Length original:", len(input))
	fmt.Println("Hex Compressed:", hexCompressedText)
	fmt.Println("Length Hex Compressed:", len(hexCompressedText))
	fmt.Println("Decompressed:", decompressedText)
	fmt.Println("len Decompressed:", len(decompressedText))
}
