package en_US

import (
	"github.com/nolafw/faker/pkg/fk/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People:    CreatePeople(),
		Addresses: CreateAddresses(),
		Companies: CreateCompanies(),
	}
}
