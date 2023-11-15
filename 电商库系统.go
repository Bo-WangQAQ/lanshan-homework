package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 商品结构体
type Item struct {
	Name  string
	Price float64
	Stock int
}

// 电子产品结构体，继承自商品结构体
type ElectronicProduct struct {
	Item
	Brand string
	Model string
}

// 电子产品库存管理接口
type ElectronicInventoryManager interface {
	AddElectronicItem(item ElectronicProduct)
	RemoveElectronicItem(itemName string)
	PrintElectronicInventory()
	PrintBrandModelInfo()
}

// 电子产品管理结构体
type ElectronicInventory struct {
	ElectronicItems []ElectronicProduct
}

// 向电子产品管理中添加电子产品
func (eim *ElectronicInventory) AddElectronicItem(item ElectronicProduct) {
	eim.ElectronicItems = append(eim.ElectronicItems, item)
}

// 从电子产品管理中删除电子产品
func (eim *ElectronicInventory) RemoveElectronicItem(itemName string) {
	for i, electronicItem := range eim.ElectronicItems {
		if electronicItem.Name == itemName {
			// 从切片中删除电子产品
			eim.ElectronicItems = append(eim.ElectronicItems[:i], eim.ElectronicItems[i+1:]...)
			fmt.Printf("Electronic item '%s' removed from inventory.\n", itemName)
			return
		}
	}
	fmt.Printf("Electronic item '%s' not found in inventory.\n", itemName)
}

// 读取电子产品管理中的所有电子产品信息
func (eim *ElectronicInventory) PrintElectronicInventory() {
	fmt.Println("Electronic Inventory:")
	for _, electronicItem := range eim.ElectronicItems {
		fmt.Printf("Item: %s\nPrice: $%.2f\nStock: %d\nBrand: %s\nModel: %s\n",
			electronicItem.Name, electronicItem.Price, electronicItem.Stock, electronicItem.Brand, electronicItem.Model)
		fmt.Println("--------------")
	}
}

// 打印电子产品品牌型号信息
func (eim *ElectronicInventory) PrintBrandModelInfo() {
	fmt.Println("Brand and Model Information:")
	for _, electronicItem := range eim.ElectronicItems {
		fmt.Printf("Brand: %s\nModel: %s\n", electronicItem.Brand, electronicItem.Model)
		fmt.Println("--------------")
	}
}

func main() {
	// 创建电子产品管理实例
	electronicInventory := ElectronicInventory{}

	// 从命令行读取用户输入，并执行相应的操作
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPlease choose an action:")
		fmt.Println("1. Add electronic item")
		fmt.Println("2. Remove electronic item")
		fmt.Println("3. View electronic inventory")
		fmt.Println("4. View brand and model information")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Enter electronic item details:")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Price: ")
			var price float64
			fmt.Scan(&price)

			fmt.Print("Stock: ")
			var stock int
			fmt.Scan(&stock)

			fmt.Print("Brand: ")
			_, _ = reader.ReadString('\n') // 清除上一次输入后的换行符
			brand, _ := reader.ReadString('\n')
			brand = strings.TrimSpace(brand)

			fmt.Print("Model: ")
			model, _ := reader.ReadString('\n')
			model = strings.TrimSpace(model)

			electronicItem := ElectronicProduct{
				Item:  Item{Name: name, Price: price, Stock: stock},
				Brand: brand,
				Model: model,
			}
			electronicInventory.AddElectronicItem(electronicItem)
			fmt.Printf("Electronic item '%s' added to inventory.\n", name)

		case "2":
			fmt.Print("Enter electronic item name to remove: ")
			itemName, _ := reader.ReadString('\n')
			itemName = strings.TrimSpace(itemName)
			electronicInventory.RemoveElectronicItem(itemName)

		case "3":
			electronicInventory.PrintElectronicInventory()

		case "4":
			electronicInventory.PrintBrandModelInfo()

		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
