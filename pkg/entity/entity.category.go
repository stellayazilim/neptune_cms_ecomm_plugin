package entity

type Category struct {
	ID          uint16
	Name        string
	Description string
	Parent      *Category
}
