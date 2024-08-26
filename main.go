package main

import (
	"GameOfLife/internal/app/cells"
	"GameOfLife/internal/pkg/window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"time"
)

func main() {

	win := window.NewWindow(500, 500, "Game Of Life")

	var cell [][]*cells.Cell
	err := win.Start(
		func() {
			cell = cells.MakeCells(10, 10)
		},
		func(program uint32, window *glfw.Window) {
			t := time.Now()
			for x := range cell {
				for _, c := range cell[x] {
					c.CheckState(cell)
				}
			}
			draw(cell, window, program)
			time.Sleep(time.Second/time.Duration(2) - time.Since(t))
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func draw(cell [][]*cells.Cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, ca := range cell {
		for _, c := range ca {
			c.Draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
