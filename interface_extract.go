package main

import "fmt"

/**
接口提取
*/

// Sp 形状接口
type Sp interface {
	Area() float32
}

type TopologicalGenus interface {
	rank() int
}

// Sqr 四边形结构体
type Sqr struct {
	side float32
}

func (s *Sqr) Area() float32 {
	return s.side * s.side
}

func (s Sqr) rank() int {
	return 1
}

type Rect struct {
	width, height float32
}

func (r Rect) Area() float32 {
	return r.width * r.height
}

func (r Rect) rank() int {
	return 2
}

// 显式的指名类型实现了某个接口

type Fooer interface {
	Foo()
	ImplementsFooer()
}

type Bar struct {
}

func (b Bar) Foo() {

}

func (b Bar) ImplementsFooer() {

}

func main() {
	r := Rect{10, 5}

	// 此处 q 为 Sqr 结构体的内存地址
	q := &Sqr{5}

	shapes := []Sp{r, q}

	fmt.Println(shapes)

	for i, _ := range shapes {
		fmt.Println("Shape details: ", shapes[i])
		fmt.Println("area of this shape is: ", shapes[i].Area())
	}

	topgen := []TopologicalGenus{r, q}

	for i, _ := range topgen {
		fmt.Println("rank is", topgen[i].rank())
	}
}
