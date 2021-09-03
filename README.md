# BattleshipGo
Battleship board game for command line made in Go.


## Index

[1. About](#about)

[2. Demo](#demo)

[3. How to play](#play)

[4. Improvement ideas](#ideas)


<a name="about"/>

## About

Battleship is a strategy type guessing game for two players. It is played on ruled grids on which each player's fleet of ships are marked. The locations of the fleets are 
concealed from the other player.

Players alternate turns calling "shots" at the other player's ships, and the objective of the game is to destroy opossing player's fleet.


<a name="demo"/>

## Demo

#### Placing your fleet

![Demo placing](https://j.gifs.com/XLzMNl.gif)

#### Shooting enemy's fleet

![Demo shooting](https://j.gifs.com/6Xq0PQ.gif)

#### Winning

![Demo winning](https://j.gifs.com/gZjkll.gif)

<a name="play"/>

## How to play

Download this project and extract it in your folder of choice.

### Linux

Go to `Platforms/Linux` in the command line and run `./battleship`.

### Windows

Go to `Platforms\Windows` in the file explorer and run `battleship`.

<a name="ideas"/>

## Improvement ideas

### Player vs player

Currently, you can only play against the IA, it would be nice to be able to play against your friends.

### Add actual IA

Currently, IA makes it's turns randomly, it would be nice to add a more competitive IA. An idea would be to shot adjacent cells after hitting a ship. 
