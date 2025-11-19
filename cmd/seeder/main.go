package main

import (
	"simulation/pkg/db"
)

func main() {
	database := db.Init("./database.db")

	migrator := database.Migrator()
	dropTableErr := migrator.DropTable(&db.Sale{}, &db.Product{}, &db.Customer{}, &db.Store{}, &db.Region{})
	if dropTableErr != nil {
		return
	}

	migrationErr := database.AutoMigrate(&db.Sale{}, &db.Product{}, &db.Customer{}, &db.Store{}, &db.Region{})
	if migrationErr != nil {
		return
	}

	var regions = make([]db.Region, len(regionNames))
	var products = make([]db.Product, cap(productNames))
	var customers = make([]db.Customer, 0, len(customerFirstNames)*len(customerLastNames))
	//var stores []db.Store
	//var sales []db.Sale

	populateRegions(&regions)
	result := database.Create(&regions)
	if result.Error != nil {
		panic(result.Error)
	}

	populateProducts(&products)
	result = database.Create(&products)
	if result.Error != nil {
		panic(result.Error)
	}

	populateCustomers(&customers)
	result = database.Create(&customers)
	if result.Error != nil {
		panic(result.Error)
	}

	//populateStores(&stores, &regions)
	//result = database.Create(&stores)
	//if result.Error != nil {
	//	panic(result.Error)
	//}
	//
	//populateSales(&sales, &stores, &products)
	//result = database.Create(&sales)
	//if result.Error != nil {
	//	panic(result.Error)
	//}
}
