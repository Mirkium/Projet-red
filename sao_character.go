package main

import "fmt"

type stat_character struct {
	Pseudo      string
	Classe      string
	level       int
	PV_max      int
	PV_actuelle int
	AD          int
	Esquive     int
	point_stat  int
	inventory   map[string]int
}

func Create_Player() {
	var stat stat_character
	var input string
	fmt.Println("Bonjour player et bienvenu sur SAO ! Comment vous appeler vous ? ")
	fmt.Scanln(&input)
	stat.Pseudo = input
	stat.level = 1
	fmt.Println("Salutation", stat.Pseudo, "veuillez choisr votre classe d'arme : ", "\n")
	stat.Choose_Classe()
}

func (stat stat_character) Choose_Classe() {
	var input int
	fmt.Println("Vous avez les classes :", "\n", "1- archer : Peut utiliser l'arc ou l'arbalète comme arme, ne peut pas porte d'armure lourde.", "\n", "PV_max = 50", "\n", "AD = 15", "\n", "Esquive = 8", "\n")
	fmt.Println("2- Lancier : Utilise la lance ou la halbarde comme arme, La lance permet de réduire l'armure en fonction de la stat Perfo_armor.", "\n", "PV_max = 80", "\n", "AD = 13", "\n", "Esquive = 4", "\n")
	fmt.Println("3- épéiste : Peut utiliser comme arme la rapière, le katana, l'épée à une main et l'épée à deux main.", "\n", "PV_max = 100", "\n", "AD = 10", "\n", "Esquive = 6", "\n")
	fmt.Println("4- hacheur : L'arme utilisée est la hache rapide ou hache lourde.", "\n", "PV_max = 120", "\n", "AD = 11", "\n", "Esquive = 3", "\n")
	fmt.Println("5- Joueur de couteau : Utilisateur des dagues et couteau qui peut avoir des propriétées de saignement.", "\n", "Pv_max = 50", "\n", "AD = 14", "\n", "Esquive = 8", "\n")
	fmt.Println("6- masseur : Utilise et manie les masses et marteaux d'armes pouvant étourdir les ennemies.", "\n", "PV_max = 130", "\n", "AD = 9", "\n", "Esquive = 1", "\n")
	fmt.Println("Pour choisir une classe entrer le numéro de la classe : ")
	fmt.Scanln(&input)
	if input == 1 {
		stat.Classe = "archer"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 15
		stat.Esquive = 8
		stat.Fiche_perso()
	}
	if input == 2 {
		stat.Classe = "lancier"
		stat.PV_max = 80
		stat.PV_actuelle = 80
		stat.AD = 13
		stat.Esquive = 4
		stat.Fiche_perso()
	}
	if input == 3 {
		stat.Classe = "épéiste"
		stat.PV_max = 100
		stat.PV_actuelle = 100
		stat.AD = 10
		stat.Esquive = 6
		stat.Fiche_perso()
	}
	if input == 4 {
		stat.Classe = "hacheur"
		stat.PV_max = 120
		stat.PV_actuelle = 120
		stat.AD = 11
		stat.Esquive = 2
		stat.Fiche_perso()
	}
	if input == 5 {
		stat.Classe = "joueur de couteau"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 14
		stat.Esquive = 8
		stat.Fiche_perso()
	}
	if input == 6 {
		stat.Classe = "masseur"
		stat.PV_max = 130
		stat.PV_actuelle = 130
		stat.AD = 9
		stat.Esquive = 1
		stat.Fiche_perso()
	} else if input > 6 {
		fmt.Println("Tu as pas rentrer le numéro de l'une des classe")
		stat.Choose_Classe()
	}
}

func (stat stat_character) Fiche_perso() {
	fmt.Println("Name : ", stat.Pseudo, "                   ", "classe : ", stat.Classe)
	fmt.Println("level : ", stat.level)
	fmt.Println("___________________________STATISTIQUE___________________________")
	fmt.Println("PV max : ", stat.PV_max, "                 ", "PV actuelle : ", stat.PV_actuelle)
	fmt.Println("AD : ", stat.AD, "                         ", "Esquive : ", stat.Esquive)
	fmt.Println("Point de stat : ", stat.point_stat)
	fmt.Println("____________________________ABILITIE___________________________")
	fmt.Println(" ")
	fmt.Println("___________________________INVENTORY___________________________")
	fmt.Println(stat.inventory)
}

func main() {
	Create_Player()
}
