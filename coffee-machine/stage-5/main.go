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

	switch coffeeType {
	case 1:
		c.makeCoffee(250, 0, 16, 4)

	case 2:
		c.makeCoffee(350, 75, 20, 7)

	case 3:
		c.makeCoffee(200, 100, 12, 6)

	default:
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
		fmt.Scan(&coffeeType)
	}

}

func (c *coffeeVals) fill() {
	var addWater, addMilk, addBeans, addCups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addWater)

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addMilk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addBeans)

	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scan(&addCups)

	c.water += addWater
	c.milk += addMilk
	c.beans += addBeans
	c.cups += addCups
}

func (c *coffeeVals) makeCoffee(water, milk, beans, money int) {
	w, m, b, cupsLeft := c.water-water > 0, c.milk-milk > 0, c.beans-beans > 0, c.cups-1 > 0
	if w && m && b && cupsLeft {
		c.water -= water
		c.milk -= milk
		c.beans -= beans
		c.cups -= 1
		c.money += money
		fmt.Println("I have enough resources, making you a coffee!")
	} else if !w {
		fmt.Println("Sorry, not enough water!")
	} else if !m {
		fmt.Println("Sorry, not enough milk!")
	} else if !b {
		fmt.Println("Sorry, not enough beans!")
	} else if !cupsLeft {
		fmt.Println("Sorry, not enough cups!")
	}
}

func main() {
	// write your code here
	cV := &coffeeVals{400, 540, 120, 9, 550}

	var action string
	fmt.Println("Write action (buy, fill, take):")
	fmt.Scan(&action)

	for action != "exit" {
		switch action {
		case "buy":
			cV.buy()
		case "fill":
			cV.fill()
		case "take":
			fmt.Printf("I gave you $%d\n", cV.money)
			cV.money = 0
		case "remaining":
			fmt.Println(amountAvail(cV))
		case "exit":
			break
		default:
			fmt.Println("Wrong choice entered")
		}
		fmt.Scan(&action)
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
