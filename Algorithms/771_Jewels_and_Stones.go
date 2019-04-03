package main

func test771() {

}

func numJewelsInStones(J string, S string) int {
    m := make(map[rune]bool)
    for _, s := range J {
        m[s] = true
    }
    count := 0
    for _, s := range S {
        if m[s] == true {
            count++
        }
    }
    return count
}