package ja_JP

import (
	"github.com/nolafw/faker/pkg/faker/provider"
)

func New() *provider.Localized {
	return &provider.Localized{
		People:    CreatePeople(),
		Addresses: CreateAddresses(),
		Companies: CreateCompanies(),
	}
}
