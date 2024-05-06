package main

import "math"

// Поля x и y инкапсулированы - они не будут видны вне пакета и доступ к ним вне пакета можно получить только с помощью функций
type Point struct {
	x float64
	y float64
}

// Функция для получения X
func (p *Point) GetX() float64{
	return p.x
}

// Функция для получения Y
func (p *Point) GetY() float64{
	return p.y
}

// Конструктор для Point
func NewPoint(x float64, y float64) Point{
	return Point{x, y}
}

// Функция нахождения расстояния
func FindDistance(p1 Point, p2 Point) float64{
	dist := math.Pow(p1.GetX() - p2.GetX(), 2) + math.Pow(p1.GetY() - p2.GetY(), 2)
	return math.Sqrt(dist)
}

func main() {
	point1 := NewPoint(3.5, 5.5)
	point2 := NewPoint(5, 8.95)
	println(FindDistance(point1, point2))
}