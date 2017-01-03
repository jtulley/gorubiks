package rubik

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialCubeLayout(t *testing.T) {
	theCube := NewCube()
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, BLUE, theCube.Right(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, GREEN, theCube.Left(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, WHITE, theCube.Top(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, YELLOW, theCube.Bottom(i))
	}
	assert.Equal(t, true, theCube.isSolved())
}

func TestTurnFrontCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(FRONT)
	//theCube.DumpFace(TOP)
	//theCube.DumpFace(BOTTOM)
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, WHITE, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, BLUE, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, GREEN, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnFrontCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(FRONT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, WHITE, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, GREEN, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, BLUE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnFrontCCWTwice(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(FRONT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, GREEN, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, BLUE, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, YELLOW, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, WHITE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnFrontCCWFourTimesBackToStart(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(FRONT)
	assert.Equal(t, true, theCube.isSolved())
}

func TestTurnFrontCWFourTimesBackToStart(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(FRONT)
	theCube.TurnFaceForward(FRONT)
	theCube.TurnFaceForward(FRONT)
	theCube.TurnFaceForward(FRONT)
	assert.Equal(t, true, theCube.isSolved())
}

func TestTurnRightCWFourTimesBackToStart(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(RIGHT)
	theCube.TurnFaceForward(RIGHT)
	theCube.TurnFaceForward(RIGHT)
	theCube.TurnFaceForward(RIGHT)
	assert.Equal(t, true, theCube.isSolved())
}

func TestTurnLeftCWFourTimesBackToStart(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(LEFT)
	theCube.TurnFaceForward(LEFT)
	theCube.TurnFaceForward(LEFT)
	theCube.TurnFaceForward(LEFT)
	assert.Equal(t, true, theCube.isSolved())
}

func TestTurnBackCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(BACK)
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, WHITE, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, GREEN, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, BLUE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnBackCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(BACK)
	for i := 0; i < 9; i++ {
		assert.Equal(t, ORANGE, theCube.Back(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, RED, theCube.Front(i))
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, WHITE, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, BLUE, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, GREEN, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnRightCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(RIGHT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, BLUE, theCube.Right(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, GREEN, theCube.Left(i))
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, WHITE, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, ORANGE, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, RED, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnRightCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(RIGHT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, BLUE, theCube.Right(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, GREEN, theCube.Left(i))
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, WHITE, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, RED, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, ORANGE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnLeftCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(LEFT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, BLUE, theCube.Right(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, GREEN, theCube.Left(i))
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, WHITE, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, RED, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, ORANGE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnLeftCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(LEFT)
	for i := 0; i < 9; i++ {
		assert.Equal(t, BLUE, theCube.Right(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, GREEN, theCube.Left(i))
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, WHITE, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i+1)%3 == 0 {
			assert.Equal(t, YELLOW, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, ORANGE, theCube.Top(i))
		} else {
			assert.Equal(t, WHITE, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			assert.Equal(t, RED, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnBottomCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(BOTTOM)
	for i := 0; i < 9; i++ {
		assert.Equal(t, YELLOW, theCube.Bottom(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, WHITE, theCube.Top(i))
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, BLUE, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, GREEN, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, ORANGE, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, RED, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnBottomCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(BOTTOM)
	for i := 0; i < 9; i++ {
		assert.Equal(t, YELLOW, theCube.Bottom(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, WHITE, theCube.Top(i))
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, GREEN, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, BLUE, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, RED, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i > 5 {
			assert.Equal(t, ORANGE, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnTopCCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(TOP)
	for i := 0; i < 9; i++ {
		assert.Equal(t, YELLOW, theCube.Bottom(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, WHITE, theCube.Top(i))
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, GREEN, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, BLUE, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, RED, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, ORANGE, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestTurnTopCW(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(TOP)
	for i := 0; i < 9; i++ {
		assert.Equal(t, YELLOW, theCube.Bottom(i))
	}
	for i := 0; i < 9; i++ {
		assert.Equal(t, WHITE, theCube.Top(i))
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, BLUE, theCube.Front(i))
		} else {
			assert.Equal(t, RED, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, GREEN, theCube.Back(i))
		} else {
			assert.Equal(t, ORANGE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, ORANGE, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, RED, theCube.Left(i))
		} else {
			assert.Equal(t, GREEN, theCube.Left(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestFrontRight(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(FRONT)
	theCube.TurnFaceForward(RIGHT)
	for i := 0; i < 9; i++ {
		if i%3 < 2 {
			assert.Equal(t, RED, theCube.Front(i))
		} else if i == 2 {
			assert.Equal(t, BLUE, theCube.Front(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Front(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i%3 > 0 {
			assert.Equal(t, ORANGE, theCube.Back(i))
		} else if i == 0 {
			assert.Equal(t, GREEN, theCube.Back(i))
		} else {
			assert.Equal(t, WHITE, theCube.Back(i))
		}
	}
	for i := 0; i < 9; i++ {
		if i < 3 {
			assert.Equal(t, WHITE, theCube.Right(i))
		} else {
			assert.Equal(t, BLUE, theCube.Right(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i % 3) < 2 {
			assert.Equal(t, GREEN, theCube.Left(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Left(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i % 3) > 1 {
			assert.Equal(t, RED, theCube.Top(i))
		} else if i < 6 {
			assert.Equal(t, WHITE, theCube.Top(i))
		} else {
			assert.Equal(t, GREEN, theCube.Top(i))
		}
	}
	for i := 0; i < 9; i++ {
		if (i % 3) > 1 {
			assert.Equal(t, ORANGE, theCube.Bottom(i))
		} else if i < 3 {
			assert.Equal(t, BLUE, theCube.Bottom(i))
		} else {
			assert.Equal(t, YELLOW, theCube.Bottom(i))
		}
	}
	assert.Equal(t, false, theCube.isSolved())
}

func TestBackForward(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceForward(FRONT)
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceForward(BACK)
	theCube.TurnFaceBack(BACK)
	theCube.TurnFaceForward(RIGHT)
	theCube.TurnFaceBack(RIGHT)
	theCube.TurnFaceForward(LEFT)
	theCube.TurnFaceBack(LEFT)
	theCube.TurnFaceForward(TOP)
	theCube.TurnFaceBack(TOP)
	theCube.TurnFaceForward(BOTTOM)
	theCube.TurnFaceBack(BOTTOM)

	assert.Equal(t, true, theCube.isSolved())
}

func DumpCube(theCube *Cube) {
	fmt.Println("----------FRONT---------")
	theCube.DumpFace(FRONT)
	fmt.Println("----------RIGHT---------")
	theCube.DumpFace(RIGHT)
	fmt.Println("----------BACK---------")
	theCube.DumpFace(BACK)
	fmt.Println("----------LEFT---------")
	theCube.DumpFace(LEFT)
	fmt.Println("----------BOTTOM---------")
	theCube.DumpFace(BOTTOM)
	fmt.Println("----------TOP---------")
	theCube.DumpFace(TOP)
}

//F'R'T'F'L'Ba'Bo'Ba'
func TestGoodScramble(t *testing.T) {
	theCube := NewCube()
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(RIGHT)
	theCube.TurnFaceBack(TOP)
	theCube.TurnFaceBack(FRONT)
	theCube.TurnFaceBack(LEFT)
	theCube.TurnFaceBack(BACK)
	theCube.TurnFaceBack(BOTTOM)
	theCube.TurnFaceBack(BACK)

	assert.Equal(t, ORANGE, theCube.Front(0))
	assert.Equal(t, WHITE, theCube.Front(1))
	assert.Equal(t, BLUE, theCube.Front(2))
	assert.Equal(t, YELLOW, theCube.Front(3))
	assert.Equal(t, RED, theCube.Front(4))
	assert.Equal(t, RED, theCube.Front(5))
	assert.Equal(t, GREEN, theCube.Front(6))
	assert.Equal(t, YELLOW, theCube.Front(7))
	assert.Equal(t, ORANGE, theCube.Front(8))

	assert.Equal(t, ORANGE, theCube.Back(0))
	assert.Equal(t, ORANGE, theCube.Back(1))
	assert.Equal(t, GREEN, theCube.Back(2))
	assert.Equal(t, WHITE, theCube.Back(3))
	assert.Equal(t, ORANGE, theCube.Back(4))
	assert.Equal(t, GREEN, theCube.Back(5))
	assert.Equal(t, RED, theCube.Back(6))
	assert.Equal(t, BLUE, theCube.Back(7))
	assert.Equal(t, RED, theCube.Back(8))

	assert.Equal(t, RED, theCube.Right(0))
	assert.Equal(t, RED, theCube.Right(1))
	assert.Equal(t, YELLOW, theCube.Right(2))
	assert.Equal(t, GREEN, theCube.Right(3))
	assert.Equal(t, BLUE, theCube.Right(4))
	assert.Equal(t, ORANGE, theCube.Right(5))
	assert.Equal(t, BLUE, theCube.Right(6))
	assert.Equal(t, YELLOW, theCube.Right(7))
	assert.Equal(t, BLUE, theCube.Right(8))

	assert.Equal(t, ORANGE, theCube.Left(0))
	assert.Equal(t, WHITE, theCube.Left(1))
	assert.Equal(t, WHITE, theCube.Left(2))
	assert.Equal(t, ORANGE, theCube.Left(3))
	assert.Equal(t, GREEN, theCube.Left(4))
	assert.Equal(t, GREEN, theCube.Left(5))
	assert.Equal(t, YELLOW, theCube.Left(6))
	assert.Equal(t, RED, theCube.Left(7))
	assert.Equal(t, RED, theCube.Left(8))

	assert.Equal(t, YELLOW, theCube.Top(0))
	assert.Equal(t, YELLOW, theCube.Top(1))
	assert.Equal(t, BLUE, theCube.Top(2))
	assert.Equal(t, GREEN, theCube.Top(3))
	assert.Equal(t, WHITE, theCube.Top(4))
	assert.Equal(t, BLUE, theCube.Top(5))
	assert.Equal(t, GREEN, theCube.Top(6))
	assert.Equal(t, BLUE, theCube.Top(7))
	assert.Equal(t, YELLOW, theCube.Top(8))

	assert.Equal(t, WHITE, theCube.Bottom(0))
	assert.Equal(t, RED, theCube.Bottom(1))
	assert.Equal(t, WHITE, theCube.Bottom(2))
	assert.Equal(t, WHITE, theCube.Bottom(3))
	assert.Equal(t, YELLOW, theCube.Bottom(4))
	assert.Equal(t, BLUE, theCube.Bottom(5))
	assert.Equal(t, GREEN, theCube.Bottom(6))
	assert.Equal(t, ORANGE, theCube.Bottom(7))
	assert.Equal(t, WHITE, theCube.Bottom(8))

	assert.Equal(t, false, theCube.isSolved())
}

func TestRobertsTheory(t *testing.T) {
	theCube := NewCube()
	for i := 0; i < 10000; {
		theCube.TurnFaceForward(RIGHT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(TOP)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(LEFT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(FRONT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceForward(BACK)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(BOTTOM)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceForward(RIGHT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(TOP)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceBack(RIGHT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}

		theCube.TurnFaceForward(LEFT)
		i++
		if checkCorrectness(theCube, i) {
			break
		}
	}
}

/*
func TestRobertsTheory_2(t *testing.T) {
	theCube := NewCube()
    for i := 0; i < 10000; {
	    theCube.TurnFaceForward(TOP);
		i++
		if (checkCorrectness(theCube, i)) {
			break;
		}
	    theCube.TurnFaceBack(RIGHT);
		i++
		if (checkCorrectness(theCube, i)) {
			break;
		}
	    theCube.TurnFaceBack(BOTTOM);
		i++
		if (checkCorrectness(theCube, i)) {
			break;
		}
    }
}
*/

func checkCorrectness(c *Cube, count int) bool {
	if c.isSolved() {
		fmt.Println("Solved at count: ", count)
		return true
	} else {
		if count%1000 == 0 {
			fmt.Printf("%d turns\n", count)
		}
		return false
	}
}
