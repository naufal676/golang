package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Buat channel untuk mengirim data
	dataChan := make(chan map[string]int)

	// Jalankan goroutine untuk mengupdate data setiap 15 detik
	go updateData(dataChan)

	// Loop utama untuk menerima dan mencetak data yang diterima
	for {
		data := <-dataChan
		fmt.Printf("Water: %d, Wind: %d\n", data["water"], data["wind"])
	}
}

func updateData(dataChan chan map[string]int) {
	for {
		// Buat data baru dengan angka acak antara 1-100
		data := map[string]int{
			"water": rand.Intn(100) + 1,
			"wind":  rand.Intn(100) + 1,
		}

		// Kirim data ke channel
		dataChan <- data

		// Tunggu selama 15 detik sebelum mengupdate data berikutnya
		time.Sleep(15 * time.Second)
	}
}
