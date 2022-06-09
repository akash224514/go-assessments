package main

import "fmt"

func main() {
	fmt.Println("\n Question 01 -- Create a small program which accepts two slices of colors.It must return a slice containing the colors that appear in either or both slices")

	c1 := []string{"Red", "Black", "White"}
	c2 := []string{"Black", "Yellow", "Orange"}

	res := make([]string, 0)
	fmt.Println("\n**Inputed slices**")
	fmt.Println("Slice 1:", c1)
	fmt.Println("Slice 2:", c2)

	for _, ele := range c1 { // By this loop we are finding a common element from both slices and making that element blank
		for j, ele2 := range c2 {
			if ele == ele2 {
				c2[j] = ""
			}
		}
	}

	res = append(res, c1...)
	res = append(res, c2...) // Slice res contains all elements from both slices including a blanked element

	finalOutput := []string{} // Creating a new slice to remove that blank element

	for _, ele := range res { // By this loop we are omitting the blanked element
		if ele != "" {
			finalOutput = append(finalOutput, ele)
		}
	}

	fmt.Println("\n**Final output slice**")
	fmt.Println("Final Ouput:", finalOutput) // Desired output
}

/* APPROACH ->
- There is NO DELETE in a slice in a Golang
- In Golang slices are not that high level
- Here, initially we found out the common element from both slices
- After that we blanked that common element and appended to a another slice called "res"
- Now this "res" slice contains all the elements from both slices including that blanked element.
- In order to remove or delete that element, we simply appended all the elements from "res" except that blanked element to a slice called "finalOutput"

*/
