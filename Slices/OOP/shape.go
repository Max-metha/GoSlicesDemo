package main

type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float32{
	return r.width * 2 + r.height * 2
}