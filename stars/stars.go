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
			if mat[j][i] { s = '*' }
			fmt.Print(string(s))
		}
		fmt.Print("\n")
	}
}

func inArray(i int, a []int) bool {
	for _, e := range a { if i == e { return true }	}

	return false
}

func solve(line int, pos int, s Stars) matrix {
	mat := make(matrix)
	mat[line] = make(map[int] bool)
	mat[line][(pos - 1)] = true
	a := make([]int, s.x)
	a[line] = pos
	for j := 0; j < s.y; j++ {
		if j == line { continue }
		mat[j] = make(map[int] bool)
		for i := 1; i <= s.y; i++ {
			if inArray(i, a) { continue }
			mat[j][(i - 1)] = true
			a[j] = i
			break
		}
	}

	return mat
}

func inSolutions(mat matrix, mats []matrix) bool {
	var same bool
	for _, m := range mats {
		same = true
		if len(m) == 0 { continue }
		for i, _ := range mat {
			if mat[i][i] != m[i][i] {
				same = false
				break
			}
		}

		if same { return true }
	}

        return false
}

func isValid(m matrix, s Stars) bool {
	g := make(map[string] bool, len(s.groups))
	for j := 0; j < s.y; j++ {
		for i := 0; i < s.x; i++ {
                        if m[j][i] {
				if g[s.input[j][i]] {
					return false
				}
				g[s.input[j][i]] = true
			}
		}
	}

	return true
}

func main() {
	s := parse()
	mat := make([]matrix, (s.x * s.y))

	num := 0
	for j := 0; j < s.y; j++ {
		for i := 1; i <= s.x; i++ {
			m := solve(j, i, s)
			if inSolutions(m, mat) {
				continue
			}
			mat[num] = m
			num++
		}
	}

	for _, m := range mat {
		if len(m) == 0 { continue }
		if isValid(m, s) {
			print(s.x, s.y, m)
		}
	}
}