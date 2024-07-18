package lib

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)


func Kasir() {
	defer fmt.Println("terimakasih")

	products := []map[string]interface{}{
		{"id": 1, "name": "mangga", "harga": 12000},
		{"id": 2, "name": "jeruk", "harga": 2000},
		{"id": 3, "name": "nanas", "harga": 9000},
	}

	for _, v := range products {
		fmt.Println("id: ", v["id"], "nama: ", v["name"], "harga: ", v["harga"])
	}

	println("silahkan masukan no id untuk memesan")
	
	var input string
	var qty string
	fmt.Scanln(&input)
	
	println("silahkan masukan jumlah")
	fmt.Scanln(&qty)

	number, err := strconv.Atoi(input)
	qtynum, err_qty := strconv.Atoi(qty)

	if err != nil {
		println("harus memasukkan angka")
	}
	if err_qty != nil {
		println("harus memasukkan angka")
	}

	isExist := false
	
	println("============================================")
	for _, v := range products {
		if v["id"] == number {
			total := qtynum * int(reflect.ValueOf(v["harga"]).Int())
			fmt.Println("total pesanan anda adalah :", total)
			isExist = true
		}
	}

	if !isExist {
		println("produk tidak ditemukan")
	}
	println("============================================")
	println("apakah anda akan melakukan transaksi lagi ? y/n")
	var isAgain string
	fmt.Scanln(&isAgain)

	if (isAgain == "y"){
		Kasir()
	}else{
		os.Exit(1)
	}

}
