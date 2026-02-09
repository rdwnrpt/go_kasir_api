package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go_kasir_api/config"
	"go_kasir_api/database"
	"go_kasir_api/handlers"
	"go_kasir_api/repositories"
	"go_kasir_api/services"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("❌ Failed to load config:", err)
	}

	// Init database
	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	defer db.Close()

	// Product DI
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Category DI
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Routes (NO TRAILING SPACES!)
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running with Supabase ✅",
		})
	})

	// Start server
	addr := "0.0.0.0:" + cfg.Port
	fmt.Printf("🚀 Server running at http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}