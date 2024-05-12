package main

import (
	"fmt"
	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
var k = koanf.New(".")

func main() {

	//Load dotenv config
	//f := file.Provider("mock/mock.env")
	f := file.Provider(".env")
	if err := k.Load(f, dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println("DB_HOST = ", k.String("DB_HOST"))
	fmt.Println("DB_USER = ", k.String("DB_USER"))
	fmt.Println("DB_PASS = ", k.String("DB_PASS"))
	//fmt.Println("COMMENT = ", k.String("COMMENT"))

	//// Load JSON config.
	//f := file.Provider("mock/mock.json")
	//if err := k.Load(f, json.Parser()); err != nil {
	//	log.Fatalf("error loading config: %v", err)
	//}
	//
	//// Load YAML config and merge into the previously loaded config (because we can).
	//k.Load(file.Provider("mock/mock.yml"), yaml.Parser())
	//
	//fmt.Println("parent's name is = ", k.String("parent1.name"))
	//fmt.Println("parent's ID is = ", k.Int("parent1.id"))
	//
	//// Watch the file and get a callback on change. The callback can do whatever,
	//// like re-load the configuration.
	//// File provider always returns a nil `event`.
	//f.Watch(func(event interface{}, err error) {
	//	if err != nil {
	//		log.Printf("watch error: %v", err)
	//		return
	//	}
	//
	//	log.Println("config changed. Reloading ...")
	//	k.Load(f, json.Parser())
	//	k.Print()
	//})
	//
	//// Block forever (and manually make a change to mock/mock.json) to
	//// reload the config.
	//log.Println("waiting forever. Try making a change to mock/mock.json to live reload")
	//<-make(chan bool)
}
