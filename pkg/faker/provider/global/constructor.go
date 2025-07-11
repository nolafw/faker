package global

import "github.com/nolafw/faker/pkg/faker/provider"

func New() *provider.Global {
	return &provider.Global{
		Colors:    CreateColors(),
		Files:     CreateFiles(),
		Images:    CreateImages(),
		Internets: CreateInternets(),
		Lorems:    CreateLorems(),
	}
}
