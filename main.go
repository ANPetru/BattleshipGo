package main

import( "bufio"
		"strings"
		"fmt"
		"github.com/fatih/color"
		"os")

		//"math/rand"
		//"time"

type Ship struct{
	size int
	points []Point
	sunk bool
}

type Point struct{
	x , y int
}


var ships = [5]int{2,3,3,4,5}
var board [10][10]int
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
	placeShips()
}

func placeShips(){
	buf := bufio.NewReader(os.Stdin)
	i:=0
	printPlayerBoard()
	//TODO: don't let ships be placed on other ship
	for i<len(ships){
		fmt.Printf("Insert ship of size %d \n Horizontal or Vertical? (H/V)", ships[i])
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
					if placeShipOnBoard(getPointFromString(getMessageFromString(pos)), ships[i], orientation){
						i++
						printPlayerBoard()
					}
				}
			} else {
				errorColor.Println("Incorrect orientation")
			}
		}
	}
	fmt.Println("Placing enemy ships")
	//TODO: place enemy ships
}

func printPlayerBoard(){
	fmt.Println("  A B C D E F G H I J")

	for i:=range board{
		fmt.Print(i)
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

func placeShipOnBoard(pos [2]int, size int, orientation string) bool{
	fmt.Println(pos)

	if orientation == "V"{
		if pos[0] >= 0 && pos[1] + size -1 < len(board) && pos[0] >=0 && pos[1]<len(board){

			for i:= 0;i<size;i++{
				board[pos[1] + i][pos[0]] = 1
			}

		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	} else if orientation == "H"{
		if pos [1] >=0 && pos [0] + size -1  < len(board) && pos[0] >=0 && pos[0]<len(board){

			for i:= 0;i<size;i++{
				board[pos[1]][pos[0] + i] = 1
			}

		} else {
			errorColor.Println("Incorrect position")
			return false
		}
	}
	fmt.Println("Ship placed")
	return true
}