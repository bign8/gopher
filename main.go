package main

import (
	"math"
	"math/cmplx"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

// TODO: turn into proper object orientation
var (
	doc    = js.Global.Get("document")
	eye    = doc.Call("getElementById", "eye")
	canvas = doc.Call("createElement", "canvas")
	ctx    = canvas.Call("getContext", "2d")
	mouse  = complex(25, 57) // relative to canvas
	left   = complex(81, 52)
	right  = complex(161, 49)
	shift  = getCmpl(eye, "height", "width") / 2 // offset drawing to senter pupil
)

func drawEye(origin complex128) {
	// thanks https://stackoverflow.com/a/24639607 (with complex numbers)
	radius, angle := cmplx.Polar(mouse - origin)
	radius = math.Min(radius, 17)
	pupil := cmplx.Rect(radius, angle) + origin - shift

	// clear eyeball
	ctx.Call("beginPath")
	ctx.Call("arc", real(origin), imag(origin), 29, 0, 2*math.Pi)
	ctx.Call("fill")
	ctx.Call("drawImage", eye, real(pupil), imag(pupil)) // draw pupil
}

func draw() {
	drawEye(left)
	drawEye(right)
}

func getCmpl(obj *js.Object, x, y string) complex128 {
	return complex(obj.Get(x).Float(), obj.Get(y).Float())
}

func mouseMove(e *js.Object) {
	// thanks https://stackoverflow.com/a/17130415 (with complex numbers)
	rect := canvas.Call("getBoundingClientRect") // TODO: only on doc resizes
	mouse = getCmpl(e, "clientX", "clientY") - getCmpl(rect, "x", "y")
	js.Global.Call("requestAnimationFrame", draw)
}

func main() {
	js.Global.Set("onload", start)
	// TODO: add image loaded listener
	// TODO: dual promisify loading of image and dom
}

func start() {
	doc.Call("getElementById", "frame").Call("appendChild", canvas)
	canvas.Set("style", "position:absolute;top:0;left:0;height:350px;width:250px;")
	canvas.Set("height", strconv.Itoa(350)) // badness example
	canvas.Set("width", strconv.Itoa(250))  // badness example

	ratio := float64(1) //js.Global.Get("window").Get("devicePixelRatio").Float()
	println("ratio", ratio)
	canvas.Set("height", strconv.FormatFloat(350*ratio, 'f', -1, 64)) // badness example
	canvas.Set("width", strconv.FormatFloat(250*ratio, 'f', -1, 64))  // badness example

	// log.Printf("device pixel ratio: %f", ratio)
	// c.SetAttribute("width", gorenderer.ConvertFloat(float64(width)*ratio))
	// c.SetAttribute("height", gorenderer.ConvertFloat(float64(height)*ratio))
	// c.SetStyle("height", gorenderer.ConvertInt(height)+"px") //convert to pixels
	// c.SetStyle("width", gorenderer.ConvertInt(width)+"px")   //convert to pixels

	doc.Call("addEventListener", "mousemove", mouseMove)
	ctx.Set("fillStyle", "#FFFFFF")
	draw()
}
