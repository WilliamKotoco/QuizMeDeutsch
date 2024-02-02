package main

import (
	"fmt"
	"github.com/WilliamKotoco/QuizMeDeutsch/internal/datenbetreiber"
)

func main() {

	var test datenbetreiber.Datenbetreiber

	test = &datenbetreiber.ZufälligDatenbetreiber{
		Frage:   "",
		Antwort: "",
	}

	test.LadenFrageAntwort()

	zufälligDatenbetreiber := test.(*datenbetreiber.ZufälligDatenbetreiber)

	fmt.Println(zufälligDatenbetreiber.Frage)
	fmt.Println(zufälligDatenbetreiber.Antwort)
}
