package lib

import (
	"fmt"
	"gopkg.in/qamarian-lib/str.v3"
	"testing"
)

// Function TestConf () tests function Conf ().
func TestConf (t *testing.T) {
	str.PrintEtr ("Function Conf () test has started...", "std", "TestConf ()")
	a, b, c, d, e, f, g := Conf ()
	str.PrintEtr ("Use test result to determine is test was passed:", "std",
		"TestConf ()")
	fmt.Println (fmt.Sprintf ("a: %s, b: %s, c: %d, d: %d, e: %d, f: %s, g: %v", a, b,
		c, d, e, f, g))
	str.PrintEtr ("Test completed.", "std", "TestConf ()")
}
