package format

import (
	"fmt"
	"math"
)

func FormatSize(size int64, humanize bool) string {
	resultSize := float64(size)
	units := map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
	unit := units[0]

	if humanize {
		var devider float64 = 1024

		pow := math.Floor(math.Log(resultSize) / math.Log(devider))
		_unit, ok := units[int(pow)]
		if !ok {
			unit = "MoreThanYouCanImagineBytes"
		} else {
			unit = _unit
		}
		resultSize = math.Round((resultSize/math.Pow(devider, pow))*10) / 10
	}
	if unit == units[0] {
		return fmt.Sprintf("%d%s", int64(resultSize), unit)
	}
	return fmt.Sprintf("%.1f%s", resultSize, unit)
}
