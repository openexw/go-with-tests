/*
  Created By 简单7月 at 2020-05-07
*/
package iteration

const repeatCount = 5

func Repeat(str string) string {
	var res string
	var i int
	for i = 0; i < repeatCount; i++ {
		res += str
	}
	return res
}
