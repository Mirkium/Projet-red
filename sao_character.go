package sao

import "fmt"

type stat_character struct {
	Pseudo      string
	Classe      string
	level       int
	xp          int
	PV_max      int
	PV_actuelle int
	AD          int
	Esquive     int
	point_stat  int
	inventory   map[string]int
	equipement  map[string]string
}

type weapon struct {
	bow               string
	spear             string
	rapier            string
	katana            string
	sword             string
	longsword         string
	knife             string
	axe               string
	mace              string
	durability_weapon int
	stat_bonus        string
	stat_leave        string
}

type armor struct {
	head_armor       string
	chest_armor      string
	leg_armor        string
	foot_armor       string
	durability_heat  int
	durability_chest int
	durability_leg   int
	durability_foot  int
	stat_bonus       string
	stat_leave       string
}

var input_classe int
var input_player string
var input_epee int

func Create_Player() {
	var stat stat_character
	fmt.Println("Bonjour player et bienvenu sur SAO ! Comment vous appeler vous ? ")
	fmt.Scanln(&input_player)
	stat.Pseudo = input_player
	stat.level = 1
	fmt.Println("Salutation", stat.Pseudo, "veuillez choisr votre classe d'arme : ", "\n")
	stat.Choose_Classe()
}

func (stat stat_character) Choose_Classe() {
	var armor armor
	fmt.Println("Vous avez les classes :", "\n", "1- archer : Peut utiliser l'arc ou l'arbalète comme arme, ne peut pas porte d'armure lourde.", "\n", "PV_max = 50", "\n", "AD = 15", "\n", "Esquive = 8", "\n")
	fmt.Println("2- Lancier : Utilise la lance ou la halbarde comme arme, La lance permet de réduire l'armure en fonction de la stat Perfo_armor.", "\n", "PV_max = 80", "\n", "AD = 13", "\n", "Esquive = 4", "\n")
	fmt.Println("3- épéiste : Peut utiliser comme arme la rapière, le katana, l'épée à une main et l'épée à deux main.", "\n", "PV_max = 100", "\n", "AD = 10", "\n", "Esquive = 6", "\n")
	fmt.Println("4- hacheur : L'arme utilisée est la hache rapide ou hache lourde.", "\n", "PV_max = 120", "\n", "AD = 11", "\n", "Esquive = 3", "\n")
	fmt.Println("5- Joueur de couteau : Utilisateur des dagues et couteau qui peut avoir des propriétées de saignement.", "\n", "Pv_max = 50", "\n", "AD = 14", "\n", "Esquive = 8", "\n")
	fmt.Println("6- masseur : Utilise et manie les masses et marteaux d'armes pouvant étourdir les ennemies.", "\n", "PV_max = 130", "\n", "AD = 9", "\n", "Esquive = 1", "\n")
	fmt.Println("Pour choisir une classe entrer le numéro de la classe : ")
	fmt.Scanln(&input_classe)
	if input_classe == 1 {
		stat.Classe = "archer"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 15
		stat.Esquive = 8
		stat.Equipement()
	}
	if input_classe == 2 {
		stat.Classe = "lancier"
		stat.PV_max = 80
		stat.PV_actuelle = 80
		stat.AD = 13
		stat.Esquive = 4
		stat.Equipement()
	}
	if input_classe == 3 {
		stat.Classe = "épéiste"
		stat.PV_max = 100
		stat.PV_actuelle = 100
		stat.AD = 10
		stat.Esquive = 6
		stat.Equipement()
	}
	if input_classe == 4 {
		stat.Classe = "hacheur"
		stat.PV_max = 120
		stat.PV_actuelle = 120
		stat.AD = 11
		stat.Esquive = 2
		stat.Equipement()
	}
	if input_classe == 5 {
		stat.Classe = "joueur de couteau"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 14
		stat.Esquive = 8
		stat.Equipement()
	}
	if input_classe == 6 {
		stat.Classe = "masseur"
		stat.PV_max = 130
		stat.PV_actuelle = 130
		stat.AD = 9
		stat.Esquive = 1
		stat.Equipement()
	} else if input_classe > 6 {
		fmt.Println("Tu as pas rentrer le numéro de l'une des classe")
		stat.Choose_Classe()
	}
	stat.xp = 0
	armor.head_armor = "casque de départ"
	armor.chest_armor = "buste de départ"
	armor.leg_armor = "pantalon de départ"
	armor.foot_armor = "botte de départ"
	armor.durability_heat = 40
	armor.durability_chest = 70
	armor.durability_leg = 60
	armor.durability_foot = 50
}

