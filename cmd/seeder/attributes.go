package main

// Region attributes
var regionNames = []string{"North America", "South America", "Europe", "Oceania", "Asia"}

// Product attributes
var productAdjectives = []string{"Vintage", "Modern", "Eco friendly", "Luxury", "High end", "Enthusiast", "Refurbished"}
var productCategories = []string{"Appliances", "Clothing", "Books", "Miscellaneous", "Kitchenware", "Computers", "Furniture", "Beauty"}
var productNames = make([]string, 0, len(productAdjectives)*len(productCategories))

// Customer attributes
var customerFirstNames = []string{
	"Javier",
	"Westbrooke",
	"Norry",
	"Cody",
	"Marlow",
	"Dillie",
	"Arie",
	"Vernen",
	"Phil",
	"Emmerich",
}
var customerLastNames = []string{
	"Lupins",
	"Broad",
	"Petroula",
	"Tofts",
	"Montrose",
	"Simmons",
	"Teaser",
	"Clemenceau",
	"Lamb",
	"Kimble",
	"Essam",
	"Hurleigh",
	"Bonifacio",
	"Huller",
	"Grut",
}

// Store attributes
var storeNames = []string{}

// Sales attributes
