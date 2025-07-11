package en_US

import "github.com/nolafw/faker/pkg/faker/provider"

func CreateAddresses() *provider.Addresses {
	return &provider.Addresses{
		CityPrefixes:            CityPrefixes,
		CitySuffixes:            CitySuffixes,
		CityNames:               *CityNames,
		BuildingNumbers:         BuildingNumbers,
		StreetSuffixes:          StreetSuffixes,
		Postcodes:               Postcodes,
		States:                  States,
		StateAbbrs:              StateAbbrs,
		Countries:               Countries,
		CityFormats:             CityFormats,
		StreetFormats:           StreetFormats,
		StreetNames:             *StreetNames,
		StreetAddressFormats:    StreetAddressFormats,
		SecondaryAddressFormats: SecondaryAddressFormats,
		AddressFormats:          AddressFormats,
		CreateCity:              CreateUSCityFullName,
		CreateStreet:            CreateUSStreet,
		CreateStreetAddress:     CreateUSStreetAddress,
		CreateAddress:           CreateUSAddress,
		CreateSecondaryAddress:  CreateEnUsSecondaryAddress,
	}
}

var Countries = []string{
	"Afghanistan", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antarctica (the territory South of 60 deg S)", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan",
	"Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia and Herzegovina", "Botswana", "Bouvet Island (Bouvetoya)", "Brazil", "British Indian Ocean Territory (Chagos Archipelago)", "British Virgin Islands", "Brunei Darussalam", "Bulgaria", "Burkina Faso", "Burundi",
	"Cambodia", "Cameroon", "Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China", "Christmas Island", "Cocos (Keeling) Islands", "Colombia", "Comoros", "Congo", "Cook Islands", "Costa Rica", "Cote d'Ivoire", "Croatia", "Cuba", "Cyprus", "Czech Republic",
	"Denmark", "Djibouti", "Dominica", "Dominican Republic",
	"Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Ethiopia",
	"Faroe Islands", "Falkland Islands (Malvinas)", "Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern Territories",
	"Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana",
	"Haiti", "Heard Island and McDonald Islands", "Holy See (Vatican City State)", "Honduras", "Hong Kong", "Hungary",
	"Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy",
	"Jamaica", "Japan", "Jersey", "Jordan",
	"Kazakhstan", "Kenya", "Kiribati", "Korea", "Korea", "Kuwait", "Kyrgyz Republic",
	"Lao People's Democratic Republic", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libyan Arab Jamahiriya", "Liechtenstein", "Lithuania", "Luxembourg",
	"Macao", "Macedonia", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia", "Moldova", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar",
	"Namibia", "Nauru", "Nepal", "Netherlands Antilles", "Netherlands", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk Island", "Northern Mariana Islands", "Norway",
	"Oman",
	"Pakistan", "Palau", "Palestinian Territories", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Pitcairn Islands", "Poland", "Portugal", "Puerto Rico",
	"Qatar",
	"Reunion", "Romania", "Russian Federation", "Rwanda",
	"Saint Barthelemy", "Saint Helena", "Saint Kitts and Nevis", "Saint Lucia", "Saint Martin", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia (Slovak Republic)", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Georgia and the South Sandwich Islands", "Spain", "Sri Lanka", "Sudan", "Suriname", "Svalbard & Jan Mayen Islands", "Swaziland", "Sweden", "Switzerland", "Syrian Arab Republic",
	"Taiwan", "Tajikistan", "Tanzania", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu",
	"Uganda", "Ukraine", "United Arab Emirates", "United Kingdom", "United States of America", "United States Minor Outlying Islands", "United States Virgin Islands", "Uruguay", "Uzbekistan",
	"Vanuatu", "Venezuela", "Vietnam",
	"Wallis and Futuna", "Western Sahara",
	"Yemen",
	"Zambia", "Zimbabwe",
}

var Postcodes = []string{
	"#####", "#####-####",
}

var States = []string{
	"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "District of Columbia", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
}

var StateAbbrs = []string{
	"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME", "MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM", "NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VA", "VT", "WA", "WI", "WV", "WY",
}

// City

var CityNames = &LastNames
var CityPrefixes = []string{
	"North", "East", "West", "South", "New", "Lake", "Port",
}
var CitySuffixes = []string{
	"town", "ton", "land", "ville", "berg", "burgh", "borough", "bury", "view", "port", "mouth", "stad", "furt", "chester", "mouth", "fort", "haven", "side", "shire",
}
var CityFormats = []string{
	"{{.CityPrefix}} {{.CityName}}{{.CitySuffix}}",
	"{{.CityPrefix}} {{.CityName}}",
	"{{.CityName}}{{.CitySuffix}}",
	"{{.CityName}}{{.CitySuffix}}",
}

type USCityName struct {
	CityPrefix string
	CitySuffix string
	CityName   string
}

type CityNameGenerator interface {
	CityPrefix() string
	CitySuffix() string
	CityName() string
}

func CreateUSCityFullName(a any) any {
	g := a.(CityNameGenerator)
	return &USCityName{
		CityPrefix: g.CityPrefix(),
		CitySuffix: g.CitySuffix(),
		CityName:   g.CityName(),
	}
}

// Street

