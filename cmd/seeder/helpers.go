package main

import (
	"fmt"
	"math/rand/v2"
	"simulation/pkg/db"
	"time"

	"gorm.io/gorm"
)

func populateRegions(regions *[]db.Region) {
	for index, regionName := range regionNames {
		(*regions)[index] = db.Region{Name: regionName}
	}
}

func populateProducts(products *[]db.Product) {
	for _, adjective := range productAdjectives {
		for _, category := range productCategories {
			productNames = append(productNames, fmt.Sprintf("%s %s", adjective, category))
		}
	}
	for index, name := range productNames {
		(*products)[index] = db.Product{Name: name, Category: productCategories[rand.IntN(len(productCategories))], Price: float64(rand.IntN(150000) / 10)}
	}
}

func populateCustomers(customers *[]db.Customer) {
	for _, firstName := range customerFirstNames {
		for _, lastName := range customerLastNames {
			*customers = append(*customers, db.Customer{FirstName: firstName, LastName: lastName})
		}
	}
}

func populateStores(stores *[]db.Store, regions *[]db.Region) {
	for _, adjective := range storeAdjectives {
		for _, brand := range storeBrands {
			storeNames = append(storeNames, fmt.Sprintf("%s %s", adjective, brand))
		}
	}

	for index, name := range storeNames {
		randomIndex := rand.IntN(len(*regions))
		randomRegionID := (*regions)[randomIndex].ID

		(*stores)[index] = db.Store{
			Name:     name,
			RegionID: randomRegionID,
		}
	}
}

func populateSales(sales *[]db.Sale, stores *[]db.Store, products *[]db.Product, customers *[]db.Customer) {
	for index := 0; index < 10000; index++ {
		randomYear := 2020 + rand.IntN(6)
		randomMonth := time.Month(1 + rand.IntN(12))
		randomDay := 1 + rand.IntN(31)

		randomDate := time.Date(
			randomYear,
			randomMonth,
			randomDay,
			rand.IntN(23),
			rand.IntN(59),
			rand.IntN(59),
			0,
			time.UTC,
		)

		randomAmount := rand.UintN(99)
		randomStoreID := (*stores)[rand.IntN(len(*stores))].ID
		randomProductID := (*products)[rand.IntN(len(*products))].ID
		randomCustomerID := (*customers)[rand.IntN(len(*customers))].ID

		(*sales)[index] = db.Sale{
			Model: gorm.Model{
				CreatedAt: randomDate,
				UpdatedAt: randomDate,
			},
			Amount:     randomAmount,
			StoreID:    randomStoreID,
			ProductID:  randomProductID,
			CustomerID: randomCustomerID,
			Date:       randomDate,
		}
	}
}
