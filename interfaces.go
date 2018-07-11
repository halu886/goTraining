package main
 import(
	 "fmt"
	 "math"
 )

 type geometry interface{
	 area() float64
	 perim() float64
 }

 type react struct{
	 width,height float64
 }

 type circle struct{
	 redius float64
 }

 func (r react) area() float64 {
	 return r.width*r.height
 }

 func (r react) perim() float64{
	 return 2*r.height+2*r.width
 }

 func (c circle) area() float64  {
	 return math.Pi*c.redius*c.redius
 }

 func (c circle) perim() float64 {
	 return 2*math.Pi*c.redius
 }

 func measure(g geometry)  {
	 fmt.Println(g.area())
	 fmt.Println(g.perim())
 }

 func Interfaces()  {
	 r:=react{width:3,height:4}
	 c:=circle{redius:2}

	 measure(r)
	 measure(c)
 }

 func main() {
	 Interfaces()
 }