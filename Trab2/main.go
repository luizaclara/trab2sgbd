// Trabalho 02 SGBD Prof Zé Maria
// Luiza Clara de Albuquerque Pacheco 493478
// Sabrina Silveira Oliveira 494013

package main

import (
	//"Trab2/directory"
	"fmt"
)

func main() {
	var option string
	
	for {
		fmt.Println("\n ******************** NIVEL ISOLAMENTO ******************** ")
		fmt.Println("[ 1 ] READ UNCOMMITTED")
		fmt.Println("[ 2 ] READ COMMITTED")
		fmt.Println("[ 3 ] REPEATABLE READ")
		fmt.Println("[ 4 ] SERIALIZABLE")
		fmt.Println("[ 5 ] SAIR")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&option)

		switch option {
		case "1":
		fmt.Println(" _____________ READ UNCOMMITTED _____________ ")

		case "2": 
		fmt.Println(" _____________ READ COMMITTED _____________ ")

		case "3": 
		fmt.Println(" _____________ REPEATABLE READ _____________ ")

		case "4":
		fmt.Println(" _____________ SERIALIZABLE _____________ ")
		var schedule string
		fmt.Print("Digite um Schedule: ")
		fmt.Scan(&schedule)

		scheduler := Scheduler.newScheduler(schedule)
		scheduler.SchedulerSerializable()
		fmt.Printf("Escalonamento final: " scheduler.ScheduleOut)

		case "5":
			fmt.Println("Saindo do programa...")
			return
		default:
			fmt.Println("Entrada inválida.")
			option = "0"
		} 
	}

func Schedule() {
	var schedule string
	fmt.Print("Digite um Schedule: ")
	fmt.Scan(&schedule)
}