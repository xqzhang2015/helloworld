package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	// "golang.org/x/sync/singleflight"
	"myGo/internal/singleflight"
)

// https://github.com/golang/go/blob/master/src/internal/singleflight/singleflight_test.go
// https://github.com/golang/go/blob/master/src/net/lookup.go#L152
// https://blog.csdn.net/weixin_31422487/article/details/112672038

var (
	myGroup singleflight.Group
)

func main() {
	fmt.Printf("hello world\n")

	testDo()
	testDoChan()
}

func myHead(url string) func() (interface{}, error) {
	fn := func() (interface{}, error) {
		res, err := http.Head(url)
		return res, err
	}
	return fn
}

func myDo(url string) {
	urlKey := url + ":test"
	v, err, shared := myGroup.Do(urlKey, myHead(url))
	fmt.Printf("key: %v, v: %v (%T), err: %v, shared: %v\n", urlKey, v.(*http.Response).Status, v, err, shared)
}

// go run main.go
// === testDo() ===
// key: https://www.baidu.com:test, v: 200 OK (*http.Response), err: <nil>, shared: false
// key: https://golang.org:test, v: 200 OK (*http.Response), err: <nil>, shared: true
// key: https://golang.org:test, v: 200 OK (*http.Response), err: <nil>, shared: true
func testDo() {
	fmt.Println("=== testDo() ===")
	var wg sync.WaitGroup
	url1 := "https://golang.org"
	url2 := "https://www.baidu.com"

	wg.Add(1)
	go func(url string) {
		myDo(url)
		wg.Done()
	}(url1)

	wg.Add(1)
	go func(url string) {
		myDo(url)
		wg.Done()
	}(url1)

	wg.Add(1)
	go func(url string) {
		myDo(url)
		wg.Done()
	}(url2)

	wg.Wait()
}

func myDoChan(ctx context.Context, url string) {
	urlKey := url + ":test"
	ch, called := myGroup.DoChan(urlKey, myHead(url))

	fmt.Printf("fn is called: %v\n", called)

	select {
	case <-ctx.Done():
		if myGroup.ForgetUnshared(urlKey) {
			fmt.Printf("I'm the only goroutine looking up this key")
		} else {
			go func() {
				<-ch
				// child-ctx.Cancel()
			}()
		}
	case r := <-ch:
		fmt.Printf("key: %v, v: %v (%T), err: %v, shared: %v\n", urlKey, r.Val.(*http.Response).Status, r.Val, r.Err, r.Shared)
	}
}

func testDoChan() {
	fmt.Println("=== testDoChan() ===")

	var wg sync.WaitGroup
	url1 := "https://golang.org"
	url2 := "https://www.baidu.com"

	wg.Add(1)
	go func(url string) {
		myDoChan(context.Background(), url)
		wg.Done()
	}(url1)

	wg.Add(1)
	go func(url string) {
		myDoChan(context.Background(), url)
		wg.Done()
	}(url1)

	wg.Add(1)
	go func(url string) {
		myDoChan(context.Background(), url)
		wg.Done()
	}(url2)

	wg.Wait()
}
