package main

import "fmt"

func main() {

	//ARRAY
	multidimensional := [2][3]int{
		[3]int{11, 2, 3},
		[3]int{11, 2, 3},
	}

	fmt.Println(multidimensional, len(multidimensional))

	multidimensional2 := [2][3]float32{
		[3]float32{11, 2, 3},
		[3]float32{11, 2, 3},
	}

	fmt.Println(multidimensional2, len(multidimensional2))

	//===========================================================

	Tic_Tac_Toe := [3][3]bool{
		[3]bool{false, true, true},
		[3]bool{false, true, true},
		[3]bool{false, true, true},
	}

	fmt.Println(Tic_Tac_Toe, len(Tic_Tac_Toe))

	Tic_Tac_Toe[0] = [3]bool{true, false, true}

	fmt.Println(Tic_Tac_Toe, len(Tic_Tac_Toe))

	//===========================================================

	Tic_Tac_Toe[0] = [3]bool{}
	fmt.Println(Tic_Tac_Toe, len(Tic_Tac_Toe))

	//============================================================

	Tic_Tac_Toe[0][0] = true
	fmt.Println(Tic_Tac_Toe, len(Tic_Tac_Toe))

	//============================================================
	//SLICE

	multidimensional_slice := [][]int{
		[]int{21, 1, 5},
		[]int{5, 4, 84},
		[]int{54, 87, 8},
	}

	fmt.Println(multidimensional_slice, len(multidimensional_slice))

	//==============================================================

	multidimensional_slice[0] = []int{21, 5, 12}
	fmt.Println(multidimensional_slice, len(multidimensional_slice))

	//==============================================================

	multidimensional_slice[0] = []int{}
	fmt.Println(multidimensional_slice, len(multidimensional_slice))

	//==============================================================
	multidimensional_slice[1] = append(multidimensional_slice[2], 11)
	fmt.Println(multidimensional_slice, len(multidimensional_slice))
}
