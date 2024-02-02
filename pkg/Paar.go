package pkg

// Paar /Paar diese Struktur ermöglicht das Speichern eines Objektpaars in einer einzigen Struktur
type Paar[T, U any] struct {
	Erste  T
	Zweite U
}
