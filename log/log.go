package log

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

// Errorf logs as an error with the given params
func Errorf(params ...interface{}) {
	red := color.BgRed.Render
	fmt.Printf(fmt.Sprintf("%s %v\n", red("ERROR:"), params[0]), params[1:]...)
	os.Exit(3)
}