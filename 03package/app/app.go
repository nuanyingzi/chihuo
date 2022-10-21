package app

import "fmt"

var AppName = "chihuo"
var appVersion = "1.0.0"

func init() {
	fmt.Printf("app name is %s\n", AppName)
	fmt.Printf("app version is %s\n", appVersion)
}
