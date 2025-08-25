package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	num := 0
    
    for _, b := range cb[file] {
        if b {
            num++
        }
    }

    return num
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
	
    num := 0
    var fileName string
    
	for r := 'A'; r <= 'H'; r++ {
        fileName = string(r)
        if cb[fileName][rank - 1] {
            num++
        }
    }

    return num
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	num := 0

	for _, file := range cb {
        for range file {
            num++
        }
    }

    return num
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	num := 0
    var fileName string

	for r := 'A'; r <= 'H'; r++ {
        fileName = string(r)
        num += CountInFile(cb, fileName)
    }

    return num
}
