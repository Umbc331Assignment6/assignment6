package main

import ("os"
		"fmt"
		"io"
		"image"
		"image/color"
		"image/png"
 		"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
 img, _, err := image.Decode(r)
 if err != nil {
  return err
 }
 return png.Encode(w, img)
}

func main() {

	arg1, e := strconv.Atoi(os.Args[1])
	arg2, e := strconv.Atoi(os.Args[2])
	if e != nil {
        fmt.Println(e)
    }
	fmt.Printf("Arg1 %d, Arg2 %d\n", arg1, arg2)
	
	r := image.Rect(0, 255, 255, 0)
	//Create file
	f, err := os.Create("x.png")
    if err != nil {
		check(err)
	}
    m := image.NewNRGBA(r)
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			if x % 2 == 0 {
				m.Set(x,y, color.Black)
			} else {
				m.Set(x,y, color.White)
			}
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



