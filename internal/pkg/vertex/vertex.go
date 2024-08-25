package vertex

import "github.com/go-gl/gl/v4.1-core/gl"

func NewVertexArray(points []float32) uint32 {
	var VertexBuffer uint32
	gl.GenBuffers(1, &VertexBuffer)
	gl.GenBuffers(1, &VertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, VertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var VertexArray uint32
	gl.GenVertexArrays(1, &VertexArray)
	gl.BindVertexArray(VertexArray)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, VertexBuffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return VertexArray
}
