package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	initDB()
	migrate()
	genProducts()

	retCode := m.Run()

	os.Exit(retCode)
}

func genProducts() {
	commitsCount := 50
	commits := make(chan bool, commitsCount)

	for i := 0; i < commitsCount; i++ {
		go func() {
			values := []string{}
			for j := 0; j < 200; j++ {
				values = append(values, "('product %[1]v','%[1]v. She stared at him in astonishment, and as she read something of the significant hieroglyphic of his battered face, her lips whitened.','2015-12-11T15:51:26+08:00','2015-12-11T15:51:26+08:00',null)")
			}
			tx := DB.Begin()
			tx.Exec(fmt.Sprintf(`
				INSERT INTO "products" ("title","description","updated_at","created_at","deleted_at") VALUES
				%v;
			`, strings.Join(values, ",")))
			if err := tx.Commit().Error; err != nil {
				log.Println("failed generate test data")
				commits <- false
			}
			commits <- true
		}()
	}

	for i := 0; i < commitsCount; i++ {
		<-commits
	}
}

func BenchmarkLoadProductsWithRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadProductsWithRows()
	}
}

func BenchmarkLoadProductsWithFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadProductsWithFind()
	}
}

func TestLoadProductsWithRows(t *testing.T) {
	products, err := loadProductsWithRows()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(products) == 0 {
		t.Fatalf("Unexpected empty products.")
	}
}

func TestLoadProductsWithFind(t *testing.T) {
	products, err := loadProductsWithFind()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(products) == 0 {
		t.Fatalf("Unexpected empty products.")
	}
}
