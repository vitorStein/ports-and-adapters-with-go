package app

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

func (b *Book) IsValid(bool, error) {

}

func (b *Book) Enable(error) {

}

func (b *Book) Disable(error) {

}

func (b *Book) GetID(string) {

}

func (b *Book) GetName(string) {

}

func (b *Book) GetStatus(string) {

}

func (b *Book) GetPrice(float32) {

}
