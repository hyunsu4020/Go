package something

import "fmt"

func SayBye() { // Go 함수는 대문자가 공유가능, 소문자는 privete
	fmt.Println("bye")
}

func SayHello() {
	fmt.Println("Hello")
}
