package bondApi

type User struct {
	Name     string
	LastName string
	Email    string
	Password string
	Bonds    []Bond
}
