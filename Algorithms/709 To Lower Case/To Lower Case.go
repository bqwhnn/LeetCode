package main

func main() {

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