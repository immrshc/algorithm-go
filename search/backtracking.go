package search

import (
	"fmt"
	"strings"
)

type board struct {
	N                 int
	rows              []int  // それぞれの行ごとの配置されている列番号
	columns           []bool // それぞれの列が占有されているかどうか
	upDiagonalLines   []bool // それぞれの斜め右上の線上が占有されているかどうか
	downDiagonalLines []bool // それぞれの斜め右上の線上が占有されているかどうか
}

func NewBoard(n int) *board {
	if n <= 0 {
		return nil
	}
	rows := make([]int, n)
	for i := 0; i < len(rows); i++ {
		rows[i] = -1
	}
	return &board{
		N:                 n,
		rows:              rows,
		columns:           make([]bool, n),
		upDiagonalLines:   make([]bool, 2*n-1),
		downDiagonalLines: make([]bool, 2*n-1),
	}
}

func (b board) CanPutQueen(r, c int) bool {
	if r >= b.N || c >= b.N {
		return false
	}
	return !b.columns[c] && !b.downDiagonalLines[r-c+b.N-1] && !b.upDiagonalLines[r+c]
}

func (b *board) PutQueen(r, c int) {
	b.rows[r] = c
	b.columns[c] = true
	b.downDiagonalLines[r-c+b.N-1] = true
	b.upDiagonalLines[r+c] = true
}

func (b *board) RemoveQueen(r, c int) {
	b.rows[r] = -1
	b.columns[c] = false
	b.downDiagonalLines[r-c+b.N-1] = false
	b.upDiagonalLines[r+c] = false
}

func (b board) PrintQueen() {
	line := strings.Repeat("----", b.N) + "-"
	fmt.Println(line)
	for r := 0; r < b.N; r++ {
		for c := 0; c < b.N; c++ {
			if b.rows[r] == c {
				fmt.Print("| Q ")
			} else {
				fmt.Print("|   ")
			}
		}
		fmt.Println("|")
		fmt.Println(line)
	}
}

func Backtracking() {
	b := NewBoard(8)
	backtracking(b, 0)
	b.PrintQueen()
}

func backtracking(b *board, r int) bool {
	for c := 0; c < b.N; c++ {
		if !b.CanPutQueen(r, c) {
			continue
		}
		b.PutQueen(r, c)
		// この行が最後であれば終了を返す
		if r+1 >= b.N {
			return true
		}
		// まだ次の行に配置できるのであれば配置し、無事できれば終了を返す
		if backtracking(b, r+1) {
			return true
		}
		// 次の行のどの列にも配置できなかった場合はこの行と現在の列の組み合わせの配置をやめて、次の列に配置する
		b.RemoveQueen(r, c)
	}
	return false
}
