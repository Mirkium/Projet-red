	fmt.Printf("%s entre en combat avec %s!\n", player.Name, mob.nom)

	// Boucle de combat
	for player.Health > 0 && mob.Health > 0 {
		// Joueur attaque le monstre
		playerDamage := rand.Intn(player.Attack)
		mob.Health -= playerDamage
		fmt.Printf("%s inflige %d points de dégâts à %s!\n", player.Name, playerDamage, mob.Name)

		// Vérifier si le monstre est vaincu
		if mob.Health <= 0 {
			fmt.Printf("%s a vaincu le monstre!\n", player.Name)
			break
		}

		// Monstre attaque le joueur
		mobDamage := rand.Intn(mob.Attack)
		player.Health -= mobDamage
		fmt.Printf("%s inflige %d points de dégâts à %s!\n", monster.Name, mobDamage, player.Name)

		// Vérifier si le joueur est vaincu
		if player.Health <= 0 {
			fmt.Printf("%s a été vaincu par le monstre!\n", player.Name)
			break
		}

		// Pause avant le prochain tour
		time.Sleep(time.Second)
	}
}
