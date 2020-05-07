/*
  Created By 简单7月 at 2020-05-07
*/
package main

import "fmt"

// 定义常量，提高应用程序的性能，避免了每次调用 Hello() 时都要创建 `Hello, ` 字符串实例
const englishHelloPrefix = "Hello, "

func Hello(some string) string {
	if some == "" {
		some = "world"
	}
	return englishHelloPrefix + some
}

func main() {
	fmt.Println(Hello("world"))
}
