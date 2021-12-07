package main

import "fmt"

func main() {

	//dentro do colchetes é o tipo das chaves
	//depois é o tipo dos valores
	//ambos devem ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
		"cidade":    "Alfenas",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
