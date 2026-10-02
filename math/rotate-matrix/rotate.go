package main

import "fmt"

// Give an image represented by an NxN matrix where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees
// Can you do it in-place

func main() {
	image := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	/* intended Output
	{
	   {13,9,5,1},
	   {14, 10,6,2},
	   {15,11,7,3},
	   {16,12,8,4},
	}

	*/
	fmt.Println(image)
	n := 4
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			temp := image[i][j]                   //temp = 1
			image[i][j] = image[n-1-j][i]         // image[0][0] = 13
			image[n-1-j][i] = image[n-1-j][n-1-i] // image[3][0] = 16
			image[n-1-j][n-1-i] = image[j][n-1-i] // image[3][3] = 4
			image[j][n-1-i] = temp
		}
	}
	fmt.Println(image)
}