var StreetNames = CityNames
var StreetSuffixes = []string{
	"Alley", "Avenue", "Branch", "Bridge", "Brook", "Brooks", "Burg", "Burgs", "Bypass", "Camp", "Canyon", "Cape", "Causeway", "Center", "Centers", "Circle", "Circles", "Cliff", "Cliffs", "Club", "Common", "Corner", "Corners", "Course", "Court", "Courts", "Cove", "Coves", "Creek", "Crescent", "Crest", "Crossing", "Crossroad", "Curve", "Dale", "Dam", "Divide", "Drive", "Drive", "Drives", "Estate", "Estates", "Expressway", "Extension", "Extensions", "Fall", "Falls", "Ferry", "Field", "Fields", "Flat", "Flats", "Ford", "Fords", "Forest", "Forge", "Forges", "Fork", "Forks", "Fort", "Freeway", "Garden", "Gardens", "Gateway", "Glen", "Glens", "Green", "Greens", "Grove", "Groves", "Harbor", "Harbors", "Haven", "Heights", "Highway", "Hill", "Hills", "Hollow", "Inlet", "Inlet", "Island", "Island", "Islands", "Islands", "Isle", "Isle", "Junction", "Junctions", "Key", "Keys", "Knoll", "Knolls", "Lake", "Lakes", "Land", "Landing", "Lane", "Light", "Lights", "Loaf", "Lock", "Locks", "Locks", "Lodge", "Lodge", "Loop", "Mall", "Manor", "Manors", "Meadow", "Meadows", "Mews", "Mill", "Mills", "Mission", "Mission", "Motorway", "Mount", "Mountain", "Mountain", "Mountains", "Mountains", "Neck", "Orchard", "Oval", "Overpass", "Park", "Parks", "Parkway", "Parkways", "Pass", "Passage", "Path", "Pike", "Pine", "Pines", "Place", "Plain", "Plains", "Plains", "Plaza", "Plaza", "Point", "Points", "Port", "Port", "Ports", "Ports", "Prairie", "Prairie", "Radial", "Ramp", "Ranch", "Rapid", "Rapids", "Rest", "Ridge", "Ridges", "River", "Road", "Road", "Roads", "Roads", "Route", "Row", "Rue", "Run", "Shoal", "Shoals", "Shore", "Shores", "Skyway", "Spring", "Springs", "Springs", "Spur", "Spurs", "Square", "Square", "Squares", "Squares", "Station", "Station", "Stravenue", "Stravenue", "Stream", "Stream", "Street", "Street", "Streets", "Summit", "Summit", "Terrace", "Throughway", "Trace", "Track", "Trafficway", "Trail", "Trail", "Tunnel", "Tunnel", "Turnpike", "Turnpike", "Underpass", "Union", "Unions", "Valley", "Valleys", "Via", "Viaduct", "View", "Views", "Village", "Village", "Villages", "Ville", "Vista", "Vista", "Walk", "Walks", "Wall", "Way", "Ways", "Well", "Wells",
}
var StreetFormats = []string{
	"{{.StreetName}}",
	"{{.StreetName}} {{.StreetSuffix}}",
	"{{.StreetName}} {{.StreetSuffix}}",
}

type USStreetName struct {
	StreetName   string
	StreetSuffix string
}

type StreetNameGenerator interface {
	StreetName() string
	StreetSuffix() string
}

func CreateUSStreet(a any) any {
	g := a.(StreetNameGenerator)
	return &USStreetName{
		StreetName:   g.StreetName(),
		StreetSuffix: g.StreetSuffix(),
	}
}

// Street Address

var StreetAddressFormats = []string{
	"{{.BuildingNumber}} {{.StreetName}}",
	"{{.BuildingNumber}} {{.StreetName}} {{.SecondaryAddress}}",
}

type USStreetAddress struct {
	BuildingNumber   string
	StreetName       string
	SecondaryAddress string
}

type StreetAddressGenerator interface {
	BuildingNumber() string
	StreetName() string
	SecondaryAddress() string
}

func CreateUSStreetAddress(a any) any {
	g := a.(StreetAddressGenerator)
	return &USStreetAddress{
		BuildingNumber:   g.BuildingNumber(),
		StreetName:       g.StreetName(),
		SecondaryAddress: g.SecondaryAddress(),
	}
}

// Secondary Address

var BuildingNumbers = []string{
	"####", "###", "##",
}
var SecondaryAddressFormats = []string{
	"Apt. {{.BuildingNumber}}", "Suite {{.BuildingNumber}}",
}

type SecondaryAddress struct {
	BuildingNumber string
}

type SecondaryAddressGenerator interface {
	BuildingNumber() string
}

func CreateEnUsSecondaryAddress(a any) any {
	g := a.(SecondaryAddressGenerator)

	return &SecondaryAddress{
		BuildingNumber: g.BuildingNumber(),
	}
}

// Address

var AddressFormats = []string{
	"{{.StreetAddress}}\n{{.City}}, {{.StateAbbr}} {{.Postcode}}",
}

type USAddress struct {
	StreetAddress string
	City          string
	StateAbbr     string
	Postcode      string
}

type AddressGenerator interface {
	StreetAddress() string
	City() string
	StateAbbr() string
	Postcode() string
}

func CreateUSAddress(a any) any {
	g := a.(AddressGenerator)
	return &USAddress{
		StreetAddress: g.StreetAddress(),
		City:          g.City(),
		StateAbbr:     g.StateAbbr(),
		Postcode:      g.Postcode(),
	}
}
