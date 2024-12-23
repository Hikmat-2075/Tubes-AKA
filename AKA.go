package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Struktur data untuk item inventaris
type Item struct {
	ID    string
	Nama  string
	Lokasi string
}

// **Pencarian Linear Iteratif**
func linearSearchIterative(inventory []Item, targetID string) int {
	time.Sleep(10 * time.Millisecond) 
	for i, item := range inventory {
		if item.ID == targetID {
			return i 
		}
	}
	return -1 
}

// **Pencarian Linear Rekursif**
func linearSearchRecursive(inventory []Item, targetID string, index int) int {
	time.Sleep(10 * time.Millisecond) 
	if index >= len(inventory) {
		return -1 
	}
	if inventory[index].ID == targetID {
		return index // Base case: ditemukan
	}
	return linearSearchRecursive(inventory, targetID, index+1) 
}

// Fungsi untuk membuat data inventaris secara acak
func generateInventory(size int) []Item {
	inventory := make([]Item, size)
	for i := 0; i < size; i++ {
		randId := rand.Intn(10000000000)
		inventory[i] = Item{
			ID:    fmt.Sprintf("ITEM-%d", randId),
			Nama:  fmt.Sprintf("Barang-%d", i+1),
			Lokasi: fmt.Sprintf("Rak %d", (i%10)+1),
		}
	}
	return inventory
}

func main() {
	sizes := []int{1, 10, 20, 50, 100, 500, 1000, 2000, 5000, 10000}

	const iterations = 1 // Jumlah pencarian untuk di-rata-ratakan

	// Cetak header tabel dengan jarak yang rapi
	fmt.Printf("%-15s %-25s %-25s\n", "Ukuran Input", "Linear Iteratif (s)", "Linear Rekursif (s)")
	fmt.Println("-------------------------------------------------------------")

	for _, size := range sizes {
		inventory := generateInventory(size)
		var targetItemID string

		if size > 0 {
			targetItemID = inventory[size-1].ID // mencari item terakhir di list, untuk kasus terburuk linear search
		} else {
			targetItemID = "ITEM-NONEXISTENT" // untuk pencarian biner dengan input 0
		}

		// Linear Iteratif
		var totalLinearIterativeTime int64
		for i := 0; i < iterations; i++ {
			start := time.Now()
			linearSearchIterative(inventory, targetItemID)
			duration := time.Since(start).Nanoseconds() // Menggunakan Nanoseconds untuk presisi tinggi
			totalLinearIterativeTime += duration
		}
		avgLinearIterativeTime := float64(totalLinearIterativeTime) / 1_000_000_000 // Konversi ke detik

		// Linear Rekursif
		var totalLinearRecursiveTime int64
		for i := 0; i < iterations; i++ {
			start := time.Now()
			linearSearchRecursive(inventory, targetItemID, 0)
			duration := time.Since(start).Nanoseconds() // Menggunakan Nanoseconds untuk presisi tinggi
			totalLinearRecursiveTime += duration
		}
		avgLinearRecursiveTime := float64(totalLinearRecursiveTime) / 1_000_000_000 // Konversi ke detik

		// Print output dalam format tabel yang rapi dengan jarak antar kolom
		fmt.Printf("%-15d %-25.8f %-25.8f\n", size, avgLinearIterativeTime, avgLinearRecursiveTime)
	}
}
