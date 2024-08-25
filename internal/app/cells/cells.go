package cells

import (
	"GameOfLife/internal/pkg/vertex"
	"github.com/go-gl/gl/v4.1-core/gl"
	"math/rand"
	"time"
)

func newCell(x, y, row, col int) *Cell {
	points := make([]float32, len(vertex.SquareVertex), len(vertex.SquareVertex))
	copy(points, vertex.SquareVertex)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(col)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(row)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &Cell{
		Drawable: vertex.NewVertexArray(points),
		X:        x,
		Y:        y,
	}
}

func MakeCells(row, column int) [][]*Cell {
	rand.Seed(time.Now().UnixNano())

	cells := make([][]*Cell, row, row)
	for x := 0; x < row; x++ {
		for y := 0; y < column; y++ {
			c := newCell(x, y, row, column)
			c.Alive = rand.Float64() < .15
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}
	return cells
}

func (c *Cell) Draw() {
	if !c.Alive {
		return
	}

	gl.BindVertexArray(c.Drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertex.SquareVertex)/3))
}

func (c *Cell) CheckState(cells [][]*Cell) {
	c.Alive = c.AliveNext
	c.AliveNext = c.Alive

	liveCount := c.liveNeighbors(cells)
	if c.Alive {
		if liveCount < 2 {
			c.AliveNext = false
		}
		if liveCount == 2 || liveCount == 3 {
			c.AliveNext = true
		}
		if liveCount > 3 {
			c.AliveNext = false
		}
	} else {
		if liveCount == 3 {
			c.AliveNext = true
		}
	}
}

func (c *Cell) liveNeighbors(cells [][]*Cell) int {
	var liveCount int
	aliveCheck := func(x, y int) {
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].Alive {
			liveCount++
		}
	}

	aliveCheck(c.X-1, c.Y)   // To the left
	aliveCheck(c.X+1, c.Y)   // To the right
	aliveCheck(c.X, c.Y+1)   // up
	aliveCheck(c.X, c.Y-1)   // down
	aliveCheck(c.X-1, c.Y+1) // top-left
	aliveCheck(c.X+1, c.Y+1) // top-right
	aliveCheck(c.X-1, c.Y-1) // bottom-left
	aliveCheck(c.X+1, c.Y-1) // bottom-right

	return liveCount
}
