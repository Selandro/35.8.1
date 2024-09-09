package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var pogovorki = []string{
	"Don’t communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Закрыть соединение в конце работы

	for {
		// Генерация случайной поговорки
		pogovorka := pogovorki[rand.Intn(len(pogovorki))]

		// Отправка поговорки клиенту с добавлением возврата каретки и новой строки
		_, err := fmt.Fprintf(conn, "%s\r\n", pogovorka)
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
			return
		}

		// Ожидание 3 секунд перед отправкой следующей поговорки
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Прослушивание порта 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080...")

	for {
		// Ожидание подключения клиента
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении клиента:", err)
			continue
		}

		// Обработка каждого подключения в отдельной горутине
		go handleConnection(conn)
	}
}
