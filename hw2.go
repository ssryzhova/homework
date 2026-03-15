package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func main() {
	//easy()
	//medium()
	hard()
}

// easy
type Figure struct {
	Name   string
	a      float64
	b      float64
	r      float64
	square float64
}
type Shape interface {
	Area()
}

func (f *Figure) Area() {
	if f.r != 0 {
		f.square = f.r * f.r * 3.14
	} else {
		f.square = f.a * f.b
	}
}
func square(s Shape) {
	s.Area()
}
func easy() {
	circle := Figure{
		Name: "Circle",
		r:    21.4,
	}
	rectangle := Figure{
		Name: "Rectangle",
		a:    12.4,
		b:    13.5,
	}
	square(&circle)
	square(&rectangle)

	slog.Info("Square of circle:")
	fmt.Println(circle.square)
	slog.Info("Square of rectangle:")
	fmt.Println(rectangle.square)
}

// medium
type transport struct {
	Name string
	V    int
}
type Vechile interface {
	Start(speed int)
	Stop()
}

func (t *transport) Start(speed int) {
	t.V = speed
	fmt.Println("The", t.Name, "start, speed:", t.V)

}
func (t *transport) Stop() {
	t.V = 0
	fmt.Println("The", t.Name, "stop, speed:", t.V)
}
func run(v Vechile, speed int) {
	v.Start(speed)
	v.Stop()
}
func medium() {
	machine := transport{
		Name: "Machine",
	}
	motobyke := transport{
		Name: "Motobyke",
	}
	run(&machine, 100)
	run(&motobyke, 130)
	fmt.Println(machine, motobyke)
}

// hard
type File struct {
	text []byte
}
type Console struct {
}
type Writer interface {
	Write([]byte) (int, error)
}

func (f *File) Write(text []byte) (int, error) {
	file, err := os.Create("hard.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return 0, err
	}
	defer file.Close()
	n, err := file.Write(text)
	if err != nil {
		return n, err
	}
	fmt.Println("Done")
	return n, nil
}
func (c *Console) Write(text []byte) (int, error) {
	n, err := fmt.Println(string(text))
	return n, err
}
func (c *Console) Read() {
	file, err := os.Open("hard.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	text := make([]byte, 64)
	for {
		n, err := file.Read(text)
		if err == io.EOF {
			break
		}
		fmt.Print(string(text[:n]))
	}
	fmt.Println()
}

func write(w Writer, text []byte) {
	w.Write(text)
}
func hard() {
	file := &File{}
	write(file, []byte("1 2 3"))
	reading := &Console{}
	reading.Read()
}
