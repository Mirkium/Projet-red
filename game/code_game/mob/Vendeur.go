type marchant struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMarchant() {
	var marchant marchant
	marchant.nom = "Le Forgeron"
	marchant.classe = "Vendeur"
	marchant.niveau = 50
	marchant.pointsDeVieMax = 10000
	marchant.pointsDeVieActuels = 10000
	marchant.inventaire = map[string]int{"potion": 100, "épée": 50}
}

type initStrMarchant2 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMarchant2() {
	var marchand2 Marchand
	marchand2.nom = "Entraineur"
	marchand2.classe = "Vendeur"
	marchand2.niveau = 100
	marchand2.pointsDeVieMax = 10000
	marchand2.pointsDeVieActuels = 10000
	marchand2.inventaire = map[string]int{"AD": 1, "Esquive": 1, "PV": 5}
}

type initStrMarchant3 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMarchant3() {
	var marchand2 Marchand
	marchand2.nom = "Alchimiste"
	marchand2.classe = "Vendeur"
	marchand2.niveau = 100
	marchand2.pointsDeVieMax = 10000
	marchand2.pointsDeVieActuels = 10000
	marchand2.inventaire = map[string]int{"potion de soin": 5, "potion de poison": 5}
}
