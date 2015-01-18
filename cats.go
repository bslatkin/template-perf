package main

import (
	"math/rand"
	"time"
)

var (
	// From http://en.wikipedia.org/wiki/Popular_cat_names
	names = []string{
		"Alfie",
		"Angel",
		"Baby",
		"Bella",
		"Blackie",
		"Blacky",
		"Buddy",
		"Caramel",
		"Casper",
		"Chanel",
		"Charlie",
		"Charlotte",
		"Charly",
		"Chloe",
		"Cleo",
		"Daisy",
		"Felix",
		"Fluffy",
		"Fraidy",
		"Félix",
		"Ginger",
		"Grisou",
		"Jack",
		"Jasper",
		"Kitty",
		"Lily",
		"Lisa",
		"Lucky",
		"Lucy",
		"Maggie",
		"Max",
		"Millie",
		"Mimi",
		"Minette",
		"Minka",
		"Minou",
		"Missy",
		"Misty",
		"Molly",
		"Moritz",
		"Muffin",
		"Muschi",
		"Oliver",
		"Oreo",
		"Oscar",
		"Pacha",
		"Patch",
		"Patches",
		"Poppy",
		"Princess",
		"Puss",
		"Sam",
		"Samantha",
		"Sammy",
		"Sassy",
		"Scaredy",
		"Shadow",
		"Simba",
		"Simon",
		"Smokey",
		"Smudge",
		"Sooty",
		"Sophie",
		"Susi",
		"Sylvester",
		"Ti-Mine",
		"Tiger",
		"Tigger",
		"Tom",
		"Whiskers",
	}
	// From http://www.seregiontica.org/Colors/intro.htm
	colors = []string{
		"Bengal",
		"Black",
		"Black Classic Tabby with White",
		"Black Classic Torbie with White",
		"Black Mac Tabby",
		"Black Mackerel Tabby",
		"Black Mackerel Torbie with White",
		"Black Spotted Tabby",
		"Black Ticked Tabby",
		"Black Tortie",
		"Black Tortie with White",
		"Black with White",
		"Black with White (van pattern)",
		"Black with White (Van Pattern)",
		"Blue",
		"Blue & Black Mac Tabby",
		"Blue and a Brown Mackerel Tabby",
		"Blue Classic Tabby",
		"Blue Cream Tortie",
		"Blue Cream with White",
		"Blue Lynx Point",
		"Blue Mackerel Tabby",
		"Blue Mackerel Tabby with White",
		"Blue Point",
		"Blue Spotted Tabby",
		"Blue Tortie",
		"Blue Tortie with White",
		"Brown Classic Tabby",
		"Brown Classic Tabby with White",
		"Brown Classic Torbie",
		"Brown Classic Torbie with White",
		"Brown Mackeral Torbie",
		"Brown Mackerel Tabby",
		"Brown Mackerel Tabby with White",
		"Brown Mackerel Torbie",
		"Brown Mackerel Torbie with White",
		"Brown Spotted Tabby",
		"Brown Spotted Torbie with White",
		"Brown Ticked Tabby",
		"Chocolate",
		"Classic Tabby",
		"Cream",
		"Cream Classic Tabby",
		"Cream Point",
		"Cream Spotted Tabby",
		"Cream with White",
		"Pointed with White",
		"Red",
		"Red & Black Classic Tabby",
		"Red & Brown Classic Tabby",
		"Red and a Cream Mackerel Tabby",
		"Red and Cream Mac Tabby",
		"Red Classic Tabby",
		"Red Classic Tabby with White",
		"Red Mackerel Tabby",
		"Red Point",
		"Red Ticked Tabby",
		"Red Ticked Tabby with White",
		"Seal & Blue Point with White (Bi-Color Pattern)",
		"Seal Blue Point with White",
		"Seal Lynx Point",
		"Seal Point",
		"Seal Point with White (Mitted Pattern)",
		"Seal Tortie Point with White",
		"Silver Classic Tabby",
		"Silver Classic Torbie",
		"Silver Classic Torbie with White",
		"Silver Mackerel Tabby",
		"Silver Spotted Tabby",
		"Silver Tabby with White",
		"Silver Tabby with White (Van Pattern)",
		"Silver Ticked Tabby",
		"Siver Ticked Tabby",
		"Solid Black",
		"Solid Chocolate",
		"Solid White",
		"Spotted Tabby",
		"Tabby with White",
		"Ticked Tabby",
		"Tortie/Torbie with White",
		"White",
	}
	sexes = []string{
		"Male",
		"Male (neutered)",
		"Female",
		"Female (neutered)",
		"Other",
		"Unknown",
	}
	legOptions []int
)

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		legOptions = append(legOptions, 4)
	}
	for i := 0; i < 10; i++ {
		legOptions = append(legOptions, 3)
	}
	for i := 0; i < 5; i++ {
		legOptions = append(legOptions, 2)
	}
	for i := 0; i < 2; i++ {
		legOptions = append(legOptions, 2)
	}
	legOptions = append(legOptions, 0)
}

func getCats(count int) []Data {
	result := make([]Data, 0, count)
	for i := 0; i < count; i++ {
		name := names[rand.Intn(len(names))]
		color := colors[rand.Intn(len(colors))]
		sex := sexes[rand.Intn(len(sexes))]
		legs := legOptions[rand.Intn(len(legOptions))]
		result = append(result, Data{name, color, sex, legs})
	}
	return result
}
