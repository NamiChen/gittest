package game

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

const _rows, _cols = 4, 4

//Board is
type Board interface {
	Display()
}
type board struct {
	board  [][]int
	nx, ny int
}

type key int

const (
	UP key = iota
	DOWN
	LEFT
	RIGHT
	QUIT
	Error_Key
)

func (b *board) TakeInput() bool {
	key, err := GetKeyStrokes()
	if err != nil {
		fmt.Printf(err.Error())
	}
	if key == Error_Key {
		b.TakeInput()
	}
	switch key {
	case UP:
		b.moveUp()
	case DOWN:
		b.moveDown()
	case LEFT:
		b.moveLeft()
	case RIGHT:
		b.moveRight()
	case QUIT:
		fmt.Print("You press ESC ,game exit")
		return false
	}
	return true
}
func (b *board) moveLeft() {
	for i := 0; i < _rows; i++ {
		old := b.board[i]
		b.board[i] = moveRow(old)
	}
}
func (b *board) moveRight() {
	b.Reverse()
	b.moveRight()
	b.Reverse()
}

func (b *board) moveUp() {
	b.leftRotate90()
	b.moveLeft()
	b.rightRotate90()
}
func (b *board) moveDown() {
	b.rightRotate90()
	b.moveLeft()
	b.leftRotate90()
}
func (b *board) leftRotate90() {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[_cols-1-j][i] = b.board[i][j]
		}
	}
	b.board = matrix

}
func (b *board) rightRotate90() {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[j][_cols-1] = b.board[i][j]
		}
	}
	b.board = matrix

}

func (b *board) Reverse() {
	for i := 0; i < _rows; i++ {
		for j, K := 0, _cols; j < K; j, K = j+1, K-1 {
			b.board[i][j], b.board[i][K] = b.board[i][K], b.board[i][j]
		}
	}
}

func moveRow(row []int) []int {
	index := 0
	for i := 0; i < len(row); i++ {
		if row[i] != 0 {
			row[index], row[i] = row[i], row[index]
			index++
		}
	}
	for i := 0; i < len(row)-1; i++ {
		if row[i] == row[i+1] {
			row[i] += row[i+1]
			row[i+1] = 0
			i++
		}
	}
	index = 0
	for i := 0; i < len(row); i++ {

		if row[i] != 0 {
			row[index], row[i] = row[i], row[index]
			index++
		}
	}
	return row
}

func GetKeyStrokes() (key, error) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		return Error_Key, err
	}
	if int(char) == 0 {
		switch key {
		case keyboard.KeyArrowUp:
			return UP, nil
		case keyboard.KeyArrowDown:
			return DOWN, nil
		case keyboard.KeyArrowLeft:
			return LEFT, nil
		case keyboard.KeyArrowRight:
			return RIGHT, nil
		case keyboard.KeyEsc:
			return QUIT, nil
		default:
			return Error_Key, errors.New("Invalid ,please press again")
		}

	} else {
		switch char {
		case 119:
			return UP, nil
		case 97:
			return LEFT, nil
		case 115:
			return DOWN, nil
		case 100:
			return RIGHT, nil
		default:
			return Error_Key, errors.New("Invalid ,please press again")
		}
	}

}

func (b *board) AddElement() {
	rand.Seed(time.Now().UnixNano())
	index := make([][2]int, 0)
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				index = append(index, [2]int{i, j})
			}

		}
	}
	next := rand.Int() % len(index)
	nx, ny := index[next][0], index[next][1]
	//number
	var number int
	if rand.Int()%100 < 80 {
		number = 2
	} else {
		number = 4

	}
	b.nx, b.ny = nx, ny
	b.board[nx][ny] = number
}

func (b *board) Display() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	b.board = generate()
	printHorizontalLine()
	for i := 0; i < _rows; i++ {
		printHorizontalLine()
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				fmt.Printf("%7s", "")
			} else {
				fmt.Printf("%4d%3s", b.board[i][j], "")
			}
			printVertical()
		}
		fmt.Println()
		printHorizontalLine()
	}
	// fmt.Println(b)
	// fmt.Println("aaa")
}

func printHorizontalLine() {
	for i := 0; i < 33; i++ {
		fmt.Printf("-")

	}
	fmt.Println()

}

func printVertical() {
	fmt.Printf("|")
}

//生成需要的随机数
func generate() [][]int {
	// Store all available numbers from 2 to 2048
	nums := make([]int, 0)
	nums = append(nums, 0)
	for i := 2; i <= 2048; i *= 2 {
		nums = append(nums, i)
	}

	//generate random numbers for init board
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[i][j] = nums[rand.Int()%len(nums)]
		}
	}
	return matrix
}

//NewBoard is
func NewBoard() *board {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	return &board{board: matrix}
}

func (b *board) CountScore() (int, int) {
	total, max := 0, 0
	matrix := b.board
	for i := 0; i < _rows; i++ {
		for k := 0; k < _cols; k++ {
			total += matrix[i][k]
			max = maxInts(max, matrix[i][k])
		}
	}
	return max, total
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *board) IsOver() bool {
	blank := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				blank++
			}
		}
	}
	return blank == 0
}
