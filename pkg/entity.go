package pkg

type Machine struct {
	Name     string
	Products []*Product
}

type Product struct {
	Name  string
	Count int32
	Price int32
}
