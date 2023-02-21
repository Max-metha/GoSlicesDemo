
package main

import "fmt"

func main(){
	doubleNum := []int{1,2,3,4,5}
	doubleSliceElements (doubleNum)
	fmt.Print("Das ist der doppelte Wert von doubleNum: ", doubleNum)

}

func doubleSliceElements(doubleNum []int) {
    for i := range doubleNum {
		doubleNum[i] *= 2

	}
}