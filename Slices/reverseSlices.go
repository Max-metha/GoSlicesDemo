//Sortieren Sie ein Slice mit mindestens 10 Werten in umgekehrter Reihenfolge. 
//Geben Sie die umgekehrte Reihenfolge auf der Konsole aus.

package main

import (
	"fmt"
	"sort"
)

func main() {
	//Slice mit mindestens 10 Werten erstellen
	numbers := []int{4, 7, 1, 3, 5, 9, 8, 2, 6, 0}

	//Die sort.Reverse()-Funktion verwenden, um den Slice in umgekehrter Reihenfolge zu sortieren
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	//Ausgabe der sortierten Werte
	fmt.Println("Sortierte Werte in umgekehrter Reihenfolge:", numbers)
}