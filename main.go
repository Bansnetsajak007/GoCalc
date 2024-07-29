//file with the main function

package main

import (
	"Gocalc/repl"
	"fmt"
)
func main() {
    fmt.Println(`
   ______      ______      __    
  / ____/___  / ____/___ _/ /____
 / / __/ __ \/ /   / __ '/ / ___/
/ /_/ / /_/ / /___/ /_/ / / /__  
\____/\____/\____/\__,_/_/\___/  
                                 
    `)
    repl.Start()
}