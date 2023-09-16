package main

import (
	"fmt"
	"unicode/utf8"
)

//问题：建立一个程序统计字符串里的字符数量： asSASA ddd dsjkdsjs dk 同时输出这个字符串的字节数。 提示： 看看 unicode/utf8 包。
/*对于普通的ASCII编码字符串,它的字节数和字符数是一样的。

对于UTF-8编码的字符串,它的字节数和字符数可能不同。

UTF-8中的英文字符占1个字节
中文字符占3个字节
表情符号占4个字节
len()返回的是字节数,而不是字符数。
如果要获得字符数,需要使用utf8.RuneCountInString()。*/

func test(s string) (int, int) {
	l := len(s)                    //字节数
	u := utf8.RuneCount([]byte(s)) //字符数
	return l, u
}
func main() {
	s := "asSASA ddd dsjkdsjs dk"
	l, u := test(s)
	fmt.Printf("字节数=%v,字符数=%v", l, u)
}
