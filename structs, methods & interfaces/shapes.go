package shapes

// Perimeter function gets the 2*length + 2*width the perimeter of a rectangle
func Perimeter(length float64, width float64) float64{
	return 2 * (length + width)
}

//Area function finds the area of a rectangle
func Area(length float64, width float64) float64{
	return (length * width)
}
