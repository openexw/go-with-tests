/*
  Created By 简单7月 at 2020-05-07
*/
package _0_hello

import "fmt"

// 定义常量，提高应用程序的性能，避免了每次调用 Hello() 时都要创建 `Hello, ` 字符串实例
const englishHelloPrefix = "Hello, "

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "En"))
}
