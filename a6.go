package main

import ("os"
		"fmt"
		"image"
		"image/color"
		"image/png"
 		"strconv"
)

var maxX, maxY int
var stepX, stepY float64

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func mandlebrot(c complex128,iterations int) bool{
	z := c
	i := int(0)
	//x, y := real(c), imag(c)
	//fmt.Printf("mandlebrotX Value %f\n", x)
	for (i < iterations) {
		z = z*z + c
		x, y := real(z), imag(z)
		if x*x + y*y > 4 {
			fmt.Println("\nescapes!\n")
			return true
		}
		i++
		x, y = real(z), imag(c)
	}
	return false
}

func normal2imageCoordinate(x, y int) (int,int) {
	x = x + maxX/2
	y = y + maxY/2
	return x, y
}

func scale_pixel(x, y int) (float64,float64) {
	newx := float64(x)/(float64(maxX)) - 1.0
	newy := float64(y)/(float64(maxY)) - 1.0
	fmt.Printf("X Value %f\n", newx)
	fmt.Printf("Y Value %f\n", newy)
	return newx, newy
}


func main() {

	arg1, e := strconv.Atoi(os.Args[1])
	maxY = arg1 * 2
	
	stepY = 2.0/float64(arg1)
	maxX = arg1 * 3
	fmt.Println(maxX)
	arg2, e := strconv.Atoi(os.Args[2])
	if e != nil {
        fmt.Println(e)
    }
	fmt.Printf("Arg1 %d, Arg2 %d\n", arg1, arg2)
	
	r := image.Rect(0, maxY, maxX, 0)
	//Create file
	f, err := os.Create("mandlebrot.png")
    if err != nil {
		check(err)
	}
    m := image.NewNRGBA(r)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
		
			//fmt.Println(x)
			xx, yy := scale_pixel(x,y)
			//fmt.Println(xx,yy)
			if mandlebrot( complex(xx,yy),arg2) == true {
				newx, newy := normal2imageCoordinate(x,y)
				fmt.Println(newx,newy)
				m.Set(newx,newy, color.Black)	
			} else {
				newx, newy := normal2imageCoordinate(x,y)
				m.Set(newx,newy, color.White)
			}
			
			/*
			if x % 2 == 0 {
				newx, newy := normalx2imageCoordinate(x,y)
				//fmt.Println(newx,newy)
				m.Set(newx,newy, color.Black)
			} else {
				newx, newy := normal2imageCoordinate(x,y)
				m.Set(newx,newy, color.White)
			}
			*/
		} 
	}
	
	if err = png.Encode(f, m); err != nil {
		fmt.Println(err)
	os.Exit(1)
	}


	// Dx and Dy return a rectangle's width and height.
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // prints 3 4 false
	
	//err := ioutil.WriteFile("test.txt", r, 0644)
	//check(err)
	
}



