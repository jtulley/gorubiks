package rubik

import (
	"fmt"
)

const (
	RED    = iota //`0`
	ORANGE        //`1`
	BLUE          //`2`
	GREEN         //`3`
	WHITE         //`4`
	YELLOW        //`5`
)

const (
	FRONT = iota
	BACK
	RIGHT
	LEFT
	TOP
	BOTTOM
)

/*
type Corner struct {
    color1, color2, color3 int
    orientation
}

type MiddleEdge struct {
}

type Middle struct {
}
*/

type Cube struct {
	//pieces [3][9]int
	front  [9]int
	back   [9]int
	right  [9]int
	left   [9]int
	top    [9]int
	bottom [9]int
}

/*
func (c *Cube) defaultInitializeCube() {
    pieces[0][0] = RED | WHITE | GREEN
    pieces[0][1] = RED | WHITE
    pieces[0][2] = RED | WHITE | BLUE
    pieces[0][3] = RED | GREEN
    pieces[0][4] = RED
    pieces[0][5] = RED | BLUE
    pieces[0][6] = RED | YELLOW | GREEN
    pieces[0][7] = RED | YELLOW
    pieces[0][8] = RED | YELLOW | BLUE

    pieces[1][0] = WHITE | GREEN
    pieces[1][1] = RED | WHITE
    pieces[1][2] = RED | WHITE | BLUE
    pieces[1][3] = RED | GREEN
    pieces[1][4] = RED
    pieces[1][5] = RED | BLUE
    pieces[1][6] = RED | YELLOW | GREEN
    pieces[1][7] = RED | YELLOW
    pieces[1][8] = RED | YELLOW | BLUE

    pieces[0][0] = RED | WHITE | GREEN
    pieces[0][1] = RED | WHITE
    pieces[0][2] = RED | WHITE | BLUE
    pieces[0][3] = RED | GREEN
    pieces[0][4] = RED
    pieces[0][5] = RED | BLUE
    pieces[0][6] = RED | YELLOW | GREEN
    pieces[0][7] = RED | YELLOW
    pieces[0][8] = RED | YELLOW | BLUE
}
*/

func (c *Cube) initializeSide(side *[9]int, color int) {
	for i := range side {
		side[i] = color
	}
}

func (c *Cube) DumpFace(sideNumber int) {
	var side *[9]int
	switch sideNumber {
	case FRONT:
		side = &c.front
	case BACK:
		side = &c.back
	case RIGHT:
		side = &c.right
	case LEFT:
		side = &c.left
	case TOP:
		side = &c.top
	case BOTTOM:
		side = &c.bottom
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("%d: %d\n", i, side[i])
	}
}

func (c *Cube) TurnFaceForward(side int) {
	c.turnFace(side, true)
}

func (c *Cube) TurnFaceBack(side int) {
	c.turnFace(side, false)
}

