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
	sexyCount("Tommy")

	// 따라서 강제로 time.Sleep(time.Second * 6) 로
	// 종료를 막으면 모두 go를 붙여도 함수가 모두 정상 동작한다.
}

func sexyCount(person string) {
	for i := 0; i < 5; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
