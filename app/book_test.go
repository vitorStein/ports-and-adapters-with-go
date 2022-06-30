package app_test

import (
	"ports-and-adapter-with-go/app"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBook_Enable(t *testing.T) {
	book := app.Book{}
	book.Name = "Revolução dos Bichos"
	book.Status = app.DISABLED
	book.Price = 29.99

	err := book.Enable()
	require.Nil(t, err)

	book.Price = 0
	err = book.Enable()
	require.Equal(t, "o preço do livro deve ser maior que zero para habilitar o livro", err.Error())
}

func TestBook_Disable(t *testing.T) {
	book := app.Book{}
	book.Price = 0
	book.Status = app.ENABLE

	err := book.Disable()
	require.Nil(t, err)

	book.Price = 29.99
	err = book.Disable()
	require.Equal(t, "o preço do livro deve ser zero para ser desabilitado", err.Error())
}
