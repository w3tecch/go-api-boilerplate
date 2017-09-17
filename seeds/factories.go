package seeds

import (
	"github.com/icrowley/fake"
	"github.com/w3tecch/go-api-boilerplate/app/models"
	"github.com/w3tecch/go-api-boilerplate/lib/random"
)

// FakeUser creates a user with faked data
func FakeUser() (u models.User) {
	firstName := fake.FirstName()
	lastName := fake.LastName()
	email := firstName + "." + lastName + "@" + fake.DomainName()
	birthday := random.Date()
	passCode := random.Int(1000000, 999999)
	weight := random.Float64(30, 120)

	u.FirstName = &firstName
	u.LastName = &lastName
	u.Email = &email
	u.Birthday = &birthday
	u.PassCode = &passCode
	u.Weight = &weight
	return u
}
