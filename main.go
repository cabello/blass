package main

import (
	"fmt"
    "bloom"
    "math"
    "text/tabwriter"
    "os"
    // "random"
)

func Btoi(b bool) float64 {
    if b {
        return 1.0
    }
    return 0.0
}

func bench() {
    var results map[string]float64
    results = make(map[string]float64)

    w := new(tabwriter.Writer)

    // Format in tab-separated columns with a tab stop of 8.
    w.Init(os.Stdout, 0, 8, 0, '\t', 0)
    fmt.Fprintln(w, "N Elements\tError Prob\tM Bytes\tFalse Negative\tFalse Positive\tActual Prob\tTest")

    // fmt.Fprintln(w, "a\tb\tc\td\t.")
    // fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")

    // results[] = ['N Elements','Error Prob','M Bytes','False Negative','False Positive','Actual Prob'];
    for i := 1.0; i < 6; i++ {
        n := float64(math.Pow(10,i));
        for p := 0.1; p > 0.0001; p /= 10 {
            results["N Elements"] = n;
            results["Error Prob"] = p;

            fmt.Fprintf(w, "%v\t%v\t", n, p)

            filter := bloom.New(int(n), p)

            results["M Bytes"] = float64(filter.GetSizeInBytes())

            fmt.Fprintf(w, "%v\t", filter.GetSizeInBytes())

            falseNegative := 0.0;
            falsePositive := 0.0;

            limit := n * 3;

            for k := 0.0; k < limit; k+= 3.0 {
                filter.Add([]byte(fmt.Sprintf("T%v", k)));
            }

            // fmt.Println(filter.BitArray)

            samples := n * 9;
            for k := 0.0; k < samples; k++ {
                if int(k) % 3 == 0 && k < limit {
                    falseNegative += Btoi(!filter.Contains([]byte(fmt.Sprintf("T%v", k))));
                } else if k > limit {
                    falsePositive += Btoi(filter.Contains([]byte(fmt.Sprintf("T%v", k))));
                }
            }

            results["False Negative"] = falseNegative;
            results["False Positive"] = falsePositive;
            results["Actual Prob"] = falsePositive / samples;
            if falsePositive / samples < p && falseNegative == 0 {
                results["Test"] = 1
            } else {
                results["Test"] = 0
            }

            fmt.Fprintf(w, "%v\t%v\t%.4g\t%v\n", falseNegative, falsePositive, falsePositive / samples, falsePositive / samples < p && falseNegative == 0)

            // fmt.Println(results);
            // fmt.Println(filter.BitArray)
        }
    }
    w.Flush()
}

func main() {
    // var keys []string = make([]string, 100)

    // var b bloom.Filter = bloom.New(250, 0.001)

    // rs := random.NewRS("abcdefghijklmnopqrstuvxz")
    // for i := 0; i < 100; i++ {
    //     keys[i] = rs.NewRandomString(16)
    //     b.Add([]byte(keys[i]))
    // }

    // fmt.Printf("%+v\n", b);

    // fmt.Printf("%v\n", b.Contains([]byte("guilherme")))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[0])))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[50])))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[99])))
    // fmt.Printf("%v\n", b.Contains([]byte("adalberto")))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[23])))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[44])))
    // fmt.Printf("%v\n", b.Contains([]byte(keys[11])))
    bench()
}
