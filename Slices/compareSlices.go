package main

import "fmt"

func main (){
	slices1 := []int {1,2,3,4}
	slices2 := []int {4,8,6,7}
	slices3 := []int {1,2,3,4}
	slices4 := []int {4,8,6,7}

	//equal :=compareSlices(slices1, slices2)

    fmt.Println(compareSlices(slices1, slices2))
	fmt.Println(compareSlices(slices1, slices3))
	fmt.Println(compareSlices(slices1, slices4))
}

//Die Funktion compareSlices vergleicht zwei Slices und gibt true zurück
//wenn die Slices Identisch sind oder false, wenn sie unterschiedlich sind
func compareSlices(slices1, slices2 []int) bool{
    //erster Vergleich ob beide Slices die selbe Länge haben
	//Operator "==" prüft auf Gleichheit
	//Operator "!=" prüft auf Ungleichheit
	if len(slices1) != len(slices2){
		
		return false
	}
	
	//zweiter Vergleich: sind beide Slices identisch
	for i := range slices1 {
        if slices1[i] != slices2[i] {
	    	
			return false
		}
	}
	//Slices sind gleich: generieren keine Rückgabewerte "return false" daher wird,
	//return true wird als Rückgabewert für die Funktion compareSlices generiert!
    return true

}