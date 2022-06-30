package app

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

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
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float32 `valid:"float, optional"`
	Status string  `valid:"required"`
}

func (b *Book) IsValid() (bool, error) {

	if b.Status == "" {
		b.Status = DISABLED
	}

	if b.Status != ENABLE && b.Status != DISABLED {
		return false, errors.New("o status deve ser ENABLE ou DISABLED")
	}

	if b.Price < 0 {
		return false, errors.New("o preço deve ser maior ou igual a 0")
	}

	_, err := govalidator.ValidateStruct(b)

	if err != nil {
		return false, err
	}

	return true, nil
}

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
