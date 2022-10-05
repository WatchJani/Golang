package main


import "fmt"



// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int 

	for _, value:= range cb[file]{
			if value == true{
				counter++
			}
	}

	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	if rank < 1 || rank > 8{
		return 0
	}

	var counter int 

	for _, value:= range cb{
		// fmt.Println(value)
		if value[rank-1] == true{
			counter++
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var counter int 

	for _,column:= range cb{
		for range column{
			counter++
		}
	}

	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int 

	for _,column:= range cb{
		for _, row := range column{
			if row == true{
				counter++
			}
		}
	}

	return counter
}


func main(){

	//   A B C D E F G H
	// 8 # _ _ _ # _ _ # 8
	// 7 _ _ _ _ _ _ _ _ 7
	// 6 _ _ _ _ # _ _ # 6
	// 5 _ # _ _ _ _ _ # 5
	// 4 _ _ _ _ _ _ # # 4
	// 3 # _ # _ _ _ _ # 3
	// 2 _ _ _ _ _ _ _ # 2
	// 1 # _ _ _ _ _ _ # 1
	//   A B C D E F G H

	var board Chessboard
	board = map[string]File{
		"A": File([]bool{true, false, true, false, false, false, false, true}),
		"B": File([]bool{false, false, false, false, true, false, false, false}),
		"C": File([]bool{false, false, false, true, false, false, false, false}),
		"D": File([]bool{false, false, false, false, false, false, false, false}),
		"E": File([]bool{false, false, false, false, false, true, false, true}),
		"F": File([]bool{false, false, false, false, false, false, false, false}),
		"G": File([]bool{false, false, false, true, false, false, false, false}),
		"H": File([]bool{true, true, true, true, true, true, false, true}),
	}
	

	fmt.Println(CountInFile(board, "A"))
	fmt.Println(CountInRank(board, 99))
	fmt.Println(CountAll(board))
}