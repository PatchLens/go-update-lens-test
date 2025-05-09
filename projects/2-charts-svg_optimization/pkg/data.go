package pkg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/go-analyze/charts"
)

const startMM = 60
const endMM = 510

var lensDefinitions = map[string]string{
	"70-200mm f/2.8":           "70f2.8,201f-", // use a string encoding to define the f/stop point changes
	"70-200mm f/2.8 + 1.4x TC": "98f4,281f-",
	"70-200mm f/2.8 + 2x TC":   "140f5.6,401f-",
	"100-500mm f/4.5-7.1":      "100f4.5,151f5,254f5.6,363f6.3,472f7.1,501f-",
}

func PopulateData() (values [][]float64, xAxisLabels []string, labels []string, err error) {
	for k := range lensDefinitions {
		labels = append(labels, k)
	}
	sort.Slice(labels, func(i, j int) bool {
		return labels[i] < labels[j]
	})

	for i := startMM; i <= endMM; i++ {
		xAxisLabels = append(xAxisLabels, fmt.Sprintf("%vmm", i))
	}

	for _, lens := range labels {
		parts := strings.Split(lensDefinitions[lens], ",")
		count := (endMM - startMM) + 1
		lensValues := make([]float64, count)
		currentFValue := charts.GetNullValue()
		// for code simplicity, we assume startMM is strictly BEFORE the first lens, this allows us to set null
		// values until the start point (which will be loaded on the first run of the loop)
		nextPartIndex := 0
		nextMM := startMM
		nextFStop := currentFValue
		for i := 0; i < count; i++ {
			if i+startMM == nextMM {
				currentFValue = nextFStop
				if nextPartIndex < len(parts) {
					nextFStop, nextMM, err = parseFStopMM(parts[nextPartIndex])
					nextPartIndex++
					if err != nil {
						return
					}
				} else {
					nextFStop = charts.GetNullValue()
				}
			}
			lensValues[i] = currentFValue
		}
		values = append(values, lensValues)
	}

	return
}

func parseFStopMM(str string) (float64, int, error) {
	parts := strings.Split(str, "f")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid lens spec str: '%s'", str)
	}
	mm, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if parts[1] == "-" {
		return charts.GetNullValue(), mm, nil
	}
	fstop, err := strconv.ParseFloat(parts[1], 64)
	return fstop, mm, err
}
