// 实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。

package main

func test709() {

}

func toLowerCase(str string) string {
    c := []byte(str)
    for k, v := range c {
        if v >= 65 && v <= 90 {
            c[k] = v + 32
        }
    }
    return string(c)
}