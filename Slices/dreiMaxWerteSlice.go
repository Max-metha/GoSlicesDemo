//Legen Sie ein Slice mit mindestens 10 Werten an 
//geben auf der Konsole die drei größten Werte aus

package main
import "fmt"
import "sort"

func main() {
  // Anlegen eines Slices mit 10 Werten
  mySlice := []int{30, 50, 15, 20, 10, 45, 70, 1, 100, 500}

  // Sortieren des Slices
  sort.Ints(mySlice)

  // Ausgabe der größten drei Werte
  fmt.Println("Die größten drei Werte sind:", mySlice[len(mySlice)-1], mySlice[len(mySlice)-2], mySlice[len(mySlice)-3])
}