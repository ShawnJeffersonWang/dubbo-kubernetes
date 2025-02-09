package art

import (
	_ "embed"
	"github.com/fatih/color"
)

//go:embed dubbo-ascii.txt
var dubboASCIIArt string

func DubboColoredArt() string {
	return color.New(color.FgHiBlue).Add(color.Bold).Sprint(dubboASCIIArt)
}
