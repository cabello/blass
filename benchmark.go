package main

import (
	"bloom"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func boolToFloat(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func benchmark(capacityBase, capacityPowStart, capacityPowEnd float64) {
	errorRateBase := 10.0
	errorRateStart := 0.1
	errorRateEnd := 0.00001

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Elements\tError Prob.\tSize (Bytes)\tFalse Negative\tFalse Positive\tObserved Error\tPass?")

	for i := capacityPowStart; i < capacityPowEnd; i++ {
		capacity := math.Pow(capacityBase, i)

		for errorRate := errorRateStart; errorRate > errorRateEnd; errorRate /= errorRateBase {
			filter := bloom.New(int(capacity), errorRate)

			falseNegative := 0.0
			falsePositive := 0.0

			limit := capacity * 3
			for k := 0.0; k < limit; k += 3.0 {
				filter.Add([]byte(fmt.Sprintf("T%v", k)))
			}

			samples := capacity * 9
			for k := 0.0; k < samples; k++ {
				if int(k)%3 == 0 && k < limit {
					falseNegative += boolToFloat(!filter.Contains([]byte(fmt.Sprintf("T%v", k))))
				} else if k > limit {
					falsePositive += boolToFloat(filter.Contains([]byte(fmt.Sprintf("T%v", k))))
				}
			}

			observedError := falsePositive / samples
			pass := falsePositive/samples < errorRate && falseNegative == 0

			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%.4g\t%v\n", capacity, errorRate, filter.GetSizeInBytes(), falseNegative, falsePositive, observedError, pass)
		}
	}
	w.Flush()
}

func main() {
	// from capacity 10 ^ 1 to 10 ^ 6
	benchmark(10.0, 1.0, 6.0)
	// from capacity 2 ^4 to 2 ^ 16
	benchmark(2.0, 4.0, 16.0)
}