func (stat stat_character) Fiche_perso() {
	fmt.Println("Name : ", stat.Pseudo, "                   ", "classe : ", stat.Classe, "        ", stat.equipement["tete"])
	fmt.Println("level : ", stat.level, "                        xp :", stat.xp, "                  ", stat.equipement["buste"])
	fmt.Println("                                                            ", stat.equipement["jambe"])
	fmt.Println("                                                            ", stat.equipement["botte"])
	fmt.Println("                                                         ", stat.equipement["main gauche"], "   ", stat.equipement["main droite"])

	fmt.Println("___________________________STATISTIQUE___________________________")
	fmt.Println("PV max : ", stat.PV_max, "                 ", "PV actuelle : ", stat.PV_actuelle)
	fmt.Println("AD : ", stat.AD, "                         ", "Esquive : ", stat.Esquive)
	fmt.Println("Point de stat : ", stat.point_stat)
	fmt.Println("____________________________ABILITIE___________________________")
	fmt.Println(" ")
	fmt.Println("___________________________INVENTORY___________________________")
	fmt.Println(stat.inventory)
}

func (stat stat_character) Equipement() {
	var armor armor
	var weapon weapon
	var input_axe int
	if input_classe == 1 {
		weapon.bow = "arc de départ"
		stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": " ", "main gauche": weapon.bow}
		stat.Fiche_perso()
	}
	if input_classe == 2 {
		weapon.spear = "lance de départ"
		stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.spear, "main gauche": "X"}
		stat.Fiche_perso()
	}
	if input_classe == 3 {
		fmt.Println("Pour choisir votre voix de l'épée entrée : 1 pour la rapière, 2 pour le katana, 3pour l'épée à une main et 4 pour l'épée à deux main")
		fmt.Scanln(&input_epee)

		if input_epee == 1 {
			weapon.rapier = "rapier de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.rapier, "main gauche": " "}
			stat.Fiche_perso()
		}
		if input_epee == 2 {
			weapon.katana = "katana de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.katana, "main gauche": " "}
			stat.Fiche_perso()
		}
		if input_epee == 3 {
			weapon.sword = "épée à une main de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.sword, "main gauche": " "}
			stat.Fiche_perso()
		}
		if input_epee == 4 {
			weapon.longsword = "épée à deux main de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.longsword, "main gauche": "X"}
			stat.Fiche_perso()
		}
	}
	if input_classe == 4 {
		fmt.Println("Pour chosir votre voix du maniement de la hache, entrer 1 pour la hache lourde et 2 pour la hache rapide : ")
		fmt.Scanln(&input_axe)
		if input_axe == 1 {
			weapon.axe = "hache lourde de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.axe, "main gauche": " "}
			stat.Fiche_perso()
		}
		if input_axe == 2 {
			weapon.axe = "hache rapide de départ"
			stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.axe, "main gauche": weapon.axe}
			stat.Fiche_perso()
		}
	}
	if input_classe == 5 {
		weapon.knife = "dague de départ"
		stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.knife, "main gauche": " "}
		stat.Fiche_perso()
	}
	if input_classe == 6 {
		weapon.mace = "masse de départ"
		stat.equipement = map[string]string{"tete": armor.head_armor, "buste": armor.chest_armor, "jambe": armor.leg_armor, "botte": armor.foot_armor, "main droite": weapon.mace, "main gauche": " "}
		stat.Fiche_perso()
	}
}
