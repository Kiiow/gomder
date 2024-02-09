package format

import (
	"fmt"
	"math"
)

func Bytes_fmt(num int64) string {
	numf := float64(num)
	for _, unit := range []string{" b", "kb", "mb", "Gb", "Tb"} {
		if math.Abs(numf) < 1024 {
			return fmt.Sprintf("%3.1f %s", numf, unit)
		}
		numf /= 1024
	}
	return fmt.Sprintf("%3.1f %s", numf, "Pb")
}
