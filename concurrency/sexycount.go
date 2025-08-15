package concurrency

import (
	"fmt"
	"time"
)

func Test() {
	go sexyCount("Alice")
	// 만약 아래 두 함수 모두 go를 붙이면
	// goroutine은 main함수가 끝나기 전까지만
	// 실행되고 프로그램이 종료되기 때문에
	// 아무일도 일어나지 않는다.
	sexyCount("Kane")
	// sexyCount("Tommy")

	// 따라서 강제로 time.Sleep(time.Second * 6) 로
	// 종료를 막으면 모두 go를 붙여도 함수가 모두 정상 동작한다.

	fmt.Println("Channel Example======")
	// 그렇다면, goroutine 쓸 때마다 sleep 걸어줘야할까?
	// goroutine 응답값을 활용해야할 때는?
	// -> channel을 사용한다.
	ch := make(chan string)
	go convertToSexyName("Alice", ch)
	go convertToSexyName("Kane", ch)
	go convertToSexyName("Tommy", ch)
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func convertToSexyName(person string, ch chan string) {
	fmt.Println("START", person)
	time.Sleep(time.Second * 2)
	ch <- person + " is sexy"
	// ch input이 종료시점에 관련된 부분이라
	// 아래 END 출력되지 않을 수 있다.
	fmt.Println("END", person)
}

func sexyCount(person string) {
	for i := 0; i < 5; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
