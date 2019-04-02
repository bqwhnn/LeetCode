package main

import "sort"

func main() {

}

func sortedSquares(A []int) []int {
    r := []int{}
    for _, v := range A {
        r = append(r, v*v)
    }
    sort.Ints(r)
    return r
}

func sortedSquares_2(A []int) []int {
    r := []int{}
    t := []int{}
    a := -1
    for k, v := range A {
        v2 := v * v
        if v < 0 {
            t = append(t, v2)
            a = k
        } else {
            for i:=a; i>=0; i--{
                if t[i] <= v2 {
                    r = append(r, t[i])
                    a--
                } else {
                    break
                }
            }
            r = append(r, v2)
        }
    }
    for i:=a;i>=0;i--{
        r = append(r, t[i])
    }
    return r
}

func sortedSquares_3(A []int) []int {
    r := make([]int, len(A))
    head := 0
    tail := len(A)-1
    
    for i:= len(A)-1; i>=0; i-- {
        a := A[head] * A[head]
        b := A[tail] * A[tail]
        if a > b {
            r[i] = a
            head++
        } else {
            r[i] = b
            tail--
        }
    }
    return r
}