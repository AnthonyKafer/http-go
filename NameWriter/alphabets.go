package nameWriter

import "go-server/NameWriter/Fonts"

var alphabets = map[string]map[string][]string{
	"BIG":      Fonts.Big,
	"BIGMONEY": Fonts.BigMoney,
	"CUSTOM":   Fonts.Custom,
}
