package main

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func abs(x int) int {
    mask := x >> 31
    ax := x ^ mask
    ax = ax - mask
    return ax
}
