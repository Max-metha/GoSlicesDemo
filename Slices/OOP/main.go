package main

import (
	"fmt"
)
func main(){
	r:= Rectangle{width: 20, height: 10}

	fmt.Println("Die Flaeche betraegt: ", r.Aria())
	fmt.Println("Der Umfang betraegt: ", r.Perimeter())
}