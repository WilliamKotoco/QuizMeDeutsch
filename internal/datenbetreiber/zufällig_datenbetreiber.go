package datenbetreiber

import (
	"github.com/WilliamKotoco/QuizMeDeutsch/pkg"
	"log"
	"math/rand"
	"os"
)

// ZufälligDatenbetreiber ist die struktur inder zufällig ausgewählte fragen und
// antwort aus der Datei gespeichert werden.
type ZufälligDatenbetreiber struct {
	Frage   string
	Antwort string
}

// LadenFrageAntwort / diese Funktion wählt zufällig eine Zeile aus dem Datensatz, die eine Frage und eine Antwort
// enthält
func (zd *ZufälligDatenbetreiber) LadenFrageAntwort() {
	datei, fehler := os.Open(DATEINAME)

	if fehler != nil {
		log.Fatal("Fehler beim Lesen der Datei")
	}

	defer datei.Close()

	linie := rand.Intn(ANZAHL_ZEILEN + 1)

	var paar pkg.Paar[string, string] = LeseZeile(linie, datei)

	zd.Frage = paar.Erste
	zd.Antwort = paar.Zweite
}
