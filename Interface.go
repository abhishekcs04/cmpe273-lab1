package main

import "math"

type Shape interface { 

 Perimeter() float64
} 

type Rectangle struct {
 

 height float64
 width float64
} 

type Circle struct { 

radius float64
} 

func (r *Rectangle) Perimeter() float64 { 


 return  2 * (r.height + r.width)
} 

func (c *Circle) Perimeter() float64  {

 return 2 * math.Pi * c.radius
} 

func Perimeter(s Shape) float64  {  
 
 return s.Perimeter()
} 



