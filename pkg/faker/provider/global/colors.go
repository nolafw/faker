package global

import "github.com/nolafw/faker/pkg/faker/provider"

func CreateColors() *provider.Colors {
	return &provider.Colors{
		SafeColorNames: safeColorNames,
		AllColorNames:  allColorNames,
	}
}

var safeColorNames = []string{
	"black", "maroon", "green", "navy", "olive",
	"purple", "teal", "lime", "blue", "silver",
	"gray", "yellow", "fuchsia", "aqua", "white",
}

var allColorNames = []string{
	"AliceBlue", "AntiqueWhite", "Aqua", "Aquamarine",
	"Azure", "Beige", "Bisque", "Black", "BlanchedAlmond",
	"Blue", "BlueViolet", "Brown", "BurlyWood", "CadetBlue",
	"Chartreuse", "Chocolate", "Coral", "CornflowerBlue",
	"Cornsilk", "Crimson", "Cyan", "DarkBlue", "DarkCyan",
	"DarkGoldenRod", "DarkGray", "DarkGreen", "DarkKhaki",
	"DarkMagenta", "DarkOliveGreen", "Darkorange", "DarkOrchid",
	"DarkRed", "DarkSalmon", "DarkSeaGreen", "DarkSlateBlue",
	"DarkSlateGray", "DarkTurquoise", "DarkViolet", "DeepPink",
	"DeepSkyBlue", "DimGray", "DimGrey", "DodgerBlue", "FireBrick",
	"FloralWhite", "ForestGreen", "Fuchsia", "Gainsboro", "GhostWhite",
	"Gold", "GoldenRod", "Gray", "Green", "GreenYellow", "HoneyDew",
	"HotPink", "IndianRed", "Indigo", "Ivory", "Khaki", "Lavender",
	"LavenderBlush", "LawnGreen", "LemonChiffon", "LightBlue", "LightCoral",
	"LightCyan", "LightGoldenRodYellow", "LightGray", "LightGreen", "LightPink",
	"LightSalmon", "LightSeaGreen", "LightSkyBlue", "LightSlateGray", "LightSteelBlue",
	"LightYellow", "Lime", "LimeGreen", "Linen", "Magenta", "Maroon", "MediumAquaMarine",
	"MediumBlue", "MediumOrchid", "MediumPurple", "MediumSeaGreen", "MediumSlateBlue",
	"MediumSpringGreen", "MediumTurquoise", "MediumVioletRed", "MidnightBlue",
	"MintCream", "MistyRose", "Moccasin", "NavajoWhite", "Navy", "OldLace", "Olive",
	"OliveDrab", "Orange", "OrangeRed", "Orchid", "PaleGoldenRod", "PaleGreen",
	"PaleTurquoise", "PaleVioletRed", "PapayaWhip", "PeachPuff", "Peru", "Pink", "Plum",
	"PowderBlue", "Purple", "Red", "RosyBrown", "RoyalBlue", "SaddleBrown", "Salmon",
	"SandyBrown", "SeaGreen", "SeaShell", "Sienna", "Silver", "SkyBlue", "SlateBlue",
	"SlateGray", "Snow", "SpringGreen", "SteelBlue", "Tan", "Teal", "Thistle", "Tomato",
	"Turquoise", "Violet", "Wheat", "White", "WhiteSmoke", "Yellow", "YellowGreen",
}
