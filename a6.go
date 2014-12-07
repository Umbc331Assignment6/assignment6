package main

import ("os"
		"fmt"
		"image"
		"image/color"
		"image/png"
 		"strconv"
)

var maxX, maxY int	//globals for maximum X and Y cooridinate values

type point struct {
	xf, yf float64
	xi, yi int
}
//TODO gen_points is wrong outer loop doesnt ever finsh
//		if the first argument is greater than 2
//		Once it hits 4 it starts to finsh less times
/*
 * Generates channel to emit list of point structs
 * (stage 1)
 */
func gen_points() <-chan point {
	fmt.Printf("MaxY: %d MaxX: %d\n", maxY, maxX)
	out := make(chan point)
	go func() {
		for y := -maxY/2; y < maxY; y++ {//TODO out for loop only executes oncefor larger inputs
			fmt.Printf("Outer loop\n")
			fmt.Println(y)
			for x := -maxX/2; x < maxX; x++ {
				xx, yy := scale_pixel(x, y)
				p := point{xf: xx, yf: yy, xi: x, yi: y}
				//fmt.Println(p.yf, p.yi)//TODO doesnt change????
				//fmt.Println(p.xf, p.xi)
				//fmt.Printf("MaxY: %d MaxX: %d\n", maxY, maxX)
				out <- p
			}//end for
		}//end for
		close(out)
	}()//end go func
	return out
}

/*
 * Calculates if point is in the mandlebrot set
 * (stage 2)
 */
func mandlebrot_routine(in <-chan point,iterations int, m *image.NRGBA) {
	go func() {
		for n := range in {
			if mandlebrot( complex(n.xf,n.yf), iterations) {
				newx, newy := normal2imageCoordinate(n.xi,n.yi)
				//fmt.Println(newx, newy)
				m.Set(newx,newy, color.Black)
			} else {
				newx, newy := normal2imageCoordinate(n.xi,n.yi)
				//fmt.Println(newx, newy)
				m.Set(newx,newy, color.White)
			}
		}//end for loop

	}()//end go func 
}

/*
 * Error catching function
 */
func check(e error) {
    if e != nil {
    	fmt.Println(e)
        panic(e)
    }
}

/*	Takes in Complex value and computes if it
 *	is in the mandlebrot set
 */ 
func mandlebrot(c complex128, iterations int) bool{
	z := c
	i := int(0)
	for (i < iterations) {
		z = z*z + c
		x, y := real(z), imag(z)
		if x*x + y*y > 4 {
			return true
		}
		i++
		x, y = real(z), imag(c)
	}
	return false
}
/* Takes zero being at the top left of the screen 
 * translates to zero being at the center
 */
func normal2imageCoordinate(x, y int) (int, int) {
	x = x + maxX/3
	y = y + maxY/2
	return x, y
}

/*
 *	Scales the pixels to lie between;	X: -2, 2
 *										Y: -1, 1
 */
func scale_pixel(x, y int) (float64, float64) {
	newx := float64(x)/(float64(maxX)/3) - 2.0
	newy := float64(y)/(float64(maxY)/2) - 0.0
	//fmt.Printf("X Value %f\n", newx)
	//fmt.Printf("Y Value %f\n", newy)
	return newx, newy
}

func unscale_pixel(x, y float64) (int,int) {
	newx := float64(x)*(float64(maxX)*3) + 2.0
	newy := float64(y)*(float64(maxY)*2) + 0.0
	//fmt.Printf("X Value %f\n", newx)
	//fmt.Printf("Y Value %f\n", newy)
	return int(newx), int(newy)
}


func main() {

	arg1, e := strconv.Atoi(os.Args[1])
	check(e) //check if an error happened
	maxY = arg1 * 2	//garuntees you have plus and minus 1*(arg1) from the center
	
	maxX = arg1 * 3 //garuntees you have plus and minus 1.5*(arg1) from the center
	arg2, e := strconv.Atoi(os.Args[2])
	check(e) //check if an error happened

	fmt.Printf("max X: %d max Y: %d\n",maxX,maxY)
	fmt.Printf("Arg1 %d, Arg2 %d\n", arg1, arg2)
	
	r := image.Rect(0, maxY, maxX, 0) //Make rectangle to store the mandlebrot
	//Create file, f being the filedescriptor
	f, err := os.Create("mandlebrot.png")
	check(err)	//Catch any errors

    m := image.NewNRGBA(r)	//Put the rectangle into a new image object
	/* Parrellized way (doesnt work)*/
	//########################
	the_points := gen_points()
	mandlebrot_routine(the_points, arg2, m)
	//########################
	/* Sequenctial way (works)*/
	/*/////////////////////////////////////////////////////////////////
	for y := -maxY/2; y < maxY; y++ {
		for x := -maxX/2; x < maxX; x++ {
		
			xx, yy := scale_pixel(x,y)	//scales the image coordinates
										//to our mandlebrot range coordinates 
			if mandlebrot( complex(xx,yy), arg2) == true {		//if the pixel escapes then color it black
				newx, newy := normal2imageCoordinate(x,y)
				m.Set(newx,newy, color.Black)	
			} else {											//otherwise color it white
				newx, newy := normal2imageCoordinate(x,y)
				m.Set(newx,newy, color.White)
			}
						
		}//end x for 
	}//end y for
	*//////////////////////////////////////////////////////////////////

	//Makes the image object into a png 
	//f is the filediscriptor created earlier
	if err = png.Encode(f, m); err != nil {
		fmt.Println(err)
	os.Exit(1)
	}
}//end main



