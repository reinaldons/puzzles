package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stars struct {
	num int
	x int
	y int
	groups map[string] bool
	corners int
	lines int
	columns int
	input map[int] map[int] string
}

func parse() (s Stars) {
	buf := bufio.NewReader(os.Stdin)
	line, _, err := buf.ReadLine()
	s.num, _ = strconv.Atoi(string(line[0]))
	s.x, _ = strconv.Atoi(string(line[2])) 
	s.y, _ = strconv.Atoi(string(line[4]))

	s.groups = make(map[string] bool)
	s.input = make(map[int] map[int] string)
	l := 0
        for err != os.EOF {
                fmt.Println(string(line))
                line, _, err = buf.ReadLine()
		s.input[l] = make(map[int] string)
		for pos, letter := range line {
			s.groups[string(letter)] = true
			s.input[l][pos] = string(letter)
                }
		l++
        }

	s.corners = (s.x - 1) * (s.y -1)
	fmt.Println("corners: ", s.corners)
	s.lines = s.x * s.y + s.corners
	fmt.Println("lines: ", s.lines)
	s.columns = s.x + s.y + len(s.groups) + s.corners
	fmt.Println("columns: ", s.columns)

	return
}

type matrix map[int] map[int] bool

func print(x int, y int, mat matrix) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			s := '.'
			if mat[j][i] {
				s = '*'
			}
			fmt.Print(string(s))
		}
		fmt.Print("\n")
	}
}

func main() {
	s := parse()
	//fmt.Println("--> ", s)

	mat := make(matrix)

	for j := 0; j < s.y; j++ {
		for i := 0; i < s.x; i++ {
			line := j * s.x + i;
			mat[line] = make(map[int] bool)
			mat[line][i] = true
			mat[line][s.y + j] = true
			mat[line][s.x + s.y + (int([]byte(s.input[j][i])[0]) - 49)] = true
		}
	}

	//fmt.Println(mat)
	print(s.x, s.y, mat)
}