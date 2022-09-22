package main

import "fmt"

var banner = `
     _                   _             
    | |                 | |            
  _ | | ___  ____  ____ | | ____  ____ 
 / || |/ _ \|  _ \|  _ \| |/ _  )/ ___)
( (_| | |_| | | | | | | | ( (/ /| |    
 \____|\___/| ||_/| ||_/|_|\____)_|    
            |_|   |_|                  

v%s
			`

func main() {
	fmt.Println(fmt.Sprintf(banner, "1.0.0"))
}
