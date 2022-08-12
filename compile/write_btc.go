package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func f_write(f *os.File, code []byte) {
	_, err := f.Write(code)
	check(err)
}

func f_read(file_path string) []byte {
	dat, err := os.ReadFile(file_path)
	check(err)
	return dat
}

func open_file(file_path string) *os.File {
	var btcfile *os.File
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		btcfile, _ = os.Create(file_path)
	} else {
		btcfile, _ = os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, 0666)
	}
	return btcfile
}

func main() {
	file_path := "./5.btc"
	btcfile := open_file(file_path)
	defer btcfile.Close()
	code := []byte{1, 10, 2, 3, 4}
	f_write(btcfile, code)

	data := f_read(file_path)

	for _, bt := range data {
		fmt.Printf("%v ", bt)
	}
}
