package sao

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Stat_character struct {
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
	monnaie     int
}

type Weapon struct {
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

type Armor struct {
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
var input_axe int

// =============================================================================================================================================
func Create_Player() {
	var stat Stat_character
	fmt.Println("Bonjour player et bienvenu sur SAO ! Comment vous appeler vous ? ")
	fmt.Scanln(&input_player)
	stat.Pseudo = input_player
	stat.level = 1
	fmt.Println("Salutation", stat.Pseudo, "veuillez choisr votre classe d'arme : ", "\n")
	stat.Choose_Classe()
}

// =============================================================================================================================================
func (stat Stat_character) Choose_Classe() {
	var armor Armor
	fmt.Println("Vous avez les classes :", "\n")
	fmt.Println("   O    |----------------------------------------------------------------------------------------------------------------------------------|    O")
	fmt.Println("   ▒    | 1- archer : Peut utiliser l'arc ou l'arbalète comme arme, ne peut pas porte d'armure lourde.                                     |    ▒")
	fmt.Println("   ▒    | statistique de base :   PV_max = 50   //   AD = 15   //   Esquive = 8                                                            |    ▒")
	fmt.Println("o==▓==o |----------------------------------------------------------------------------------------------------------------------------------| o==▓==o")
	fmt.Println("   █    | 2- Lancier : Utilise la lance ou la halbarde comme arme, La lance permet de réduire l'armure en fonction de la stat Perfo_armor. |    █")
	fmt.Println("   █    | statistique de base :   PV_max = 80   //   AD = 13   //   Esquive = 4                                                            |    █")
	fmt.Println("   █    |----------------------------------------------------------------------------------------------------------------------------------|    █")
	fmt.Println("   █    | 3- épéiste : Peut utiliser comme arme la rapière, le katana, l'épée à une main et l'épée à deux main.                            |    █")
	fmt.Println("   ▼    | statistique de base :   PV_max = 100   //   AD = 10   //   Esquive = 6                                                           |    ▼")
	fmt.Println("  ynov  |----------------------------------------------------------------------------------------------------------------------------------|   ynov")
	fmt.Println("   ▲    | 4- hacheur : L'arme utilisée est la hache rapide ou hache lourde.                                                                |    ▲")
	fmt.Println("   █    | statistique de base :   PV_max = 120   //   AD = 11   //   Esquive = 3                                                           |    █")
	fmt.Println("   █    |----------------------------------------------------------------------------------------------------------------------------------|    █")
	fmt.Println("   █    | 5- Joueur de couteau : Utilisateur des dagues et couteau qui peut avoir des propriétées de saignement.                           |    █")
	fmt.Println("   █    | statistique de base :   Pv_max = 50   //   AD = 14   //   Esquive = 8                                                            |    █")
	fmt.Println("o==▓==o |----------------------------------------------------------------------------------------------------------------------------------| o==▓==o")
	fmt.Println("   ▒    | 6- masseur : Utilise et manie les masses et marteaux d'armes pouvant étourdir les ennemies.                                      |    ▒")
	fmt.Println("   ▒    | statistique de base :   PV_max = 130   //   AD = 9   //   Esquive = 1                                                            |    ▒")
	fmt.Println("   O    |----------------------------------------------------------------------------------------------------------------------------------|    O")
	fmt.Scanln(&input_classe)
	Clear()
	if input_classe == 1 {
		stat.Classe = "archer"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 15
		stat.Esquive = 8
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else if input_classe == 2 {
		stat.Classe = "lancier"
		stat.PV_max = 80
		stat.PV_actuelle = 80
		stat.AD = 13
		stat.Esquive = 4
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else if input_classe == 3 {
		stat.Classe = "épéiste"
		stat.PV_max = 100
		stat.PV_actuelle = 100
		stat.AD = 10
		stat.Esquive = 6
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else if input_classe == 4 {
		stat.Classe = "hacheur"
		stat.PV_max = 120
		stat.PV_actuelle = 120
		stat.AD = 11
		stat.Esquive = 2
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else if input_classe == 5 {
		stat.Classe = "joueur de couteau"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 14
		stat.Esquive = 8
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else if input_classe == 6 {
		stat.Classe = "masseur"
		stat.PV_max = 130
		stat.PV_actuelle = 130
		stat.AD = 9
		stat.Esquive = 1
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		stat.Equipement()
	} else {
		fmt.Println("Tu as pas rentrer le numéro de l'une des classe")
		stat.Choose_Classe()
	}
}

// =============================================================================================================================================
func (stat Stat_character) Fiche_perso() {
	fmt.Println("Name : ", stat.Pseudo, "                   ", "classe : ", stat.Classe, "        ", stat.equipement["tete"])
	fmt.Println("level : ", stat.level, "                        xp :", stat.xp, "                  ", stat.equipement["buste"])
	fmt.Println("                                                            ", stat.equipement["jambe"])
	fmt.Println("                                                            ", stat.equipement["botte"])
	fmt.Println("                                                         ", stat.equipement["main gauche"], "   ", stat.equipement["main droite"])

	fmt.Println("_______________________________________________STATISTIQUE_______________________________________________")
	fmt.Println("PV max : ", stat.PV_max, "                 ", "PV actuelle : ", stat.PV_actuelle)
	fmt.Println("AD : ", stat.AD, "                         ", "Esquive : ", stat.Esquive)
	fmt.Println("Point de stat : ", stat.point_stat)
	fmt.Println("________________________________________________ABILITIE_________________________________________________")
	fmt.Println(" ")
	fmt.Println("________________________________________________INVENTORY________________________________________________")
	fmt.Println(stat.inventory)
	fmt.Println("gold : ", stat.monnaie)
	time.Sleep(3 * time.Second)
	stat.Quete()
}

// =============================================================================================================================================
func (stat Stat_character) Equipement() {
	var weapon Weapon
	if input_classe == 1 {
		weapon.bow = "arc de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": " ", "main gauche": weapon.bow}
		stat.Fiche_perso()
	}
	if input_classe == 2 {
		weapon.spear = "lance de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.spear, "main gauche": "X"}
		stat.Fiche_perso()
	}
	if input_classe == 3 {
		stat.Voix_Epee()
	}

	if input_classe == 4 {
		Clear()
		stat.Voix_hache()
	}
	if input_classe == 5 {
		weapon.knife = "dague de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.knife, "main gauche": " "}
		stat.Fiche_perso()
	}
	if input_classe == 6 {
		weapon.mace = "masse de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.mace, "main gauche": " "}
		stat.Fiche_perso()
	}
}

// =============================================================================================================================================
func (stat *Stat_character) Voix_Epee() {
	fmt.Println("o∷∷{=============▻ Choisiser votre voix de l'épée <ΞΞΞΞΞΞΞΞΞ{o}≠≠≠O: , ,  et ")
	fmt.Println("           ⸸-----------------------Cyril------------------⸸")
	fmt.Println("           |               1 : la rapière                 |")
	fmt.Println("           ⸸----------------------Data-IA-----------------⸸")
	fmt.Println("  ")
	fmt.Println("           🗡----------------------Ethan------------------🗡")
	fmt.Println("           |               2 : le katana                  |")
	fmt.Println("           🗡--------------------Game-Prod----------------🗡")
	fmt.Println("   ")
	fmt.Println("           ⚔️---------------------Kheir------------------⚔️")
	fmt.Println("           |               3 : l'épée à une main          |")
	fmt.Println("           ⚔️--------------------Data-IA-----------------⚔️")
	fmt.Println("  ")
	fmt.Println("           🗡️--------------------Allan-------------------🗡️")
	fmt.Println("           |               4 : l'épée à deux main         |")
	fmt.Println("           🗡️---------------------Web--------------------🗡️")
	fmt.Scanln(&input_epee)
	var weapon Weapon
	if input_epee == 1 {
		weapon.rapier = "rapier de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.rapier, "main gauche": " "}
		Clear()
		stat.Fiche_perso()
	} else if input_epee == 2 {
		weapon.katana = "katana de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.katana, "main gauche": " "}
		Clear()
		stat.Fiche_perso()
	} else if input_epee == 3 {
		weapon.sword = "épée à une main de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.sword, "main gauche": " "}
		Clear()
		stat.Fiche_perso()
	} else if input_epee == 4 {
		weapon.longsword = "épée à deux main de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.longsword, "main gauche": "X"}
		Clear()
		stat.Fiche_perso()
	} else {
		fmt.Println("Rentrer l'un des numéros corespondant à une voix de l'épée")
		Clear()
		stat.Voix_Epee()
	}
}

func (stat *Stat_character) Voix_hache() {
	var weapon Weapon
	fmt.Println("Pour chosir votre voix du maniement de la hache, entrer 1 pour la hache lourde et 2 pour la hache rapide : ")
	fmt.Scanln(&input_axe)
	if input_axe == 1 {
		weapon.axe = "hache lourde de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.axe, "main gauche": " "}
		stat.Fiche_perso()
	} else if input_axe == 2 {
		weapon.axe = "hache rapide de départ"
		stat.equipement = map[string]string{"tete": "casque de départ", "buste": "plastron de départ", "jambe": "pantalon de départ", "botte": "botte de départ", "main droite": weapon.axe, "main gauche": weapon.axe}
		stat.Fiche_perso()
	} else {
		fmt.Println("ERREUR")
		stat.Voix_hache()
	}
}

// =============================================================================================================================================================
func (stat Stat_character) Quete() {
	var string_quest int
	Clear()
	fmt.Println("Maintenant aventurier ", stat.Pseudo, " veuilliez choisir une Quête, si vous voulez partir sans prendre de quête dite 0. ", "\n")
	fmt.Println("[]====================tableau=des=quêtes====================[]")
	fmt.Println("∐               1 -Donjons de l'arche perdu- (D)            ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("[]==========================================================[]", "\n")
	fmt.Scanln(&string_quest)
	if string_quest == 1 {
		fmt.Println("Bonne chance pour votre quête !")
		fmt.Println("Mais je vous conseille de vous entrainer sur le terrain d'entrainement au combat avant.")
		fmt.Println("Vous pouvez aussi visiter la ville si vous le voulez aussi. ")
		Clear()
		fmt.Println("Sur ce bonne chance pour la conquête de 'Sword Art Online' !")
	} else if string_quest == 0 {
		fmt.Println("Aurevoir et penser à venir ici pour des quêtes de grandes envergure !")
	} else {
		fmt.Println("Faite un choix sinon l'admin sera pas contente !")
		stat.Quete()
	}
}

// =============================================================================================================================================
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
