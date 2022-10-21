package main

import (
	"encoding/json"
	"fmt"
	"github/nuanyingzi/chihuo/06struct/app"
)

func main() {
	mod := app.App{Name: "chihuo", Version: 2}
	fmt.Printf("mod name = %s", mod.Name)

	jstr, _ := json.Marshal(mod)
	fmt.Printf("json string = %s", jstr)
}
