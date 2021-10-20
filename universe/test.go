package universe

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	width = 20		// index: 1-20
	height = 10		// index: 1-10

	Rate = 25
	MinNeighbourNum = 3
)

// 二维数组
type Universe [][]bool

type Pos struct{
	h int
	w int
}

func newUniverse() Universe {
	var cells Universe
	for i := 0; i < height; i++ {
		cells = append(cells, make([]bool, width))
	}
	return cells
}

func (un Universe)show() {
	for i := 0; i < height; i++ {
		for _, v := range(un[i]) {
			fmt.Printf("%t ", v)
		}
		fmt.Printf("\n")
	}
}

func (un Universe)seed() {
	num := math.Ceil(width * height * Rate / 100)
	if num <= 0 {
		return
	}
	for {
		w := rand.Intn(width)
		h := rand.Intn(height)
		if h > 0 && w > 0 && !un[h][w] {
			un[h][w] = true
			num -= 1
			if num <= 0 {
				break
			}
		}
	}
}

// 九宫格
func getNeighbour(h, w int) []Pos {
	sourceH, sourceW := h-1, w-1
	posList := make([]Pos, MinNeighbourNum) // 至少有3个成员
	index := 0

	for i := 0; i < 3; i++ {
		trueH := sourceH + i
		if trueH < 0 {
			continue
		}
		for j := 0; j < 3; j++ {
			trueW := sourceW + j
			if trueH >= 0 &&
			trueW >= 0 &&
			trueH < height &&
			trueW < width &&
			(trueH != h || trueW != w) {
				if index < MinNeighbourNum {
					*&posList[index].h = trueH
					*&posList[index].w = trueW
					index++
				}else{
					posList = append(posList, Pos{h:trueH, w:trueW})
				}
			}
		}
	}
	return posList
}

func (un Universe)isAlive(h, w int) bool {
	posList := getNeighbour(h, w)
	if len(posList) == 0 {
		return false
	}

	aliveCount, deadCount := 0, 0
	for _, v := range(posList){
		cell := un[v.h][v.w]
		if cell {
			aliveCount++
		}else{
			deadCount++
		}
	}

	isAlive := false
	cell := un[h][w]
	if cell {
		// 存活细胞，邻近存活细胞少于2个或多于3个时，死亡
		if !(aliveCount < 2 || aliveCount > 3) {
			isAlive = true
		}
	}else{
		// 死亡细胞，邻近存活细胞3个时，存活
		if aliveCount == 3 {
			isAlive = true
		}
	}
	return isAlive
}

func (un Universe)round() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			isAlive := un.isAlive(i, j)
			un[i][j] = isAlive
		}
	}
}

func Test(){
	fmt.Println("-------------------- submodule universe --------------------")

	cells := newUniverse()
	cells.seed()
	cells.show()
	
	index := 1
	for {
		fmt.Println("universe round ", index)
		cells.round()
		cells.show()
		index++
		if index > 10 {
			break
		}
	}
}
