package main

import (
	"fmt"
	"log"
	"net/http"

	"go-crud/database"
	"go-crud/handlers"
	"go-crud/repositories"
	"go-crud/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using system env")
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}
	defer db.Close()

	// Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Category Routes
	http.HandleFunc("/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	// Product Routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Transaction Routes
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout)
	http.HandleFunc("/api/report", transactionHandler.HandleReport)

	fmt.Println("Server running di port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
