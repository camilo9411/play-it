package yellowRed

const (
	NONE    = 0
	PLAYER1 = 1
	PLAYER2 = 2
	DRAW    = 3
	ROWS    = 6 - 1
	COLUMNS = 8 - 1
	CONNECT = 4
)

type Round struct {
	Winner       int    `json:"winner"`
	RowChoice    int    `json:"row_choice"`
	ColumnChoice int    `json:"column_choice"`
	Message      string `json:"msg"`
	Reset        bool   `json:"reset"`
}

var arr = [6][8]int{{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1}}

var isFirstPlayer bool = true

func PlayRound(playerColumn int) Round {
	msg := ""
	winner := NONE
	rowChoice := -1
	reset := false

	rowChoice = getAvailableRow(playerColumn)

	if isFirstPlayer {
		winner = checkWinner(PLAYER2)
	} else {
		winner = checkWinner(PLAYER1)
	}

	//if we have any winner we need to reset game
	if winner == PLAYER1 || winner == PLAYER2 || winner == DRAW {
		ResetGame()
		reset = true

	} else {
		reset = false
	}

	switch winner {
	case PLAYER1:
		msg = "1st Player won!"
		break
	case PLAYER2:
		msg = "2nd Player won!"
		break
	case DRAW:
		msg = "NO one won! This match is Draw."
		break
	default:
		break
	}

	var result Round
	result.Winner = winner
	result.RowChoice = rowChoice
	result.ColumnChoice = playerColumn
	result.Message = msg
	result.Reset = reset
	return result
}

func getAvailableRow(col int) int {
	for r := 0; r < len(arr); r++ {
		//if this row is empty we can return it
		if arr[r][col] == -1 {
			if isFirstPlayer {
				arr[r][col] = PLAYER1
			} else {
				arr[r][col] = PLAYER2
			}
			isFirstPlayer = !isFirstPlayer
			return r
		}
	}

	return -1
}

func ResetGame() {

	arr = [6][8]int{{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1}}

	isFirstPlayer = true

}

// Returns
//  NONE    = 0
// 	PLAYER1 = 1
// 	PLAYER2 = 2
// 	DRAW    = 3
func checkWinner(player int) int {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			var count = 0
			count = getVertical(r, c, player)
			if count >= CONNECT {
				return player
			}
			count = getHorizontal(r, c, player)
			if count >= CONNECT {
				return player
			}
			count = getDiagonalVerticalUpLeft(r, c, player)
			if count >= CONNECT {
				return player
			}
			count = getDiagonalVerticalUpRight(r, c, player)
			if count >= CONNECT {
				return player
			}
		}
	}
	return checkIsTie()
}

func checkIsTie() int {
	for c := 0; c <= COLUMNS; c++ {
		if arr[ROWS][c] == -1 {
			return NONE
		}
	}
	return DRAW
}

func getVertical(row, column, player int) int {
	var startRow int = 0
	if row-CONNECT >= 0 {
		startRow = row - CONNECT + 1
	}

	counter := 0
	for ; startRow <= row; startRow++ {
		if arr[startRow][column] == player {
			counter++
		} else {
			counter = 0
		}
	}

	return counter
}

func getHorizontal(row, column, player int) int {
	var endColumn int = COLUMNS
	if column+CONNECT <= COLUMNS {
		endColumn = column + CONNECT - 1
	}

	counter := 0
	for ; column <= endColumn; column++ {
		if arr[row][column] == player {
			counter++
		} else {
			counter = 0
		}
	}

	return counter
}

func getDiagonalVerticalUpLeft(row, column, player int) int {
	var startRow int = 0

	if row-CONNECT >= 0 {
		startRow = row - CONNECT + 1
	}

	var endColumn int = COLUMNS

	if column+CONNECT <= COLUMNS {
		endColumn = column + CONNECT
	}

	counter := 0
	for ; startRow <= row && column <= endColumn; column++ {
		if arr[row][column] == player {
			counter++
		} else {
			counter = 0
		}
		row--
	}

	return counter
}

func getDiagonalVerticalUpRight(row, column, player int) int {
	var endRow int = ROWS

	if row+CONNECT <= ROWS {
		endRow = row + CONNECT - 1
	}

	var endColumn int = COLUMNS

	if column+CONNECT <= COLUMNS {
		endColumn = column + CONNECT - 1
	}

	counter := 0
	for ; row <= endRow && column <= endColumn; column++ {
		if arr[row][column] == player {
			counter++
		} else {
			counter = 0
		}
		row++
	}

	return counter
}
