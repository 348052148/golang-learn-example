package creates

type (
	Color int
	Wheels int
	Speed int
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build()CarInterface
}

type CarBuilder struct {
	Colors Color
	Wheelss Wheels
	Speed Speed
}

func NewCarBuilrer() CarBuilder {
	return CarBuilder{}
}
func (b CarBuilder)Color( c Color) Builder{
	b.Colors = c
	return b
}
func (b CarBuilder)Wheels(w Wheels) Builder {
	b.Wheelss = w
	return b
}
func (b CarBuilder)TopSpeed(s Speed) Builder {
	b.Speed = s
	return b
}
func (b  CarBuilder)Build() CarInterface {
	return &Car{}
}

type CarInterface interface {
	Drive() error
	Stop() error
}

type Car struct {

}

func (car * Car)Drive()error  {
	return nil
}
func (car * Car)Stop()error  {
	return nil
}

func Use()  {
	
}