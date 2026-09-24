package main
import "fmt"

type Basket struct{
	name string
	color string
	content []string
}
type Clothing struct{
	name string
	color string
	state bool // false = sale true = propre
}

func displayBasket(b *Basket){
	if len(b.content)<1{
		fmt.Print("panier vide")
	}else{
		fmt.Printf("==paniere %s==",b.name)
		for i := 0; i < len(b.content); i++{
			fmt.Print(b.content[i].name)
			fmt.Print(b.content[i].color)
			if b.content[i].state == false {
				fmt.Print("sale")
			}else{
				fmt.Print("prore")
			}
			
		}
	}
}
func addToBasket (c *Clothing, b *Basket){
	if c.color != b.color {
		fmt.Println("Oupss ! Ajout impossible la couleur du vêtement ne correspond pas à la corbeille")
	}else{
		fmt.Printf("Ajout d'un nouveau vêtement :%s dans la %s",c.name,b.name)
		b.content = append(b.content, c)
	}
}

func cleanBasket(b *Basket) {
	if len(b.content) == 0 {
		fmt.Println("Erreur : le panier est vide, rien à laver")
		return
	}
	count := 0
	for i := range b.content {
		if !b.content[i].state {
			b.content[i].state = true
			count++
		}
	}
	fmt.Printf("==Nettoyage du panier %s ==", b.name)
	fmt.Printf("Le lavage du panier %s est terminé, %d vêtements ont été lavés", b.name, count)
}

func emptyCleanLaundry() {
	
}
