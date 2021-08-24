package main

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sqrt-7/goodhost/pkg/autogen"
)

var (
	userId      = uuid.New().String()
	menuId1     = uuid.New().String()
	categoryId1 = uuid.New().String()
	categoryId2 = uuid.New().String()
	categoryId3 = uuid.New().String()
	foodId1     = uuid.New().String()
	foodId2     = uuid.New().String()
	foodId3     = uuid.New().String()
	foodId4     = uuid.New().String()
	foodId5     = uuid.New().String()
	foodId6     = uuid.New().String()
	foodId7     = uuid.New().String()
	foodId8     = uuid.New().String()
	foodId9     = uuid.New().String()

	currencyPounds = &autogen.AcceptedCurrency{
		Name:          "GBP",
		DisplaySyntax: "£%d",
	}

	currencyEuro = &autogen.AcceptedCurrency{
		Name:          "EUR",
		DisplaySyntax: "%d EUR",
	}

	tagVegan = &autogen.Tag{
		Id:          uuid.New().String(),
		DisplayName: "VEGAN",
	}

	tagVegetarian = &autogen.Tag{
		Id:          uuid.New().String(),
		DisplayName: "VEGETARIAN",
	}

	tagGlutenFree = &autogen.Tag{
		Id:          uuid.New().String(),
		DisplayName: "GLUTEN FREE",
	}
)

func main() {
	breakfastMenu := &autogen.Menu{
		Metadata: &autogen.Metadata{
			CreatedBy: userId,
			CreatedAt: ptypes.TimestampNow(),
		},
		Id:                    menuId1,
		DisplayName:           "Breakfast Menu",
		AcceptedCurrencyNames: []string{currencyPounds.Name, currencyEuro.Name},
		Schedule: []*autogen.MenuSchedule{
			{
				ActiveDay:   autogen.DayOfWeek_MONDAY,
				ActiveHours: []uint32{6, 7, 8, 9},
			},
			{
				ActiveDay:   autogen.DayOfWeek_TUESDAY,
				ActiveHours: []uint32{6, 7, 8, 9},
			},
			{
				ActiveDay:   autogen.DayOfWeek_WEDNESDAY,
				ActiveHours: []uint32{6, 7, 8, 9},
			},
			{
				ActiveDay:   autogen.DayOfWeek_THURSDAY,
				ActiveHours: []uint32{6, 7, 8, 9},
			},
			{
				ActiveDay:   autogen.DayOfWeek_FRIDAY,
				ActiveHours: []uint32{6, 7, 8, 9, 10, 11},
			},
		},
		Categories: []*autogen.MenuCategory{
			{
				Id:          categoryId1,
				DisplayName: "Egg Dishes",
				Items: []*autogen.MenuItem{
					{
						Id:            foodId1,
						DisplayName:   "Scrambled Eggs With Bacon",
						DescriptionMd: "Scrambled eggs served with crispy bacon and a slice of toast.",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 800},
							{CurrencyName: currencyEuro.Name, PriceInCents: 1100},
						},
						TagIds: []string{},
					},
					{
						Id:            foodId2,
						DisplayName:   "Fried Egg Muffin With Cheese",
						DescriptionMd: "Freshly baked muffin filled with fried egg and a slice of cheese.",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 500},
							{CurrencyName: currencyEuro.Name, PriceInCents: 700},
						},
						TagIds: []string{
							tagVegetarian.Id,
						},
					},
					{
						Id:            foodId3,
						DisplayName:   "Eggs Benedict",
						DescriptionMd: "Two poached eggs served with salmon, spinach, and hollandaise sauce.",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 1000},
							{CurrencyName: currencyEuro.Name, PriceInCents: 1300},
						},
						TagIds: []string{
							tagGlutenFree.Id,
						},
					},
				},
			},
			{
				Id:          categoryId2,
				DisplayName: "Pancakes & Muffins",
				Items: []*autogen.MenuItem{
					{
						Id:            foodId4,
						DisplayName:   "American Pancakes With Maple Syrup",
						DescriptionMd: "4 fluffy pancakes served with maple syrup and butter.",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 400},
							{CurrencyName: currencyEuro.Name, PriceInCents: 550},
						},
						TagIds: []string{
							tagVegan.Id,
						},
					},
					{
						Id:            foodId5,
						DisplayName:   "Crepes",
						DescriptionMd: "2 crepes filled with vanilla custard, jam, or nutella",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 400},
							{CurrencyName: currencyEuro.Name, PriceInCents: 550},
						},
						TagIds: []string{
							tagVegetarian.Id,
						},
					},
					{
						Id:            foodId6,
						DisplayName:   "Muffin Of The Day",
						DescriptionMd: "Freshly baked muffin.",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 250},
							{CurrencyName: currencyEuro.Name, PriceInCents: 300},
						},
						TagIds: []string{
							tagVegetarian.Id,
						},
					},
				},
			},
			{
				Id:          categoryId3,
				DisplayName: "Hot Drinks",
				Items: []*autogen.MenuItem{
					{
						Id:            foodId7,
						DisplayName:   "Coffee",
						DescriptionMd: "",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 200},
							{CurrencyName: currencyEuro.Name, PriceInCents: 250},
						},
						TagIds: []string{},
					},
					{
						Id:            foodId8,
						DisplayName:   "Tea",
						DescriptionMd: "",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 200},
							{CurrencyName: currencyEuro.Name, PriceInCents: 250},
						},
						TagIds: []string{},
					},
					{
						Id:            foodId9,
						DisplayName:   "Hot Chocolate",
						DescriptionMd: "",
						Prices: []*autogen.MenuItemPrice{
							{CurrencyName: currencyPounds.Name, PriceInCents: 200},
							{CurrencyName: currencyEuro.Name, PriceInCents: 250},
						},
						TagIds: []string{},
					},
				},
			},
		},
	}

}
