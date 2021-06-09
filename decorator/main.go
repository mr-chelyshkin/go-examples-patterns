package main

import "fmt"

//Shape is Component
type Shape interface {
	ShowInfo()
}

// -- >

//Square is ConcreteComponent
type Square struct{}

//ShowInfo is Operation()
func (s Square) ShowInfo() {
	fmt.Println("square")
}

// -- >

//ShapeDecorator is Decorator
type ShapeDecorator struct {
	Shape Shape
}

//ShowInfo is Operation()
func (sd ShapeDecorator) ShowInfo() {
	sd.Shape.ShowInfo()
}

// -- >

//ColorShape is ConcreteDecorator
type ColorShape struct {
	ShapeDecorator
	color string
}

//ShowInfo is Operation()
func (cs ColorShape) ShowInfo() {
	fmt.Println(cs.color + " ")
	cs.Shape.ShowInfo()
}

// -- >

//ShadowShape is ConcreteDecorator
type ShadowShape struct {
	ShapeDecorator
}

//ShowInfo is Operation()
func (ss ShadowShape) ShowInfo() {
	ss.Shape.ShowInfo()
	fmt.Println(" with shadow")
}

// -- >
func main() {
	square := Square{}
	square.ShowInfo()
	//printed: square

	colorShape := ColorShape{ShapeDecorator{square}, "red"}
	colorShape.ShowInfo()
	//printed: red square

	shadowShape := ShadowShape{ShapeDecorator{colorShape}}
	shadowShape.ShowInfo()
	//printed: red square with shadow
}
