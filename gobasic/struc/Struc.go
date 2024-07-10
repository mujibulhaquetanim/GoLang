package struc

type PersonMethods interface {
	acting()
	takingPhotos()
}

type Person struct {
	Name  string
	Age   int
	Hobby []string
}
