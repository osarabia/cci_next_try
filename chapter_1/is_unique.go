package main

import "fmt"

func main(){
    fmt.Println(isUnique("ramoo") == false)
    fmt.Println(isUnique("omar") == true)
}

func isUnique(s string) (bool){
    runes := make(map[rune]int)

    for _, runeVal := range s{
        if _,ok := runes[runeVal]; ok{
            return false
        }
        runes[runeVal] = 1
    }

    return true
}
