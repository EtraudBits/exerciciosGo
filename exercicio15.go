package main

import (
	"fmt"
)

func setAlarm(empregado, ferias bool) bool {
		if empregado == true && ferias == false {
			return true
		}
		return false
	}

func main () {

	fmt.Println (setAlarm(true, false))
	fmt.Println (setAlarm(true, true))
	fmt.Println (setAlarm(false, true))
	fmt.Println (setAlarm(false, false))

	fmt.Printf ("Empregado e não de férias: %t\n", setAlarm(true, false))
	fmt.Printf ("Empregado e de férias: %t\n", setAlarm(true, true))
	fmt.Printf ("Não empregado e de férias: %t\n", setAlarm(false, true))
	fmt.Printf ("Não empregado e não de férias: %t\n", setAlarm(false, false))
	
	

}