package entity

type Product struct {
	ID         uint64
	Name       string
	Brand      string
	Categories []Category
}