func (c *Cube) turnFace(side int, cw bool) {
	var vFront, vRight, vLeft, vTop, vBottom *[9]int

	vRight = &c.right
	vLeft = &c.left
	vTop = &c.top
	vBottom = &c.bottom
	rES := []int{0, 3, 6}
	lES := []int{8, 5, 2}
	tES := []int{6, 7, 8}
	bES := []int{2, 1, 0}

	rotateCW := cw

	switch side {
	case FRONT:
		vFront = &c.front
	case BACK:
		vFront = &c.back
		vRight = &c.left
		vLeft = &c.right
		//rotateCW = !rotateCW
		rES[0] = 0
		rES[1] = 3
		rES[2] = 6
		lES[0] = 8
		lES[1] = 5
		lES[2] = 2
		tES[0] = 2
		tES[1] = 1
		tES[2] = 0
		bES[0] = 6
		bES[1] = 7
		bES[2] = 8
	case RIGHT:
		vFront = &c.right
		vLeft = &c.front

		vRight = &c.back
		rES[0] = 6
		rES[1] = 3
		rES[2] = 0
		lES[0] = 2
		lES[1] = 5
		lES[2] = 8
		tES[0] = 2
		tES[1] = 5
		tES[2] = 8
		bES[0] = 2
		bES[1] = 5
		bES[2] = 8
	case LEFT:
		vFront = &c.left
		vRight = &c.front
		rES[0] = 0
		rES[1] = 3
		rES[2] = 6

		vLeft = &c.back
		lES[0] = 8
		lES[1] = 5
		lES[2] = 2

		tES[0] = 0
		tES[1] = 3
		tES[2] = 6
		bES[0] = 0
		bES[1] = 3
		bES[2] = 6
	case TOP:
		vFront = &c.top
		vBottom = &c.front
		rES[0] = 0
		rES[1] = 1
		rES[2] = 2
		lES[0] = 0
		lES[1] = 1
		lES[2] = 2

		vTop = &c.back
		tES[0] = 0
		tES[1] = 1
		tES[2] = 2

		bES[0] = 0
		bES[1] = 1
		bES[2] = 2
	case BOTTOM:
		vFront = &c.bottom
		vTop = &c.front
		rES[0] = 6
		rES[1] = 7
		rES[2] = 8
		lES[0] = 6
		lES[1] = 7
		lES[2] = 8
		tES[0] = 6
		tES[1] = 7
		tES[2] = 8

		vBottom = &c.back
		bES[0] = 6
		bES[1] = 7
		bES[2] = 8
	}

	if rotateCW {
		c.rotateFaceCW(vFront)
		c.fixAffectedSidesCW(vRight, vLeft, vTop, vBottom, rES, lES, tES, bES)
	} else {
		c.rotateFaceCCW(vFront)
		c.fixAffectedSidesCCW(vRight, vLeft, vTop, vBottom, rES, lES, tES, bES)
	}
}

func (c *Cube) fixAffectedSidesCCW(vRight *[9]int, vLeft *[9]int, vTop *[9]int, vBottom *[9]int,
	rES []int, lES []int, tES []int, bES []int) {
	tempRight0, tempRight1, tempRight2 := vRight[rES[0]], vRight[rES[1]], vRight[rES[2]]
	vRight[rES[0]] = vBottom[bES[0]]
	vRight[rES[1]] = vBottom[bES[1]]
	vRight[rES[2]] = vBottom[bES[2]]

	vBottom[bES[0]] = vLeft[lES[0]]
	vBottom[bES[1]] = vLeft[lES[1]]
	vBottom[bES[2]] = vLeft[lES[2]]

	vLeft[lES[0]] = vTop[tES[0]]
	vLeft[lES[1]] = vTop[tES[1]]
	vLeft[lES[2]] = vTop[tES[2]]

	vTop[tES[0]] = tempRight0
	vTop[tES[1]] = tempRight1
	vTop[tES[2]] = tempRight2
}

func (c *Cube) fixAffectedSidesCW(vRight *[9]int, vLeft *[9]int, vTop *[9]int, vBottom *[9]int,
	rES []int, lES []int, tES []int, bES []int) {
	tempRight0, tempRight1, tempRight2 := vRight[rES[0]], vRight[rES[1]], vRight[rES[2]]
	vRight[rES[0]] = vTop[tES[0]]
	vRight[rES[1]] = vTop[tES[1]]
	vRight[rES[2]] = vTop[tES[2]]

	vTop[tES[0]] = vLeft[lES[0]]
	vTop[tES[1]] = vLeft[lES[1]]
	vTop[tES[2]] = vLeft[lES[2]]

	vLeft[lES[0]] = vBottom[bES[0]]
	vLeft[lES[1]] = vBottom[bES[1]]
	vLeft[lES[2]] = vBottom[bES[2]]

	vBottom[bES[0]] = tempRight0
	vBottom[bES[1]] = tempRight1
	vBottom[bES[2]] = tempRight2
}

func (c *Cube) rotateFaceCCW(side *[9]int) {
	var newFace [9]int

	newFace[0] = side[2]
	newFace[1] = side[5]
	newFace[2] = side[8]
	newFace[3] = side[1]
	newFace[4] = side[4]
	newFace[5] = side[7]
	newFace[6] = side[0]
	newFace[7] = side[3]
	newFace[8] = side[6]

	for i := 0; i < 9; i++ {
		side[i] = newFace[i]
	}

}

