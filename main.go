//Name: Vincent G.
//Date: 5/12/2020
//Description: Project #4 Game of Nim

package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Pile struct {
  First int
  Second int
}

func main() {

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

 A := Pile{First: 10}
 B := Pile{Second: 10}

 var inputA int
 var inputB int
 var turn int = 1

fmt.Println()
fmt.Println("----------------------------------------------------")

 //rules
 fmt.Println("1.) In this game there are two piles of ten pebbles.")
 fmt.Println()
 fmt.Println("2.) You can take up to 3 pebbles but only from one pile each turn.")
 fmt.Println()
 fmt.Println("3.) To win, be the person to take the last pebble.")
 
fmt.Println("----------------------------------------------------")

 //game begins
  for x :=1; A.First >=0 && B.Second >= 0;{
//Users Turn
if x%2 ==1 { 
  fmt.Println()
fmt.Println("Pile one =", A.First)
fmt.Println("Pile two =", B.Second)
fmt.Println()
fmt.Println("Enter '1' or '2' to choose which pile.")
fmt.Scanln(&inputA)
if inputA == 1 {
  fmt.Println()
  fmt.Println("Now choose between 1 and 3 pebbles to remove from Pile 1.")
  fmt.Scanln(&inputB)
A.First -= inputB 
x++
turn++
} else if inputA == 2 {
  fmt.Println()
  fmt.Println("Now choose between 1 and 3 pebbles to remove from Pile 2.")
  fmt.Scanln(&inputB)
B.Second -= inputB
x++
turn++
}
//Computers Turn
} else if x%2 !=1 {
  if rand.Intn(2)+1 == 1 {
  
  A.First -= r1.Intn(3)+1
  
  x++
  turn++
  } else if r1.Intn(2)+1 == 2 {
    B.Second -= r1.Intn(3)+1
  
  x++
  turn++
  } else if r1.Intn(2)+1 == 0  {
    
    A.First -= r1.Intn(3)+1
    x++
    turn++
  }
}

}
//End of Game
if A.First <=0 && B.Second <=0 && turn%2 ==1 {
  fmt.Println()
  fmt.Println("Game Over! You win!")
} else {
  fmt.Println()
  fmt.Println("Game Over! The computer wins.")
}
}