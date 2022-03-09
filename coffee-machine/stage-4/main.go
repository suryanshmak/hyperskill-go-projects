package main

import (
	"fmt"
)

type coffeeVals struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

func (c *coffeeVals) buy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var coffeeType int
	fmt.Scan(&coffeeType)

	if coffeeType == 1 {
		c.water, c.beans, c.money, c.cups = c.water-250, c.beans-16, c.money+4, c.cups-1
	} else if coffeeType == 2 {
		c.water, c.milk, c.beans, c.money, c.cups = c.water-350, c.milk-75, c.beans-20, c.money+7, c.cups-1
	} else {
		c.water, c.milk, c.beans, c.money, c.cups = c.water-200, c.milk-100, c.beans-12, c.money+6, c.cups-1
	}

	fmt.Println(amountAvail(c))
}

func (c *coffeeVals) take() {
	var addWater, addMilk, addBeans, addCups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addBeans)

	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&addCups)

	c.water, c.milk, c.beans, c.cups = c.water+addWater, c.milk+addMilk, c.beans+addBeans, c.cups+addCups
	fmt.Println(amountAvail(c))
}

func main() {
	// write your code here
	cV := &coffeeVals{400, 540, 120, 9, 550}

	fmt.Println(amountAvail(cV))

	var action string
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&action)

	switch action {
	case "buy":
		cV.buy()
	case "fill":
		cV.take()
	case "take":
		cV.money = 0
		fmt.Println(amountAvail(cV))
	default:
		fmt.Println("Wrong choice entered")
	}
}

func amountAvail(cV *coffeeVals) string {
	return fmt.Sprintf("The coffee machine has:\n"+
		"%d ml of water\n"+
		"%d ml of milk\n"+
		"%d g of coffee beans\n"+
		"%d disposable cups\n"+
		"$%d of money", cV.water, cV.milk, cV.beans, cV.cups, cV.money)
}
