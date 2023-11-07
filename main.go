package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func main() {
	count := 20
	// products := make([]models.Product, count)
	products := make([][]string, count)
	for i := 0; i < count; i++ {
		products[i] = []string{uuid.New().String(), fake.CarMaker(), fake.CarType(), uuid.New().String(), fmt.Sprintf("%d", fake.IntRange(1, 100))}
	}

	f, err := os.Create("products.csv")
	if err != nil {
		log.Fatalf("failed to create, error: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	err = w.WriteAll(products)
	if err != nil {
		log.Fatalf("failed to write file, error: %v", err)
	}
}
