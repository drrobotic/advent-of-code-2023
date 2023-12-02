package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Game struct {
	id       int
	maxRed   int
	maxBlue  int
	maxGreen int
}

func (g *Game) power() int {
	return g.maxRed * g.maxBlue * g.maxGreen
}

var splitColorRegex = regexp.MustCompile("[\\;\\,]+")

func parseGame(line string) (Game, error) {
	max_red := 0
	max_green := 0
	max_blue := 0
	line_parts := strings.Split(line, ":")

	var gameid int
	_, err := fmt.Sscanf(line_parts[0], "Game %d", &gameid)
	if err != nil {
		return Game{}, errors.New("no game id found")
	}

	color_parts := splitColorRegex.Split(line_parts[1], -1)
	for _, color_part := range color_parts {
		var num int
		var color string
		_, err := fmt.Sscanf(color_part, "%d %s", &num, &color)
		if err != nil {
			fmt.Printf("couldnt parse color part %s\n", color_part)
			continue
		}

		if color == "red" && num > max_red {
			max_red = num
		} else if color == "green" && num > max_green {
			max_green = num
		} else if color == "blue" && num > max_blue {
			max_blue = num
		}
	}

	return Game{gameid, max_red, max_blue, max_green}, nil
}

func main() {
	myfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could open file: %s", err)
	}
	defer myfile.Close()

	game_ids := []int{}
	game_powers := []int{}
	reader := bufio.NewScanner(myfile)

	for reader.Scan() {
		line := reader.Text()

		if game, err := parseGame(line); err == nil {
			if game.maxRed <= 12 && game.maxBlue <= 14 && game.maxGreen <= 13 {
				game_ids = append(game_ids, game.id)
			}
			game_powers = append(game_powers, game.power())
		}
	}

	sum_1 := 0
	for _, id := range game_ids {
		sum_1 += id
	}

	sum_2 := 0
	for _, p := range game_powers {
		sum_2 += p
	}

	fmt.Println("Day 2")
	fmt.Printf("part1: %d\n", sum_1)
	fmt.Printf("part2: %d\n", sum_2)
}
