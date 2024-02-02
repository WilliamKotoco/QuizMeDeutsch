package datenbetreiber

import (
	"github.com/WilliamKotoco/QuizMeDeutsch/pkg"
	"io"
	"log"
	"os"
	"strings"
)

const ANZAHL_ZEILEN int = 9275
const DATEINAME string = "assets/data.txt"
const OFFSET int = 600 //Größe des größten string in der Datei

// Datenbetreiber / Schnittstelle, die angibt, wie ein Datenbetreiber mit Daten umgehen muss.
type Datenbetreiber interface {

	// FrageAntwortLaden / Diese Funktion öffnet die Datei und füllt die Struktur mit dem Frage-Antwort-Paar aus dem Datensatz.

	LadenFrageAntwort()
}

// LeseZeile gibt ein Frage-Antwort-Paar zurück, nachdem die Zeile aus der Datei gelesen und formatiert wurde.
func LeseZeile(zeilennummer int, datei *os.File) pkg.Paar[string, string] {

	datei_distanz := (zeilennummer - 1) * OFFSET

	_, err := datei.Seek(int64(datei_distanz), 0)

	if err != nil {
		log.Fatal("Line 31", err)
	}

	var zeile string
	// Read 601 characters from the line
	buffer := make([]byte, OFFSET)
	n, err := io.ReadFull(datei, buffer)

	zeile = string(buffer[:n])
	var paar pkg.Paar[string, string]

	paar.Erste, paar.Zweite = formatierungszeile(zeile)

	return paar
}

// Formatieren Sie die Zeile, indem zwei Zeichenfolgen durch die Trennzeichen „#“
// und „!“ getrennt werden, um eine Frage und eine Antwort zu erstellen.
func formatierungszeile(zeile string) (string, string) {
	zeile_slice := strings.Split(zeile, "#")

	return zeile_slice[0], strings.Split(zeile_slice[1], "!")[0]

}