func (c *Cube) rotateFaceCW(side *[9]int) {
	var newFace [9]int

	newFace[0] = side[6]
	newFace[1] = side[3]
	newFace[2] = side[0]
	newFace[3] = side[7]
	newFace[4] = side[4]
	newFace[5] = side[1]
	newFace[6] = side[8]
	newFace[7] = side[5]
	newFace[8] = side[2]

	for i := 0; i < 9; i++ {
		side[i] = newFace[i]
	}

}

func (c *Cube) Front(pieceNumber int) int {
	return c.front[pieceNumber]
}

func (c *Cube) Back(pieceNumber int) int {
	return c.back[pieceNumber]
}

func (c *Cube) Left(pieceNumber int) int {
	return c.left[pieceNumber]
}

func (c *Cube) Right(pieceNumber int) int {
	return c.right[pieceNumber]
}

func (c *Cube) Top(pieceNumber int) int {
	return c.top[pieceNumber]
}

func (c *Cube) Bottom(pieceNumber int) int {
	return c.bottom[pieceNumber]
}

func (c *Cube) isSolved() bool {
	for i := 0; i < 9; i++ {
		if c.Front(i) != RED {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if c.Back(i) != ORANGE {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if c.Right(i) != BLUE {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if c.Left(i) != GREEN {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if c.Top(i) != WHITE {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if c.Bottom(i) != YELLOW {
			return false
		}
	}
	return true

}

func NewSpecialCube() *Cube {
	newCube := new(Cube)
	newCube.initializeSide(&newCube.front, WHITE)
	newCube.initializeSide(&newCube.back, WHITE)
	newCube.initializeSide(&newCube.right, WHITE)
	newCube.initializeSide(&newCube.left, WHITE)
	newCube.initializeSide(&newCube.top, WHITE)
	newCube.initializeSide(&newCube.bottom, WHITE)

	newCube.bottom[6] = WHITE
	newCube.bottom[3] = ORANGE
	newCube.bottom[0] = RED

	newCube.bottom[8] = RED
	newCube.bottom[5] = ORANGE
	newCube.bottom[2] = WHITE

	newCube.top[6] = WHITE
	newCube.top[3] = RED
	newCube.top[0] = ORANGE

	newCube.top[8] = ORANGE
	newCube.top[5] = RED
	newCube.top[2] = WHITE

	newCube.back[0] = BLUE
	newCube.back[3] = GREEN
	newCube.back[6] = YELLOW

	newCube.back[2] = GREEN
	newCube.back[5] = YELLOW
	newCube.back[8] = BLUE

	newCube.left[0] = YELLOW
	newCube.left[1] = GREEN
	newCube.left[2] = ORANGE
	newCube.left[3] = RED
	newCube.left[4] = WHITE
	newCube.left[5] = BLUE
	newCube.left[6] = ORANGE
	newCube.left[7] = WHITE
	newCube.left[8] = RED

	newCube.right[0] = YELLOW
	newCube.right[1] = GREEN
	newCube.right[2] = ORANGE
	newCube.right[3] = RED
	newCube.right[4] = WHITE
	newCube.right[5] = BLUE
	newCube.right[6] = ORANGE
	newCube.right[7] = WHITE
	newCube.right[8] = RED

	newCube.front[0] = BLUE
	newCube.front[3] = GREEN
	newCube.front[6] = YELLOW

	newCube.front[2] = GREEN
	newCube.front[5] = YELLOW
	newCube.front[8] = BLUE

	return newCube
}

func NewCube() *Cube {
	newCube := new(Cube)
	newCube.initializeSide(&newCube.front, RED)
	newCube.initializeSide(&newCube.back, ORANGE)
	newCube.initializeSide(&newCube.right, BLUE)
	newCube.initializeSide(&newCube.left, GREEN)
	newCube.initializeSide(&newCube.top, WHITE)
	newCube.initializeSide(&newCube.bottom, YELLOW)

	return newCube
}
