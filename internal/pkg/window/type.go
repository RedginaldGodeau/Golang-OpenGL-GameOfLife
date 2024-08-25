package window

import "github.com/go-gl/glfw/v3.2/glfw"

type LoopFunc func(program uint32, window *glfw.Window)
