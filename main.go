package main

import (
	"fmt"
	"runtime"

	"bitdump/bitdump"
)

func main() {
	fmt.Println(`
888      d8b 888        888                                 
888      Y8P 888        888                                 
888          888        888                                 
88888b.  888 888888 .d88888 888  888 88888b.d88b.  88888b.  
888 "88b 888 888   d88" 888 888  888 888 "888 "88b 888 "88b 
888  888 888 888   888  888 888  888 888  888  888 888  888 
888 d88P 888 Y88b. Y88b 888 Y88b 888 888  888  888 888 d88P 
88888P"  888  "Y888 "Y88888  "Y88888 888  888  888 88888P"  
                                                   888      
                                                   888      
                                                   888                      									
	`)

	runtime.GOMAXPROCS(4)

	bitdump.InitClient()
}
