package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	exitCommand = "exit"
)

var (
	validDenominations = []int{5000, 2000}
	products           = map[string]int{
		"Aqua":   2000,
		"Sosro":  5000,
		"Cola":   7000,
		"Milo":   9000,
		"Coffee": 12000,
	}
)

func main() {
	sortedProducts := sortByValueDesc(products)
	printProductCatalog(products, sortedProducts)

	reader := bufio.NewReader(os.Stdin)
	for {
		text := getInputText(reader)
		if text == exitCommand {
			return
		}

		denominations, err := parseDenominations(text)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		if !validateDenominations(denominations) {
			fmt.Println("invalid denomination")
			continue
		}

		totalAmount := getTotalAmount(denominations)
		items := getBuyedItem(totalAmount, products, sortedProducts)
		printItems(items)
	}
}

func sortByValueDesc(items map[string]int) []string {
	sortedItems := make([]string, 0, len(items))

	for item := range items {
		sortedItems = append(sortedItems, item)
	}
	sort.SliceStable(sortedItems, func(i, j int) bool {
		return items[sortedItems[i]] > items[sortedItems[j]]
	})

	return sortedItems
}

func printProductCatalog(products map[string]int, sortedProducts []string) {
	fmt.Println("Product Catalog")
	fmt.Println("------------------------------")
	fmt.Println("| Item | Price |")
	fmt.Println("------------------------------")
	for _, name := range sortedProducts {
		fmt.Printf("| %s | %d |\n", name, products[name])
	}
	fmt.Println("------------------------------")
}

func getInputText(reader *bufio.Reader) string {
	fmt.Print("Input your money: ")
	text, _ := reader.ReadString('\n')

	removedChar := []string{"\n", "[", "]", "(", ")", "!", " "}
	for _, char := range removedChar {
		text = strings.Replace(text, char, "", -1)
	}

	return text
}

func parseDenominations(input string) ([]int, error) {
	input = strings.ReplaceAll(input, " ", "")
	denominationStr := strings.Split(input, ",")
	denominations := make([]int, len(denominationStr))

	for i, denomStr := range denominationStr {
		denomination, err := strconv.Atoi(denomStr)
		if err != nil {
			return nil, errors.New("invalid denomination")
		}
		denominations[i] = denomination
	}

	return denominations, nil
}

func validateDenominations(denominations []int) bool {
	for _, denomination := range denominations {
		if !containsInSlice(validDenominations, denomination) {
			return false
		}
	}
	return true
}

func getTotalAmount(denominations []int) int {
	totalAmount := 0
	for _, denom := range denominations {
		totalAmount += denom
	}
	return totalAmount
}

func getBuyedItem(totalAmount int, products map[string]int, sortedProducts []string) map[string]int {
	items := map[string]int{}
	amountRemain := totalAmount

	for _, name := range sortedProducts {
		price := products[name]
		if amountRemain >= price {
			itemCount := getMaxItemCount(amountRemain, price)
			amountRemain -= products[name] * itemCount
			items[name] += itemCount
		}
	}

	return items
}

func printItems(items map[string]int) {
	output := []string{}
	for name, count := range items {
		item := fmt.Sprintf("%d %s", count, name)
		output = append(output, item)
	}
	fmt.Printf("%s\n", strings.Join(output, ", "))
}

func getMaxItemCount(amount int, price int) int {
	if amount <= 0 || price <= 0 {
		return 0
	}
	return amount / price
}

func containsInSlice(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
