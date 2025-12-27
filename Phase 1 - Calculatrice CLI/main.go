package main

import (
	"fmt"
	"time"
)

func main() {
	var menuchoice float64
	sleep := time.Duration(2) * time.Second
	var videur string //permet de vider le bufer
	for {
		clearscreen()
		fmt.Println("Bonjour et bienvenue dans votre Calculatrice préférée !")
		fmt.Println("Quelle opération souhaitez-vous effectuer aujourd'hui ?")
		fmt.Println("1 - Une Addition")
		fmt.Println("2 - Une Soustraction")
		fmt.Println("3 - Une Multiplication")
		fmt.Println("4 - Une Division")
		fmt.Println("5 - Quittez")
		_, err := fmt.Scan(&menuchoice)
		if err != nil {
			fmt.Println("❌ Choix impossible, réessayez.")
			fmt.Scanln(&videur)
			time.Sleep(sleep)
			continue
		}
		switch menuchoice {
		case 1:
		addition:
			for {
				clearscreen()
				var a float64
				var b float64
				fmt.Println("Choississez les nombres que vous voulez additionner :")
				fmt.Println("Choississez le premier nombre :")
				_, erra := fmt.Scan(&a)
				if erra != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Choississez les nombres que vous voulez additionner :")
				fmt.Println("Choississez le deuxième nombre :")
				_, errb := fmt.Scan(&b)
				if errb != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("La somme des nombres que vous avez choisis est :", addition(a, b))
			choix1:
				for {
					fmt.Println("Que voulez vous faire maintenant ?")
					fmt.Println("1 - Continuez à additionner des nombres")
					fmt.Println("0 - Revenir au menu des opérations")
					_, err := fmt.Scan(&menuchoice)
					if err != nil {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					switch menuchoice {
					case 0:
						break addition
					case 1:
						break choix1
					}
				}
			}
		case 2:
		soustraction:
			for {
				clearscreen()
				var a float64
				var b float64
				fmt.Println("Choississez les nombres que vous voulez soustraire :")
				fmt.Println("Choississez le premier nombre :")
				_, erra := fmt.Scan(&a)
				if erra != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Choississez les nombres que vous voulez soustraire :")
				fmt.Println("Choississez le deuxième nombre :")
				_, errb := fmt.Scan(&b)
				if errb != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("La différence entre les nombres que vous avez choisis est :", soustraction(a, b))
			choix2:
				for {
					fmt.Println("Que voulez vous faire maintenant ?")
					fmt.Println("1 - Continuez à soustraire des nombres")
					fmt.Println("0 - Revenir au menu des opérations")
					_, err := fmt.Scan(&menuchoice)
					if err != nil {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					switch menuchoice {
					case 0:
						break soustraction
					case 1:
						break choix2
					}
				}
			}
		case 3:
		multiplication:
			for {
				clearscreen()
				var a float64
				var b float64
				fmt.Println("Choississez les nombres que vous voulez multiplier :")
				fmt.Println("Choississez le premier nombre :")
				_, erra := fmt.Scan(&a)
				if erra != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Choississez les nombres que vous voulez multiplier :")
				fmt.Println("Choississez le deuxième nombre :")
				_, errb := fmt.Scan(&b)
				if errb != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Le produit des nombres que vous avez choisis est :", multiplication(a, b))
			choix3:
				for {
					fmt.Println("Que voulez vous faire maintenant ?")
					fmt.Println("1 - Continuez à multiplier des nombres")
					fmt.Println("0 - Revenir au menu des opérations")
					_, err := fmt.Scan(&menuchoice)
					if err != nil {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					switch menuchoice {
					case 0:
						break multiplication
					case 1:
						break choix3
					}
				}
			}
		case 4:
		division:
			for {
				clearscreen()
				var a float64
				var b float64
				fmt.Println("Choississez les nombres que vous voulez diviser :")
				fmt.Println("Choississez le premier nombre :")
				_, erra := fmt.Scan(&a)
				if erra != nil {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Choississez les nombres que vous voulez diviser :")
				fmt.Println("Choississez le deuxième nombre :")
				_, errb := fmt.Scan(&b)
				if errb != nil || b == 0 {
					fmt.Println("❌ Choix impossible, réessayez. (La division par 0 est impossible)")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				fmt.Println("Le quotient des nombres que vous avez choisis est :", division(a, b))
			choix4:
				for {
					fmt.Println("Que voulez vous faire maintenant ?")
					fmt.Println("1 - Continuez à diviser des nombres")
					fmt.Println("0 - Revenir au menu des opérations")
					_, err := fmt.Scan(&menuchoice)
					if err != nil {
						fmt.Println("❌ Choix impossible, réessayez.")
						fmt.Scanln(&videur)
						time.Sleep(sleep)
						continue
					}
					switch menuchoice {
					case 0:
						break division
					case 1:
						break choix4
					}
				}
			}
		case 5:
			return
		default:
			clearscreen()
			fmt.Println("❌ Choix impossible, réessayez.")
			fmt.Scanln(&videur)
		}

	}
}

func addition(a, b float64) (somme float64) {
	somme = a + b
	return somme
}

func soustraction(a, b float64) (difference float64) {
	difference = a - b
	return difference
}

func multiplication(a, b float64) (produit float64) {
	produit = a * b
	return produit
}

func division(a, b float64) (quotient float64) {
	quotient = a / b
	return quotient
}

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}
