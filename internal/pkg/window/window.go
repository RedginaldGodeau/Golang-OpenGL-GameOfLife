package window

import (
	"GameOfLife/internal/pkg/opengl"
	"errors"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type Window struct {
	Win           *glfw.Window
	Title         string
	Width, Height int
}

func NewWindow(width, height int, title string) *Window {
	return &Window{
		Win:    nil,
		Width:  width,
		Height: height,
		Title:  title,
	}
}

func (w *Window) initWindow() error {
	if err := glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w.Width, w.Height, w.Title, nil, nil)
	if err != nil {
		return err
	}
	if window == nil {
		return errors.New("CreateWindow failed window is nil")
	}
	w.Win = window
	w.Win.MakeContextCurrent()

	return nil
}

func (w *Window) Start(initHandler func(), updateHandler LoopFunc) error {
	runtime.LockOSThread()

	err := w.initWindow()
	defer glfw.Terminate()
	if err != nil {
		return err
	}

	program, err := opengl.InitOpenGL()
	if err != nil {
		return err
	}

	if w.Win == nil {
		return errors.New("window is nil")
	}

	initHandler()
	for !w.Win.ShouldClose() {
		updateHandler(program, w.Win)
	}

	return nil
}
