package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/WilliamKotoco/QuizMeDeutsch/internal/datenbetreiber"
)

func HauptfensterErstellen() {
	var typ datenbetreiber.DatenbetreiberTyp
	var vall string

	/// the main window must have a option for user to choose how they want the pair
	/// question-answer to be generated; a label showing the question, text field for
	/// writing text and a button to correct the answer.
	app_ := app.New()
	fenster := app_.NewWindow("Fragefenster")
	entry := widget.NewLabel("Wählen Sie den Typ des Fragen generators")
	combo := widget.NewSelect([]string{"zufällig", "nummer", "reichweite"}, func(value string) {
		if value == "zufällig" {
			typ = datenbetreiber.ZUFFALIG
		} else if value == "nummer" {
			typ = datenbetreiber.NUMMER
		} else {
			typ = datenbetreiber.REICHWEITE
		}
		vall = value

	})
	spacer := layout.NewSpacer()
	content := widget.NewButton("ok", func() {
		fenster.Close()
	})
	fenster.SetContent(container.NewVBox(entry, combo, spacer, content))
	fenster.Resize(fyne.NewSize(300, 300))
	fenster.SetFixedSize(true)
	fenster.ShowAndRun()
	fmt.Printf("typ " + vall + " " + string(typ))
}

// / utils
