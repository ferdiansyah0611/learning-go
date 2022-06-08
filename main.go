package main

import (
	f "fmt"
)

type Product struct {
	id int
	name string
	price int
	total int
}

type Console struct {
	name string
	price int
}

func AddProduct(length int, value Product) Product {
	length += 1
	data := Product{
		id: length,
		name: value.name,
		price: value.price,
		total: value.total,
	}
	return data
}

func decorativeCli(str string){
	length := 44
	for i := 0; i < length; i++ {
		f.Print(str)
		if(i == length - 1){
			f.Println("")
		}
	}
}

func main() {
	var item []Product
	var console []Console
	var purchase int

	console = append(console, Console{
		name: "Cappucino",
		price: 15000,
	})
	console = append(console, Console{
		name: "Americano",
		price: 25000,
	})
	console = append(console, Console{
		name: "Coffe Late",
		price: 20000,
	})

	f.Println("======= Sistem Penjualan Cafe Ferdy ========")
	f.Println("=============== Pilih Produk ===============")
	for i := 0; i < len(console); i++ {
		f.Println(i + 1, console[i].name, "\t\t\t Rp.", console[i].price)
	}
	decorativeCli("=")

	var cli = func() bool {
		var reply int
		var amount int
		var again string

		var inputStart = func() {
			f.Print("Pilih\t> ")
			f.Scanln(&reply)
			f.Print("Jumlah\t> ")
			f.Scanln(&amount)
		}
		inputStart()

		if reply > len(console) {
			inputStart()
		}

		var selection Console = console[reply - 1]
		purchase += selection.price * amount
		item = append(item, AddProduct(len(item), Product{
			name: selection.name,
			price: selection.price,
			total: amount,
		}))

		f.Print("Lagi\t> ")
		f.Scanln(&again)

		if again == "y" || again == "Y"{
			return true
		} else {
			decorativeCli("=")
			for _, value := range item {
				sum := f.Sprintf("[%vx]", value.total)
				f.Println(sum, value.name, "\tRp.", value.price)
			}
			decorativeCli("-")
			f.Println("Total Purchase")
			f.Println("Rp.", purchase)
			return false
		}
	}

	for i := 0; i < 10000; i++ {
		var result bool = cli()
		if result {
			decorativeCli("=")
			cli()
		} else {
			item = []Product{}
			purchase = 0
			decorativeCli("=")
			cli()
		}
	}
}