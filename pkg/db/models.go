package db

import (
	"time"

	"gorm.io/gorm"
)

// Region -  e.g., "North America", "Europe"
type Region struct {
	gorm.Model
	Name string
}

// Store -  A physical store located in a Region
type Store struct {
	gorm.Model
	Name     string
	RegionID uint
	Region   Region // Belongs to Region
}

// Product -  Items for sale
type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    float64
}

// Customer - Names of customers to add complexity to the queries
type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
}

// Sale -  The massive table we will query
type Sale struct {
	gorm.Model
	Date       time.Time
	Amount     float64
	StoreID    uint
	ProductID  uint
	CustomerID uint
	Store      Store    // Belongs to Store
	Product    Product  // Belongs to Product
	Customer   Customer // Belongs to Customer
}
