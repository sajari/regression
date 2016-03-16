# regression
[![Build Status](https://travis-ci.org/sajari/regression.svg?branch=master)](https://travis-ci.org/sajari/regression)

Multivariable Linear Regression in Go (golang)

## Install

```
go get github.com/sajari/regression
```

## Usage

Import the package, create a regression and add data to it. You can use as many variables as you like, in the below example there are 3 variables for each observation

```go
package main

import "github.com/sajari/regression"

func main() {
    var r regression.Regression
    r.SetObservedName("Z")
    r.SetVarName(0, "A")
    r.SetVarName(1, "B")
    r.SetVarName(2, "C")
    // Z = 12 + 1A + 2B +3C
    r.AddDataPoint(regression.DataPoint{Observed: 12, Variables: []float64{0, 0, 0}})
    r.AddDataPoint(regression.DataPoint{Observed: 13, Variables: []float64{1, 0, 0}})
    r.AddDataPoint(regression.DataPoint{Observed: 14, Variables: []float64{0, 1, 0}})
    r.AddDataPoint(regression.DataPoint{Observed: 15, Variables: []float64{0, 0, 1}})
    r.AddDataPoint(regression.DataPoint{Observed: 14, Variables: []float64{2, 0, 0}})
    r.AddDataPoint(regression.DataPoint{Observed: 16, Variables: []float64{0, 2, 0}})
    r.AddDataPoint(regression.DataPoint{Observed: 18, Variables: []float64{0, 0, 2}})

    r.RunLinearRegression()
    r.Dump(true)
}
```

```shell
$ go run main.go


0. Observed = 12, Predicted = 11.999999999999995, Error = -5.329070518200751e-15
1. Observed = 13, Predicted = 12.999999999999998, Error = -1.7763568394002505e-15
2. Observed = 14, Predicted = 13.999999999999996, Error = -3.552713678800501e-15
3. Observed = 15, Predicted = 14.999999999999998, Error = -1.7763568394002505e-15
4. Observed = 14, Predicted = 14.000000000000002, Error = 1.7763568394002505e-15
5. Observed = 16, Predicted = 15.999999999999998, Error = -1.7763568394002505e-15
6. Observed = 18, Predicted = 18, Error = 0

N =  7
Variance Observed =  3.3877551020408165
Variance Predicted =  3.3877551020408205
R2 =  1.000000000000001
Formula =  Predicted = 12.0000 + A*1.0000 + B*2.0000 + C*3.0000
----------------------------------
```
