module github.com/rennis250/render_cli

go 1.19

replace github.com/rennis250/renderer => ../renderer

require (
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6
	github.com/go-gl/glfw v0.0.0-20221017161538-93cebf72946b
	github.com/rennis250/renderer v0.0.0-00010101000000-000000000000
)

require github.com/cstegel/opengl-samples-golang v0.0.0-20180607031111-1ed56cc6485a // indirect
