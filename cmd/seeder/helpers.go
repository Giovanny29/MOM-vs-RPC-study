package main

import (
	"fmt"
	"math/rand/v2"
	"simulation/pkg/db"
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

}

func populateSales(sales *[]db.Sale, stores *[]db.Store, products *[]db.Product) {

}
