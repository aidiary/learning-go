package lists

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	// (1)
	hobby := [3]string{"Reading", "Games", "Coding"}
	fmt.Println(hobby)

	// (2)
	fmt.Println(hobby[0])
	fmt.Println(hobby[1:])

	// (3)
	mainHobbies := hobby[:2]
	fmt.Println(mainHobbies)

	// (4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// (5)
	courseGoals := []string{"Learn Go", "Build Apps"}
	fmt.Println(courseGoals)

	// (6)
	courseGoals[1] = "Build Web Apps"
	courseGoals = append(courseGoals, "Get a Job")
	fmt.Println(courseGoals)

	// (7)
	products := []Product{
		{"p1", "Book", 9.99},
		{"p2", "Laptop", 999.99}}
	fmt.Println(products)

	newProduct := Product{"p3", "Phone", 699.99}
	products = append(products, newProduct)
	fmt.Println(products)

	// (8)
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 0.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}