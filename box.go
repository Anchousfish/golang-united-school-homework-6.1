package golang_united_school_homework

import "errors"

type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
		shapes:         make([]Shape, 0),
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return errors.New("can't add a shape, the box is full ")
	}

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapesCapacity >= (i+1) && len(b.shapes) >= i+1 {
		return b.shapes[i], nil
	} else {
		return nil, errors.New("no shape found")
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if b.shapesCapacity >= (i+1) && len(b.shapes) >= i+1 {
		v := b.shapes[i]
		for ; i < len(b.shapes); i++ {
			b.shapes[i] = b.shapes[i+1]
		}
		return v, nil
	} else {

		return nil, errors.New("no box found")
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapesCapacity >= (i+1) && len(b.shapes) >= i+1 {
		v := b.shapes[i]
		b.shapes[i] = shape
		return v, nil
	} else {

		return nil, errors.New("nothing to replace")
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var result float64
	for _, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var result float64
	for _, shape := range b.shapes {
		result += shape.CalcArea()
	}
	return result

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	j := 0
	for i := range b.shapes {
		if _, ok := b.shapes[i].(*Circle); ok {
			for j = i; j < len(b.shapes); j++ {
				b.shapes[j] = b.shapes[j+1]
			}
		}
	}
	if j == 0 {
		return errors.New("no circles in the list")
	} else {
		return nil
	}

}
