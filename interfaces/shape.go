package interfaces

type Shape interface {
	Area() int
}

// Square 
type Square struct{
	side int
}
func (s Square) Area() int{
	return s.side*s.side
}

// Rectangle
type Rectangle struct{
	length int
	breadth int
}
func (t Rectangle) Area() int{
	return t.length*t.breadth
}

// Hybrid
type Hybrid struct{
	shapes []Shape
}

func (h *Hybrid) addShape(s Shape){ // Add new shape
	h.shapes = append(h.shapes, s)
}

func (h Hybrid) Area() int { // calculate area of all shapes
	areaSum:=0
	for _,s := range h.shapes{
		areaSum = areaSum+s.Area()
	}
	return areaSum
}
