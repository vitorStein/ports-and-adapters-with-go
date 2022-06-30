package app

import "errors"

type BookInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

const (
	DISABLED = "disabled"
	ENABLE   = "enable"
)

type Book struct {
	ID     string
	Name   string
	Price  float32
	Status string
}

/* func (b *Book) IsValid() (bool, error) {

} */

func (b *Book) Enable() error {
	if b.Price > 0 {
		b.Status = ENABLE
		return nil
	}
	return errors.New("o preço do livro deve ser maior que zero para habilitar o livro")
}

func (b *Book) Disable() error {

	if b.Price == 0 {
		b.Status = DISABLED
		return nil
	}
	return errors.New("o preço do livro deve ser zero para ser desabilitado")
}

func (b *Book) GetID() string {
	return b.ID
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetStatus() string {
	return b.Status
}

func (b *Book) GetPrice() float32 {
	return b.Price
}
