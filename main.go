package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/bslatkin/template-perf/static"
)

type Data struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Sex   string `json:"sex"`
	Legs  int    `json:"legs"`
}

var (
	clientTemplate = template.Must(template.New("static").Parse(static.Files["client_render.tpl"]))
	serverTemplate = template.Must(template.New("dynamic").Parse(static.Files["server_render.tpl"]))
	data           = []Data{
		Data{"Susi", "Silver Tabby with White", "Female (neutered)", 4},
		Data{"Sassy", "Blue Classic Tabby", "Unknown", 4},
		Data{"Minka", "Black Mac Tabby", "Unknown", 4},
		Data{"Moritz", "Black Spotted Tabby", "Unknown", 4},
		Data{"Tigger", "Red & Brown Classic Tabby", "Female", 4},
		Data{"Sophie", "Blue Classic Tabby", "Other", 4},
		Data{"Fluffy", "Pointed with White", "Unknown", 4},
		Data{"Missy", "Red Mackerel Tabby", "Female (neutered)", 4},
		Data{"Blackie", "Solid Chocolate", "Other", 4},
		Data{"Chanel", "Blue Lynx Point", "Other", 4},
		Data{"Oreo", "Brown Mackerel Tabby with White", "Other", 4},
		Data{"Minka", "Cream Point", "Unknown", 4},
		Data{"Félix", "Cream Spotted Tabby", "Male (neutered)", 4},
		Data{"Alfie", "Red Classic Tabby", "Male (neutered)", 3},
		Data{"Lily", "Tortie/Torbie with White", "Female (neutered)", 2},
		Data{"Daisy", "Black Tortie with White", "Female", 4},
		Data{"Missy", "Black Tortie with White", "Unknown", 4},
		Data{"Sam", "White", "Other", 2},
		Data{"Casper", "Blue & Black Mac Tabby", "Unknown", 4},
		Data{"Oscar", "Red & Brown Classic Tabby", "Male", 4},
		Data{"Charlotte", "Blue Mackerel Tabby", "Female", 4},
		Data{"Maggie", "Black Mac Tabby", "Male (neutered)", 4},
		Data{"Minka", "Brown Mackerel Tabby", "Male", 4},
		Data{"Patches", "Blue Tortie", "Unknown", 4},
		Data{"Sassy", "Spotted Tabby", "Male (neutered)", 4},
		Data{"Charlie", "Red Mackerel Tabby", "Other", 4},
		Data{"Muschi", "Cream", "Other", 4},
		Data{"Cleo", "Seal Lynx Point", "Unknown", 4},
		Data{"Oliver", "Brown Classic Tabby with White", "Female (neutered)", 4},
		Data{"Samantha", "Pointed with White", "Unknown", 2},
		Data{"Misty", "Brown Spotted Tabby", "Unknown", 4},
		Data{"Sylvester", "Blue & Black Mac Tabby", "Unknown", 4},
		Data{"Félix", "Red Ticked Tabby", "Male (neutered)", 4},
		Data{"Tom", "Brown Mackerel Tabby with White", "Male (neutered)", 3},
		Data{"Max", "Tortie/Torbie with White", "Female (neutered)", 4},
		Data{"Sophie", "Black Spotted Tabby", "Female", 4},
		Data{"Muschi", "Red & Brown Classic Tabby", "Female (neutered)", 4},
		Data{"Tigger", "Silver Mackerel Tabby", "Other", 4},
		Data{"Tigger", "Red Classic Tabby", "Female", 4},
		Data{"Angel", "Blue Lynx Point", "Other", 4},
		Data{"Smokey", "Cream with White", "Female", 4},
		Data{"Charlie", "Pointed with White", "Unknown", 4},
		Data{"Molly", "Solid Black", "Unknown", 2},
		Data{"Tom", "Brown Mackerel Tabby", "Female", 4},
		Data{"Alfie", "Black with White", "Unknown", 4},
		Data{"Charlie", "Seal Point with White (Mitted Pattern)", "Female (neutered)", 4},
		Data{"Patch", "Brown Mackerel Tabby", "Male", 4},
		Data{"Poppy", "Cream Spotted Tabby", "Male", 4},
		Data{"Lisa", "Seal Point with White (Mitted Pattern)", "Male (neutered)", 4},
		Data{"Smokey", "Black Spotted Tabby", "Other", 4},
		Data{"Patches", "Red and Cream Mac Tabby", "Female", 3},
		Data{"Félix", "Seal Lynx Point", "Other", 4},
		Data{"Tigger", "Cream Point", "Unknown", 4},
		Data{"Tiger", "Red Mackerel Tabby", "Male", 3},
		Data{"Misty", "Black with White (Van Pattern)", "Male (neutered)", 3},
		Data{"Sam", "Brown Mackerel Torbie with White", "Other", 4},
		Data{"Fraidy", "Brown Spotted Tabby", "Other", 4},
		Data{"Whiskers", "Red and a Cream Mackerel Tabby", "Female", 4},
		Data{"Sooty", "Red and Cream Mac Tabby", "Male (neutered)", 4},
		Data{"Whiskers", "Blue Point", "Other", 4},
		Data{"Oreo", "Tortie/Torbie with White", "Other", 4},
		Data{"Lucky", "Red & Black Classic Tabby", "Male", 4},
		Data{"Félix", "Tabby with White", "Male (neutered)", 4},
		Data{"Charlie", "Red Classic Tabby", "Male (neutered)", 4},
		Data{"Lily", "Solid White", "Male (neutered)", 4},
		Data{"Princess", "Seal Tortie Point with White", "Unknown", 4},
		Data{"Tiger", "Black with White", "Male", 3},
		Data{"Sylvester", "Black with White", "Other", 4},
		Data{"Minka", "Brown Classic Tabby with White", "Unknown", 4},
		Data{"Felix", "Silver Tabby with White (Van Pattern)", "Female (neutered)", 4},
		Data{"Lily", "Black with White (Van Pattern)", "Male", 4},
		Data{"Angel", "Red and Cream Mac Tabby", "Male", 4},
		Data{"Oliver", "Solid Black", "Unknown", 1},
		Data{"Charlie", "Red and a Cream Mackerel Tabby", "Male (neutered)", 3},
		Data{"Alfie", "Spotted Tabby", "Unknown", 4},
		Data{"Fluffy", "Solid Chocolate", "Unknown", 4},
		Data{"Chanel", "Blue Mackerel Tabby", "Female", 4},
		Data{"Sammy", "Black with White (van pattern)", "Female (neutered)", 4},
		Data{"Grisou", "Silver Ticked Tabby", "Female (neutered)", 4},
		Data{"Caramel", "Black with White (Van Pattern)", "Female", 4},
		Data{"Sammy", "Tortie/Torbie with White", "Male", 4},
		Data{"Scaredy", "Tabby with White", "Male", 4},
		Data{"Ginger", "Red & Black Classic Tabby", "Female", 4},
		Data{"Jack", "Brown Mackerel Torbie with White", "Male", 4},
		Data{"Sylvester", "Black Classic Tabby with White", "Male", 4},
		Data{"Sooty", "Classic Tabby", "Male (neutered)", 4},
		Data{"Patches", "Black with White (Van Pattern)", "Female (neutered)", 4},
		Data{"Poppy", "Silver Spotted Tabby", "Unknown", 4},
		Data{"Patch", "Red Classic Tabby with White", "Other", 4},
		Data{"Grisou", "Cream", "Other", 4},
		Data{"Minou", "Tortie/Torbie with White", "Female (neutered)", 4},
		Data{"Ti-Mine", "Spotted Tabby", "Male (neutered)", 4},
		Data{"Maggie", "Cream Classic Tabby", "Female", 4},
		Data{"Sooty", "Brown Classic Torbie", "Male (neutered)", 4},
		Data{"Misty", "Blue", "Male (neutered)", 1},
		Data{"Tom", "Black with White", "Other", 4},
		Data{"Minou", "Blue Tortie with White", "Female (neutered)", 4},
		Data{"Sophie", "Blue Spotted Tabby", "Other", 4},
		Data{"Tom", "Red & Black Classic Tabby", "Female (neutered)", 4},
		Data{"Charlotte", "Tabby with White", "Unknown", 1},
		Data{"Sylvester", "Black Mackerel Tabby", "Male (neutered)", 1},
		Data{"Sam", "Siver Ticked Tabby", "Female", 4},
		Data{"Shadow", "Red Ticked Tabby with White", "Male (neutered)", 4},
		Data{"Sophie", "Brown Classic Torbie with White", "Male (neutered)", 4},
		Data{"Sam", "Black with White (van pattern)", "Male (neutered)", 4},
		Data{"Charlie", "Red & Brown Classic Tabby", "Female (neutered)", 4},
		Data{"Grisou", "Silver Classic Tabby", "Unknown", 4},
		Data{"Moritz", "Cream Spotted Tabby", "Unknown", 4},
		Data{"Princess", "Solid Chocolate", "Other", 2},
		Data{"Tom", "Blue", "Male", 4},
		Data{"Sammy", "Silver Classic Torbie", "Male (neutered)", 4},
		Data{"Chanel", "Blue Mackerel Tabby", "Male", 4},
		Data{"Missy", "Brown Classic Torbie with White", "Unknown", 4},
		Data{"Smudge", "Silver Classic Tabby", "Male (neutered)", 4},
		Data{"Felix", "Red Mackerel Tabby", "Male (neutered)", 4},
		Data{"Ginger", "Red Mackerel Tabby", "Female (neutered)", 4},
		Data{"Sylvester", "Blue Mackerel Tabby", "Female", 3},
		Data{"Whiskers", "Blue & Black Mac Tabby", "Female", 3},
		Data{"Tiger", "Chocolate", "Unknown", 4},
		Data{"Samantha", "Silver Tabby with White", "Male (neutered)", 4},
		Data{"Simon", "Red Ticked Tabby with White", "Female (neutered)", 4},
		Data{"Muffin", "Red Classic Tabby", "Other", 4},
		Data{"Grisou", "Blue Lynx Point", "Female", 4},
		Data{"Muschi", "Blue Tortie", "Male (neutered)", 4},
		Data{"Chloe", "Silver Ticked Tabby", "Female", 4},
		Data{"Muschi", "Blue and a Brown Mackerel Tabby", "Male (neutered)", 4},
		Data{"Tiger", "Seal Point with White (Mitted Pattern)", "Unknown", 4},
		Data{"Tom", "Silver Classic Torbie", "Unknown", 4},
		Data{"Lily", "Brown Mackerel Torbie", "Female (neutered)", 4},
		Data{"Maggie", "Blue Lynx Point", "Other", 4},
		Data{"Oliver", "Silver Classic Torbie with White", "Male (neutered)", 4},
		Data{"Minou", "Black Mac Tabby", "Female", 3},
		Data{"Samantha", "Cream", "Female (neutered)", 4},
		Data{"Caramel", "Blue Tortie with White", "Male", 4},
		Data{"Poppy", "Seal Tortie Point with White", "Unknown", 4},
		Data{"Daisy", "Seal & Blue Point with White (Bi-Color Pattern)", "Female (neutered)", 4},
		Data{"Simon", "Black Ticked Tabby", "Other", 4},
		Data{"Princess", "Seal Blue Point with White", "Female (neutered)", 4},
		Data{"Charlotte", "Silver Tabby with White", "Male", 4},
		Data{"Misty", "Brown Mackerel Torbie with White", "Unknown", 4},
		Data{"Lisa", "Cream Classic Tabby", "Unknown", 4},
		Data{"Chanel", "Brown Mackeral Torbie", "Male", 4},
		Data{"Sooty", "Red Ticked Tabby", "Male", 4},
		Data{"Ginger", "Cream Point", "Unknown", 4},
		Data{"Casper", "Blue Lynx Point", "Other", 4},
		Data{"Félix", "Black Tortie", "Male (neutered)", 4},
		Data{"Caramel", "Seal Point", "Female (neutered)", 4},
		Data{"Muschi", "Seal Tortie Point with White", "Male", 4},
		Data{"Princess", "Cream Spotted Tabby", "Female", 4},
		Data{"Pacha", "Blue and a Brown Mackerel Tabby", "Other", 4},
		Data{"Scaredy", "Chocolate", "Other", 4},
		Data{"Moritz", "Brown Classic Torbie", "Other", 4},
		Data{"Oreo", "Black Mackerel Torbie with White", "Unknown", 4},
		Data{"Missy", "Brown Classic Torbie", "Male", 4},
		Data{"Moritz", "Silver Tabby with White (Van Pattern)", "Male", 3},
		Data{"Misty", "Seal Point", "Female (neutered)", 4},
		Data{"Muschi", "Red Point", "Unknown", 4},
		Data{"Max", "Red Classic Tabby with White", "Female (neutered)", 4},
		Data{"Chloe", "White", "Female", 4},
		Data{"Minette", "Cream with White", "Male (neutered)", 4},
		Data{"Sophie", "Blue & Black Mac Tabby", "Female", 4},
		Data{"Chanel", "Seal & Blue Point with White (Bi-Color Pattern)", "Male", 4},
		Data{"Sassy", "Red and Cream Mac Tabby", "Male", 4},
		Data{"Grisou", "Silver Tabby with White", "Unknown", 4},
		Data{"Buddy", "Blue Classic Tabby", "Unknown", 4},
		Data{"Minette", "Blue Mackerel Tabby with White", "Female", 4},
		Data{"Ti-Mine", "Black Classic Torbie with White", "Male", 4},
		Data{"Poppy", "Black Ticked Tabby", "Unknown", 4},
		Data{"Missy", "Blue Cream with White", "Male", 4},
		Data{"Charlie", "Brown Mackerel Tabby", "Female", 3},
		Data{"Jack", "Brown Classic Tabby", "Male", 4},
		Data{"Sooty", "Blue and a Brown Mackerel Tabby", "Male", 2},
		Data{"Kitty", "Black Classic Tabby with White", "Male", 4},
		Data{"Daisy", "Red Ticked Tabby", "Female", 4},
		Data{"Cleo", "Spotted Tabby", "Male (neutered)", 4},
		Data{"Oliver", "Black Mac Tabby", "Female (neutered)", 4},
		Data{"Sam", "Blue Cream with White", "Other", 4},
		Data{"Puss", "Black with White (van pattern)", "Female", 4},
		Data{"Scaredy", "Blue Mackerel Tabby with White", "Male (neutered)", 4},
		Data{"Pacha", "Blue Cream with White", "Other", 4},
		Data{"Tom", "Ticked Tabby", "Unknown", 4},
		Data{"Scaredy", "Silver Classic Torbie", "Female", 4},
		Data{"Shadow", "Silver Spotted Tabby", "Other", 4},
		Data{"Casper", "Tortie/Torbie with White", "Female", 4},
		Data{"Pacha", "Red & Brown Classic Tabby", "Unknown", 4},
		Data{"Smokey", "Blue Point", "Male (neutered)", 4},
		Data{"Minou", "Silver Ticked Tabby", "Female (neutered)", 4},
		Data{"Muschi", "Red Classic Tabby", "Female", 4},
		Data{"Charlie", "Chocolate", "Unknown", 4},
		Data{"Charlotte", "Silver Ticked Tabby", "Male", 4},
		Data{"Minka", "Black Mac Tabby", "Other", 2},
		Data{"Lisa", "Red and Cream Mac Tabby", "Male (neutered)", 4},
		Data{"Whiskers", "Black with White (van pattern)", "Male (neutered)", 4},
		Data{"Lucy", "Blue Lynx Point", "Unknown", 4},
		Data{"Baby", "Black Classic Torbie with White", "Other", 4},
		Data{"Smokey", "Red and Cream Mac Tabby", "Other", 4},
		Data{"Pacha", "Brown Mackeral Torbie", "Male", 2},
		Data{"Tom", "Red Classic Tabby", "Male (neutered)", 4},
		Data{"Princess", "Pointed with White", "Female", 0},
		Data{"Oreo", "Brown Classic Tabby with White", "Male (neutered)", 4},
	}
)

func clientRenderHandler(w http.ResponseWriter, r *http.Request) {
	encoded, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buffer bytes.Buffer
	json.HTMLEscape(&buffer, encoded)

	if err := clientTemplate.Execute(w, template.JS(buffer.Bytes())); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serverRenderHandle(w http.ResponseWriter, r *http.Request) {
	if err := serverTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/client_render", clientRenderHandler)
	http.HandleFunc("/server_render", serverRenderHandle)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
