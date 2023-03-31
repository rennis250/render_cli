package main

// Adapted from this tutorial: http://www.learnopengl.com/#!Lighting/Colors

import (
	"flag"
	"image/png"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/rennis250/renderer"
)

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

var fname, imgname string

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	flag.StringVar(&fname, "filename", "scene.json", "Filename for rendered image")
	flag.StringVar(&imgname, "output", "scene.png", "Filename for rendered image")
	flag.Parse()

	r, err := renderer.NewRendererFromJSON(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	window, err := glfw.CreateWindow(int(r.WindowWidth), int(r.WindowHeight), "RobRender", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow (go function bindings)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(keyCallback)

	gl.Disable(gl.DEPTH_TEST)

	err = r.AssembleScene()
	if err != nil {
		panic(err)
	}

	for !window.ShouldClose() {
		glfw.PollEvents()

		ss, err := r.Render()
		if err != nil {
			log.Fatalln(err)
		}

		f, err := os.Create(imgname)
		if err != nil {
			log.Fatalln(err)
		}
		err = png.Encode(f, ss)
		if err != nil {
			log.Fatalln(err)
		}
		f.Close()

		window.SwapBuffers()

		break
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// When a user presses the escape key, we set the WindowShouldClose property to true,
	// which closes the application
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
