package main

import( "bufio"
		"strings"
		"fmt"
		"github.com/fatih/color"
		"math/rand"
		"time"
		"os")

type Ship struct{
	size int
	points []Point
	sunk bool
}

func (s *Ship) init(size int){
	s.size = size
	s.sunk = false
	s.points = make([]Point,size)
}

func (s *Ship) addPoint(xP,yP,index int){
	s.points[index] = Point{x:xP,y:yP}

}

type Point struct{
	x , y int
}


var shipsSize = [5]int{2,3,3,4,5}
var boardLetters = [10] string{"A","B","C","D","E","F","G","H","I","J"}
var playerShips [] Ship
var enemyShips[] Ship
var board [10][10]int
var enemyBoard [10][10] int
var waterColor = color.New(color.FgBlue)
var shipColor = color.New(color.FgHiBlack)
var errorColor = color.New(color.FgRed)

func main(){

	initBoard()
}

func initBoard(){
	for i:=range board{
		for j:=range board{
			board[i][j] = 0
		}
	}
	placePlayerShips()
	placeEnemyShips()

}

func placeEnemyShips(){
	fmt.Println("Placing Enemy ships...")
	enemyShips = make([]Ship,5)
	i:=0
	for i<len(shipsSize){
		if placeEnemyShip(shipsSize[i],i){ i++}
	}
	fmt.Println("Enemy ships placed! \n Enemy Ships:")
	printShips(enemyShips)
}

func placeEnemyShip(size, shipIndex int) bool{
	rand.Seed(time.Now().UnixNano())
	orientation := rand.Intn(2)
	rand.Seed(time.Now().UnixNano())
	randX := rand.Intn(10)
	rand.Seed(time.Now().UnixNano())
	randY := rand.Intn(10)
	if orientation == 1{
		//Vertically
		if canPlaceShipInBoard(randX,randY,orientation,size,enemyShips){
			enemyShips[shipIndex].init(size)
			for i:= 0;i<size;i++{
				enemyShips[shipIndex].addPoint(randX +i,randY,i)
				enemyBoard[randX + i][randY] = 1
			}
		} else {
			return false
		}

	} else {
		if canPlaceShipInBoard(randX,randY,orientation,size,enemyShips){
			enemyShips[shipIndex].init(size)
			for i:= 0;i<size;i++{
				enemyShips[shipIndex].addPoint(randX,randY +i,i)
				enemyBoard[randX ][randY +i] = 1
			}
		} else {
			return false
		}
	}
	return true
}

func canPlaceShipInBoard(x,y,orientation,size int,ships []Ship) bool{
	if(orientation==1){
		//Vertically
		if x+size>10{return false}
		for i:= range(ships){
			for p := range (ships[i].points){
				for s :=0;s<size;s++{
					if x +s ==ships[i].points[p].x && y == ships[i].points[p].y{return false}
				}
			}
		}
	} else {
		//Horizontally
		if y+size>10{
			return false
		}
		for i:= range(ships){
			for p := range (ships[i].points){
				for s :=0;s<size;s++{
					if x ==ships[i].points[p].x && y +s == ships[i].points[p].y{return false}
				}
			}
		}
	}
	return true
}

func placePlayerShips(){
	buf := bufio.NewReader(os.Stdin)
	playerShips = make([]Ship, 5)
	i:=0
	printPlayerBoard()
	for i<len(shipsSize){
		fmt.Printf("Insert ship of size %d \n Horizontal or Vertical? (H/V)", shipsSize[i])
		orientation, err := buf.ReadString('\n')
		orientation = getMessageFromString(orientation)
		if err != nil {
			fmt.Println(err)
		} else {
			if orientation == "V" || orientation == "H"{
				fmt.Print("Enter position (A1)")
				pos, err := buf.ReadString('\n')
				pos = getMessageFromString(pos)
				if err != nil {
				fmt.Println(err)
				} else {
					if placeShipOnBoard(getPointFromString(getMessageFromString(pos)), shipsSize[i], orientation, i){
						i++
						printPlayerBoard()
					}
				}
			} else {
				errorColor.Println("Incorrect orientation")
			}
		}
	}
	fmt.Println("Your Ships:")
	printShips(playerShips)
}

func printShips(ships []Ship){
	for i := range(ships){
		fmt.Printf("Ship of size %d, sunk = %t\n",ships[i].size,ships[i].sunk)
	}
}
func printEnemyBoard(){
	fmt.Println("ENEMY BOARD\n ")
	for i := range boardLetters{
		fmt.Printf(" %d",i)
	}
	fmt.Println()
	for i:=range enemyBoard{
		fmt.Print(boardLetters[i])
		for j:=range enemyBoard{
			if enemyBoard[i][j] == 0{
				waterColor.Print(" O")
			} else if enemyBoard[i][j] == 1{
				shipColor.Print(" S")
			}
		}
		fmt.Println()
	}
}
func printPlayerBoard(){
	fmt.Print("YOUR BOARD\n ")
	for i := range boardLetters{
		fmt.Printf(" %d",i)
	}
	fmt.Println()
	for i:=range board{
		fmt.Print(boardLetters[i])
		for j:=range board{
			if board[i][j] == 0{
				waterColor.Print(" O")
			} else if board[i][j] == 1{
				shipColor.Print(" S")
			}
		}
		fmt.Println()
	}
}

func getMessageFromString(str string) string{
	str = strings.Replace(str, "\n", "", -1)
	name:=""
	for i := range str{
		if (str[i] >64 && str[i]<91) || (str[i]>96 && str[i]<123){
			name += string(str[i])
		} else if str[i]>47 && str[i]<58{
			name += string(int(str[i]))
		}
	}
	return name
}

func getPointFromString(str string) [2]int{
	var pos [2]int
	str = strings.ToUpper(str)
	switch str[0]{
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
	if len(str) == 2{
		pos[1] =getIntFromByte(int(str[1]))
	} else {
		pos[1] = 99
	}
	return pos
}

func getIntFromByte(b int) int{
	return b - 48
}

func placeShipOnBoard(pos [2]int, size int, orientation string, shipIndex int) bool{

	if orientation == "V"{
		if pos[0] >= 0 && pos[0] + size -1 < len(board) && pos[1] >=0 && pos[1]<len(board) &&
		canPlaceShipInBoard(pos[0],pos[1],1,size,playerShips){
			playerShips[shipIndex].init(size)
			for i:= 0;i<size;i++{
					board[pos[0] + i][pos[1]] = 1
					playerShips[shipIndex].addPoint(pos[0] +i,pos[1],i)
			}
		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	} else if orientation == "H"{
		if pos [1] >=0 && pos [1] + size -1  < len(board) && pos[0] >=0 && pos[0]<len(board) &&
		canPlaceShipInBoard(pos[0],pos[1],2,size,playerShips){
			playerShips[shipIndex].init(size)
			for i:= 0;i<size;i++{
				board[pos[0]][pos[1] + i] = 1
				playerShips[shipIndex].addPoint(pos[0],pos[1]+i,i)
			}
		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	}
	fmt.Println("Ship placed")
	return true
}