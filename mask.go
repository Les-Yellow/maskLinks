package main

import (
	"bufio"
	"fmt"
	"os"
)

func maskLinks(input string) string {
	// Преобразуем строку в байтовый срез
	inputBytes := []byte(input)
	outputBytes := make([]byte, len(inputBytes))
	copy(outputBytes, inputBytes) // Копируем данные в outputBytes

	// Ищем ссылки
	i := 0
	for i < len(inputBytes) {
		// Ищем начало ссылки
		if i+4 < len(inputBytes) && inputBytes[i] == 'h' && inputBytes[i+1] == 't' && inputBytes[i+2] == 't' && inputBytes[i+3] == 'p' && inputBytes[i+4] == ':' {
			j := i + 7 // Пропускаем "http://"

			// Продолжаем пока не достигнем пробела или конца строки
			for j < len(inputBytes) && inputBytes[j] != ' ' {
				j++
			}

			// Маскируем ссылку в выходном байтовом срезе
			for k := i + 7; k < j; k++ {
				outputBytes[k] = '*' // Заменяем на '*'
			}

			// Обновляем индекс i
			i = j
		} else {
			i++
		}
	}

	// Возвращаем результирующую строку из байтового среза
	return string(outputBytes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan() // на этом месте
	str := scanner.Text()

	output := maskLinks(str)
	fmt.Println("Output:", output) // Ожидаемый вывод: "Hello, its my page: http://**************** See you"
}
