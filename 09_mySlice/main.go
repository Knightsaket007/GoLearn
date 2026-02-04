package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Welcome to Slice")

	var fruitList = []string{"Apple", "Tomato", "Peach", "Bannana", "Swayberry"}
	fmt.Printf("Type of fruitList... %T\n", fruitList)

	// fruitList=append(fruitList[:2]);
	// fmt.Println("Fruitlist is...", fruitList)

	// fruitList=append(fruitList[1:]);
	// fmt.Println("Fruitlist after [3:]...", fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println("Fruitlist after [1:4]...", fruitList)

	//--==-=--==-=- Make-=-====//
	//--==-=--==-=- Make-=-====//
	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 324

	fmt.Println("highscore...", highscore)

	// sort
	sort.Ints(highscore)
	fmt.Println("sorted highscore...", highscore)
	fmt.Println("Is highscore sorted...", sort.IntsAreSorted(highscore))

	// Delete value using append..
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("courses are..", courses)
	
	var position int = 2
	
	courses = append(courses[:2], courses[position+1:]...);

	fmt.Println("courses after append all slice data except 'position'...", courses)

}
