package main

import "fmt"

func main(){
    fmt.Println(isPermutation("omar", "ramk") == false)
    fmt.Println(isPermutation("omar", "ramo") == true)
}

func isPermutation(s1, s2 string) (bool){
   counter1 := make(map[rune]int)
   counter2 := make(map[rune]int)

   if len(s1) != len(s2){
       return false
   }

   for _, runeVal := range s1{
       counter1[runeVal] += 1
   }

   for _, runeVal2 := range s2{
       counter2[runeVal2] += 1
   }

   for k, val := range counter1{
       val2, ok := counter2[k]
       if !ok || val != val2{
           return false
       }
   }

   return true
}
