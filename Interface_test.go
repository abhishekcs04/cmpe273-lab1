package main

import  "testing"

type circleTest struct {  

   radius float64
   perimeter float64
} 		

type RecTest struct {  

   height,width float64
   perimeter float64
} 		

var CircleTests= []circleTest {  // Test cases for Circle {Radius, Perimeter} 

  {2,12.566370614359172}, 
  {3,18.84955592153876}, 
  {4.0,25.132741228718345},
} 

var RectTests = []RecTest { 
    {2,4,12}, 
   {2,5,14},
   {5,4,18},
} 
		
func TestGetPerimeter(t *testing.T) {
  
  for _,val := range CircleTests { 

      if val.perimeter !=  Perimeter(&Circle{radius:val.radius}) { 

			 t.Errorf("Test Fail")
       }
	 }
	
	for _,val := range RectTests { 

      if val.perimeter !=  Perimeter(&Rectangle{height:val.height,width:val.width}) { 
		    
			 t.Errorf("Test Fail")
       }
	 	
  }
} 


 
 


   