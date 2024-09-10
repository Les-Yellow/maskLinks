package main

import (
 "bufio"
 "fmt"
 "os"
)

func maskLinks(input string) string {
 inputBytes := []byte(input)
 outputBytes := make([]byte, len(inputBytes))
 copy(outputBytes, inputBytes) // Копируем данные в outputBytes

 i := 0
 for i < len(inputBytes) {
  // Проверяем, начинается ли с "http://" или "https://"
  if i+7 <= len(inputBytes) && string(inputBytes[i:i+7]) == "http://" {
   j := i + 7
   for j < len(inputBytes) && inputBytes[j] != ' ' {
    j++
   }

   // Маскируем ссылку в выходном байтовом срезе
   for k := i + 7; k < j; k++ {
    outputBytes[k] = '*' // Заменяем на '*'
   }

   i = j // Обновляем индекс i
  } else if i+8 <= len(inputBytes) && string(inputBytes[i:i+8]) == "https://" {
   j := i + 8
   for j < len(inputBytes) && inputBytes[j] != ' ' {
    j++
   }

   // Маскируем ссылку в выходном байтовом срезе
   for k := i + 8; k < j; k++ {
    outputBytes[k] = '*' // Заменяем на '*'
   }

   i = j // Обновляем индекс i
  } else {
   i++
  }
 }

 return string(outputBytes)
}

func main() {
 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Введите строку:")
 _ = scanner.Scan() // Считываем строку
 str := scanner.Text()

 output := maskLinks(str)
 fmt.Println("Output:", output) // Ожидаемый вывод с замаскированными адресами
}
