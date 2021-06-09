package main

import "fmt"

//Graphic is Component
type Graphic interface {
	Draw()
}

// -- >

//Circle is Leaf
type Circle struct{}

//Draw is Operation
func (c Circle) Draw() {
	fmt.Println("Draw circle")
}

// -- >

//Square is Leaf
type Square struct{}

//Draw is Operation
func (s Square) Draw() {
	fmt.Println("Draw square")
}

// -- >

//Image is Composite
type Image struct {
	graphics []Graphic
}

//Add Adds a Leaf to the Composite.
func (i *Image) Add(graphic Graphic) {
	i.graphics = append(i.graphics, graphic)
}

//Draw is Operation
func (i Image) Draw() {
	fmt.Println("Draw image")
	for _, g := range i.graphics {
		g.Draw()
	}
}

// -- >
func main() {
	image := Image{}

	image.Add(Circle{})
	image.Add(Square{})

	picture := Image{}
	picture.Add(image)
	picture.Add(Image{})

	picture.Draw()
}