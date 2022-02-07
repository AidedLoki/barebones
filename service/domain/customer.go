package domain

type customer struct {
	ID   string `json: "id"`
	Name string `json: "name"`
}

func NewCustomer(id string, name string) *customer {

	c := &customer{
		ID:   id,
		Name: name,
	}

	return c

}
