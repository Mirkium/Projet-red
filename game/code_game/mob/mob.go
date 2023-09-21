package mob


type mob1 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMob1() {
	var mob1 mob1
	mob1.nom = "apprenti chevalier"
	mob1.classe = "Traitre"
	mob1.niveau = 1
	mob1.pointsDeVieMax = 50
	mob1.pointsDeVieActuels = 50
	mob1.inventaire = map[string]int{"épée fragile": 1, "casque": 1}
}

type mob2 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMob8() {
	var mob2 mob2
	mob2.nom = "chevalier déchu"
	mob2.classe = "Traitre"
	mob2.niveau = 5
	mob2.pointsDeVieMax = 75
	mob2.pointsDeVieActuels = 75
	mob2.inventaire = map[string]int{"épée en pierre": 1, "casque": 1, "bouclier": 1}
}

type mob3 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMob3() {
	var mob3 mob3
	mob3.nom = "sergent a double face"
	mob3.classe = "Traitre"
	mob3.niveau = 10
	mob3.pointsDeVieMax = 100
	mob3.pointsDeVieActuels = 100
	mob3.inventaire = map[string]int{"Lance": 1, "casque": 1, "bouclier": 1}
}

type mob4 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMob14() {
	var mob4 mob4
	mob4.nom = "Capitaine mort"
	mob4.classe = "Traitre"
	mob4.niveau = 15
	mob4.pointsDeVieMax = 125
	mob4.pointsDeVieActuels = 125
	mob4.inventaire = map[string]int{"épée en fer": 1, "casque": 1, "bouclier": 1}
}

type mob5 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
}

func initMob5() {
	var mob5 mob5
	mob5.nom = "roi de l'ombre"
	mob5.classe = "Traitre"
	mob5.niveau = 20
	mob5.pointsDeVieMax = 150
	mob5.pointsDeVieActuels = 150
	mob5.inventaire = map[string]int{"épée en obsidienne": 2, "Armure complète": 1, "potion": 1}
}
