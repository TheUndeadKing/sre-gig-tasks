package main

import (
	"fmt"
)

func main() {
	// Define the following array "Menu"}
	Menu := make([]string, 0, 5)

	// Append to it the following item: "hamburger"
	// Append to it the following item: "salad"
	Menu = append(Menu, "salad", "hamburger")
	// Iterate over the list and print for each item Food: <Food name>. Make sure to replace <Food name> with item from the array
	for _, v := range Menu {
		fmt.Println("Food: ", v)
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// Define an array of 5 items. Iterate over it and print for each item the following: This is <ITEM> and its index in the array is <INDEX>

	Menu = append(Menu, "pizza", "biriyani", "ramen")

	for i, k := range Menu {
		fmt.Printf("This is %v and its index in the array is %v\n", k, i)
	}

}
