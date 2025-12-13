package main

import "fmt"

func main() {
	var Menuchoice int
	for {
		clearscreen()
		fmt.Println("Bonjour et bienvenue dans votre Calculatrice préférée !")
		fmt.Println("Quelle opération souhaitez-vous effectuer aujourd'hui ?")
		fmt.Println("1 - Une Addition")
		fmt.Println("2 - Une Soustraction")
		fmt.Println("3 - Une Multiplication")
		fmt.Println("4 - Une Division")
		fmt.Println("5 - Quittez")
		fmt.Scan(&Menuchoice)
		switch Menuchoice {
		case 1:
			for {
				clearscreen()
				var a float64
				var b float64
				fmt.Println("Choississez les nombres que vous voulez additionner :")
				fmt.Println("Choississez le premier nombre :")
				fmt.Scan(&a)
				fmt.Println("Choississez les nombres que vous voulez additionner :")
				fmt.Println("Choississez le deuxième nombre :")
				fmt.Scan(&b)
				fmt.Println("La somme des nombres que vous avez choisis est :", addition(a, b))
				for {
					fmt.Println("Que voulez vous faire maintenant ?")
					fmt.Println("1 - Continuez à additionner des nombres")
					fmt.Println("0 - Revenir au menu des opérations")
					fmt.Scan(&Menuchoice)
					switch Menuchoice {
					case 1:
					case 2:
					}
				}
			}
		case 2:
			clearscreen()
			var a float64
			var b float64
			fmt.Println("Choississez les nombres que vous voulez soustraire :")
			fmt.Println("Choississez le premier nombre :")
			fmt.Scan(&a)
			fmt.Println("Choississez les nombres que vous voulez soustraire :")
			fmt.Println("Choississez le deuxième nombre :")
			fmt.Scan(&b)
			fmt.Println("La différence entre les nombres que vous avez choisis est :", soustraction(a, b))
		case 3:
			clearscreen()
			var a float64
			var b float64
			fmt.Println("Choississez les nombres que vous voulez multiplier :")
			fmt.Println("Choississez le premier nombre :")
			fmt.Scan(&a)
			fmt.Println("Choississez les nombres que vous voulez multiplier :")
			fmt.Println("Choississez le deuxième nombre :")
			fmt.Scan(&b)
			fmt.Println("Le produit des nombres que vous avez choisis est :", multiplication(a, b))
		case 4:
			clearscreen()
			var a float64
			var b float64
			fmt.Println("Choississez les nombres que vous voulez diviser :")
			fmt.Println("Choississez le premier nombre :")
			fmt.Scan(&a)
			fmt.Println("Choississez les nombres que vous voulez diviser :")
			fmt.Println("Choississez le deuxième nombre :")
			fmt.Scan(&b)
			fmt.Println("Le quotient des nombres que vous avez choisis est :", division(a, b))
		case 5:
			return
		default:
			clearscreen()
			fmt.Println("Mauvaise entrée")
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
