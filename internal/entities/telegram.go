package entities

const (
	DaysKey               = "."
	PromoKey              = ","
	MonthKey              = ":"
	CompaniesKey          = ";"
	PaymentKey            = "!"
	UserAccountKey        = "@"
	AddDemoKey            = "`"
	ServersKey            = "#"
	RecurrentNonRecurrent = "?"
	RebillID              = "§"
)

var SymbolsToNum = map[string]int{
	AdminRoot:        0,
	ServerRoot:       1,
	PaymentRoot:      2,
	"d":              3,
	"e":              4,
	InstructionsRoot: 5,
	"g":              6,
	ReviewsRoot:      7,
	AddDemoRoot:      8,
	"j":              9,
	"k":              10,
	"l":              11,
	"m":              12,
	"n":              13,
	"o":              14,
	"p":              15,
	"q":              16,
	"r":              17,
	"s":              18,
	"t":              19,
	"u":              20,
	"v":              21,
	"w":              22,
	"x":              23,
	"y":              24,
	"z":              25,
}

var NumsToSymbols = map[int]string{
	0:  AdminRoot,
	1:  ServerRoot,
	2:  PaymentRoot,
	3:  "d",
	4:  "e",
	5:  InstructionsRoot,
	6:  "g",
	7:  ReviewsRoot,
	8:  AddDemoRoot,
	9:  "j",
	10: "k",
	11: "l",
	12: "m",
	13: "n",
	14: "o",
	15: "p",
	16: "q",
	17: "r",
	18: "s",
	19: "t",
	20: "u",
	21: "v",
	22: "w",
	23: "x",
	24: "y",
	25: "z",
}
