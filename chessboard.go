package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	val := 0
	for _, v := range cb[rank] {
		if v {
			val += 1
		}
	}
	return val
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	val := 0
	if file >= 1 && file <= 8 {
		for _, v := range cb {
			if v[file-1] {
				val += 1
			}
		}
	}
	return val
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	val := 0
	for _, v := range cb {
		val += len(v)
	}
	return val
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	val := 0
	for i := range cb {
		val += CountInRank(cb, i)
	}
	return val
}
