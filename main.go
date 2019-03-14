package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Ship struct {
	size   int
	points []Point
	sunk   bool
}

func (s *Ship) checkSunk() bool {
	for p := range s.points {
		if s.points[p].x != -1 || s.points[p].y != -1 {
			return false
		}
	}
	return true
}

func (s *Ship) init(size int) {
	s.size = size
	s.sunk = false
	s.points = make([]Point, size)
}

func (s *Ship) addPoint(xP, yP, index int) {
	s.points[index] = Point{x: xP, y: yP}

}

type Point struct {
	x, y int
}

var shipsSize = [5]int{2, 3, 3, 4, 5}
var boardLetters = [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
var playerShips []Ship
var enemyShips []Ship
var board [10][10]int
var enemyBoard [10][10]int
var waterColor = color.New(color.FgBlue)
var shipColor = color.New(color.FgRed)
var boardColor = color.New(color.FgHiBlack)
var errorColor = color.New(color.FgRed)
var gameOver = false

func main() {

	initBoard()
	for !gameOver {
		//TODO: add play again and stats
		playPlayerTurn()
		playEnemyTurn()
		exit := false
		for !exit {
			buf := bufio.NewReader(os.Stdin)
			fmt.Println("-----\nPlay Turn --> 1\nSee Stats --> 2\n------")
			pos, err := buf.ReadString('\n')
			instr := getMessageFromString(pos)
			if err != nil {
				fmt.Println(err)
			} else {
				if instr == "1" {
					exit = true
				} else if instr == "2" {
					fmt.Println("Your ships:")
					printShips(playerShips)
					fmt.Println("Enemy ships")
					printShips(enemyShips)
				}
			}
		}

	}
}

func playEnemyTurn() {
	if gameOver {
		return
	}
	fmt.Println("Enemy turn")
	hit := true
	var rX, rY int
	for hit {
		hit = false
		rX, rY = getRandomPoint()
		if board[rX][rY] == -1 || board[rX][rY] == -2 {
			hit = true
		}
	}
	checkPointHit(Point{x: rX, y: rY}, playerShips, &board)
	printPlayerBoard()

}

func printPlayerBoard() {
	fmt.Print("YOUR BOARD\n ")
	for i := range boardLetters {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	for i := range board {
		fmt.Print(boardLetters[i])
		for j := range board {
			if board[i][j] == -1 {
				shipColor.Print(" X")
			} else if board[i][j] == -2 {
				waterColor.Print(" W")
			} else {
				boardColor.Print(" -")
			}
		}
		fmt.Println()
	}
}

func playPlayerTurn() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Enter point to shot.")
	pos, err := buf.ReadString('\n')
	iP := getPointFromString(getMessageFromString(pos))
	point := Point{x: iP[0], y: iP[1]}
	if err != nil {
		fmt.Println(err)
	} else {
		if !checkPointHit(point, enemyShips, &enemyBoard) {
			playPlayerTurn()
		} else {
			printEnemyBoard()
		}

	}
}

func checkPointHit(point Point, ships []Ship, board *[10][10]int) bool {
	if point.x < 0 || point.x >= len(board) || point.y < 0 || point.y >= len(board) {
		return false
	}
	if board[point.x][point.y] == -1 || board[point.x][point.y] == -2 {
		fmt.Println("You already shot there!")
		return false
	}
	for s := range ships {
		for p := range ships[s].points {
			if ships[s].points[p].x == point.x && ships[s].points[p].y == point.y {
				board[point.x][point.y] = -1
				ships[s].points[p].x = -1
				ships[s].points[p].y = -1
				fmt.Printf("You hit a ship!\n")
				if ships[s].checkSunk() {
					ships[s].sunk = true
					fmt.Printf("You sunk a ship of size %d!\n", ships[s].size)
					checkWin()
				}
				return true
			}
		}
	}

	board[point.x][point.y] = -2
	fmt.Println("Nothing there!")
	return true
}

func checkWin() {
	playerWon := true
	for s := range enemyShips {
		if !enemyShips[s].sunk {
			playerWon = false
		}
	}
	if playerWon {
		fmt.Println("Player Won!!!")
		gameOver = true
	} else {
		enemyWon := true
		for s := range playerShips {
			if !playerShips[s].sunk {
				enemyWon = false
			}
		}

		if enemyWon {
			fmt.Println("Enemy Won!!!")
			gameOver = true
		}
	}
}

func initBoard() {
	for i := range board {
		for j := range board {
			board[i][j] = 0
		}
	}
	placePlayerShips()
	placeEnemyShips()

}

func placeEnemyShips() {
	fmt.Println("Placing Enemy ships...")
	enemyShips = make([]Ship, 5)
	i := 0
	for i < len(shipsSize) {
		if placeEnemyShip(shipsSize[i], i) {
			i++
		}
	}
	fmt.Println("Enemy ships placed!")
}

func getRandomPoint() (int, int) {

	rand.Seed(time.Now().UnixNano())
	randX := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	randY := rand.Intn(10)
	return randX, randY
}

func placeEnemyShip(size, shipIndex int) bool {
	rand.Seed(time.Now().UnixNano())
	orientation := rand.Intn(2)
	randX, randY := getRandomPoint()
	if orientation == 1 {
		//Vertically
		if canPlaceShipInBoard(randX, randY, orientation, size, enemyShips) {
			enemyShips[shipIndex].init(size)
			for i := 0; i < size; i++ {
				enemyShips[shipIndex].addPoint(randX+i, randY, i)
				enemyBoard[randX+i][randY] = 1
			}
		} else {
			return false
		}

	} else {
		if canPlaceShipInBoard(randX, randY, orientation, size, enemyShips) {
			enemyShips[shipIndex].init(size)
			for i := 0; i < size; i++ {
				enemyShips[shipIndex].addPoint(randX, randY+i, i)
				enemyBoard[randX][randY+i] = 1
			}
		} else {
			return false
		}
	}
	return true
}

func canPlaceShipInBoard(x, y, orientation, size int, ships []Ship) bool {
	if orientation == 1 {
		//Vertically
		if x+size > 10 {
			return false
		}
		for i := range ships {
			for p := range ships[i].points {
				for s := 0; s < size; s++ {
					if x+s == ships[i].points[p].x && y == ships[i].points[p].y {
						return false
					}
				}
			}
		}
	} else {
		//Horizontally
		if y+size > 10 {
			return false
		}
		for i := range ships {
			for p := range ships[i].points {
				for s := 0; s < size; s++ {
					if x == ships[i].points[p].x && y+s == ships[i].points[p].y {
						return false
					}
				}
			}
		}
	}
	return true
}

func placePlayerShips() {
	buf := bufio.NewReader(os.Stdin)
	playerShips = make([]Ship, 5)
	i := 0
	printPositioningPlayerBoard()
	for i < len(shipsSize) {
		fmt.Printf("Insert ship of size %d \n Horizontal or Vertical? (H/V)", shipsSize[i])
		orientation, err := buf.ReadString('\n')
		orientation = getMessageFromString(orientation)
		if err != nil {
			fmt.Println(err)
		} else {
			if orientation == "V" || orientation == "H" {
				fmt.Print("Enter position (A1)")
				pos, err := buf.ReadString('\n')
				pos = getMessageFromString(pos)
				if err != nil {
					fmt.Println(err)
				} else {
					if placeShipOnBoard(getPointFromString(getMessageFromString(pos)), shipsSize[i], orientation, i) {
						i++
						printPositioningPlayerBoard()
					}
				}
			} else {
				errorColor.Println("Incorrect orientation")
			}
		}
	}
}

func printShips(ships []Ship) {
	for i := range ships {
		fmt.Printf("Ship of size %d, sunk = %t\n", ships[i].size, ships[i].sunk)
	}
}

func printEnemyBoard() {
	fmt.Println("ENEMY BOARD\n ")
	for i := range boardLetters {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	for i := range enemyBoard {
		fmt.Print(boardLetters[i])
		for j := range enemyBoard {
			if enemyBoard[i][j] == -1 {
				shipColor.Print(" X")
			} else if enemyBoard[i][j] == -2 {
				waterColor.Print(" W")
			} else {
				boardColor.Print(" -")
			}
		}
		fmt.Println()
	}
}

func printPositioningPlayerBoard() {
	fmt.Print("YOUR BOARD\n ")
	for i := range boardLetters {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	for i := range board {
		fmt.Print(boardLetters[i])
		for j := range board {
			if board[i][j] == 0 {
				waterColor.Print(" O")
			} else if board[i][j] == 1 {
				shipColor.Print(" S")
			}
		}
		fmt.Println()
	}
}

func getMessageFromString(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	name := ""
	for i := range str {
		if (str[i] > 64 && str[i] < 91) || (str[i] > 96 && str[i] < 123) {
			name += string(str[i])
		} else if str[i] > 47 && str[i] < 58 {
			name += string(int(str[i]))
		}
	}
	return name
}

func getPointFromString(str string) [2]int {
	var pos [2]int
	str = strings.ToUpper(str)
	if len(str) != 2 {
		pos[0] = -1
		pos[1] = -1
		return pos
	}
	switch str[0] {
	case 'A':
		pos[0] = 0
	case 'B':
		pos[0] = 1
	case 'C':
		pos[0] = 2
	case 'D':
		pos[0] = 3
	case 'E':
		pos[0] = 4
	case 'F':
		pos[0] = 5
	case 'G':
		pos[0] = 6
	case 'H':
		pos[0] = 7
	case 'I':
		pos[0] = 8
	case 'J':
		pos[0] = 9
	default:
		pos[0] = 99
	}
	if len(str) == 2 {
		pos[1] = getIntFromByte(int(str[1]))
	} else {
		pos[1] = 99
	}
	return pos
}

func getIntFromByte(b int) int {
	return b - 48
}

func placeShipOnBoard(pos [2]int, size int, orientation string, shipIndex int) bool {

	if orientation == "V" {
		if pos[0] >= 0 && pos[0]+size-1 < len(board) && pos[1] >= 0 && pos[1] < len(board) &&
			canPlaceShipInBoard(pos[0], pos[1], 1, size, playerShips) {
			playerShips[shipIndex].init(size)
			for i := 0; i < size; i++ {
				board[pos[0]+i][pos[1]] = 1
				playerShips[shipIndex].addPoint(pos[0]+i, pos[1], i)
			}
		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	} else if orientation == "H" {
		if pos[1] >= 0 && pos[1]+size-1 < len(board) && pos[0] >= 0 && pos[0] < len(board) &&
			canPlaceShipInBoard(pos[0], pos[1], 2, size, playerShips) {
			playerShips[shipIndex].init(size)
			for i := 0; i < size; i++ {
				board[pos[0]][pos[1]+i] = 1
				playerShips[shipIndex].addPoint(pos[0], pos[1]+i, i)
			}
		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	}
	fmt.Println("Ship placed")
	return true
}
