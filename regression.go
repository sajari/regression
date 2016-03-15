package regression

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/skelterjohn/go.matrix"
)

type Regression struct {
	Names             Describe
	Data              []DataPoint
	RegCoeff          map[int]float64
	Rsquared          float64
	VarianceObserved  float64
	VariancePredicted float64
	Debug             bool
	Initialised       bool
	Formula           string
}

type DataPoint struct {
	Observed  float64
	Variables []float64
	Predicted float64
	Error     float64
}

type Describe struct {
	ObservedName  string
	VariableNames map[int]string
}

func (r *Regression) SetObservedName(name string) {
	r.Names.ObservedName = name
}

func (r *Regression) SetVarName(i int, name string) {
	if len(r.Names.VariableNames) == 0 {
		r.Names.VariableNames = make(map[int]string, 5)
	}
	r.Names.VariableNames[i] = name
}

func (r *Regression) GetObservedName() string {
	return r.Names.ObservedName
}

func (r *Regression) GetVarName(i int) string {
	x := r.Names.VariableNames[i]
	if x == "" {
		s := []string{"X", strconv.Itoa(i)}
		return strings.Join(s, "")
	}
	return x
}

func (r *Regression) AddDataPoint(d DataPoint) {
	r.Data = append(r.Data, d)
	r.Initialised = true
}

func (r *Regression) RunLinearRegression() error {
	if !r.Initialised {
		return errors.New("you need some observations to perform the regression")
	}

	observations := len(r.Data)
	numOfVars := len(r.Data[0].Variables)

	if observations < (numOfVars + 1) {
		return fmt.Errorf("not enough observations to to support %v variable(s)", numOfVars)
	}

	// Create some blank variable space
	observed := matrix.Zeros(observations, 1)
	variables := matrix.Zeros(observations, numOfVars+1)

	for i := 0; i < observations; i++ {
		observed.Set(i, 0, r.Data[i].Observed)
		for j := 0; j < numOfVars+1; j++ {
			if j == 0 {
				variables.Set(i, 0, 1)
			} else {
				variables.Set(i, j, r.Data[i].Variables[j-1])
			}
		}
	}
	if r.Debug {
		fmt.Println("---------- VARS ---------\n")
		fmt.Println(variables)
		fmt.Println("-------- OBSERVED -------\n")
		fmt.Println(observed)
	}

	// Now run the regression
	n := variables.Cols()
	q, reg := variables.QR()
	qty, err := q.Transpose().Times(observed)
	if err != nil {
		return err
	}
	c := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		c[i] = qty.Get(i, 0)
		for j := i + 1; j < n; j++ {
			c[i] -= c[j] * reg.Get(i, j)
		}
		c[i] /= reg.Get(i, i)
	}

	// Output the regression results
	r.RegCoeff = make(map[int]float64, numOfVars)
	for i, val := range c {
		r.RegCoeff[i] = val
	}

	// Calculate the regression formula
	for i, val := range c {
		if i == 0 {
			r.Formula = fmt.Sprintf("Predicted = %.4f", val)
		} else {
			r.Formula += fmt.Sprintf(" + %v*%.4f", r.GetVarName(i-1), val)
		}
	}

	if r.Debug {
		fmt.Println("\n-----------------------------------------------------------------")
		fmt.Printf("%v\n\n", r.Formula)
	}

	r.calcPredicted()
	r.calcVariance()
	r.calcRsquared()
	return nil
}

func (r *Regression) GetRegCoeff(i int) float64 {
	if len(r.RegCoeff) == 0 {
		return 0
	}
	return r.RegCoeff[i]
}

func (r *Regression) calcPredicted() {
	if r.Debug {
		fmt.Println("\n")
	}

	observations := len(r.Data)
	numOfVars := len(r.Data[0].Variables)
	var predicted float64

	for i := 0; i < observations; i++ {
		for j := 0; j < numOfVars+1; j++ {
			if j == 0 {
				predicted = r.GetRegCoeff(j)
			} else {
				predicted += r.GetRegCoeff(j) * r.Data[i].Variables[j-1]
			}
		}
		r.Data[i].Error = predicted - r.Data[i].Observed
		r.Data[i].Predicted = predicted
		if r.Debug {
			fmt.Printf("%v. Observed = %v, Predicted = %v, Error = %v \n", i, r.Data[i].Observed, predicted, r.Data[i].Error)
		}
	}
}

func (r *Regression) calcVariance() {
	if r.Debug {
		fmt.Println("\n")
	}

	observations := len(r.Data)
	var obtotal, prtotal, obvar, prvar float64
	for i := 0; i < observations; i++ {
		obtotal += r.Data[i].Observed
		prtotal += r.Data[i].Predicted
	}
	obaverage := obtotal / float64(observations)
	praverage := prtotal / float64(observations)

	for i := 0; i < observations; i++ {
		obvar += math.Pow(r.Data[i].Observed-obaverage, 2)
		prvar += math.Pow(r.Data[i].Predicted-praverage, 2)
	}
	r.VarianceObserved = obvar / float64(observations)
	r.VariancePredicted = prvar / float64(observations)
	if r.Debug {
		fmt.Println("N = ", observations)
		fmt.Println("Variance Observed = ", r.VarianceObserved)
		fmt.Println("Variance Predicted = ", r.VariancePredicted)
	}
}

func (r *Regression) calcRsquared() {
	r.Rsquared = r.VariancePredicted / r.VarianceObserved
	if r.Debug {
		fmt.Println("\nR2 = ", r.Rsquared)
	}
}

func (r *Regression) Dump(data bool) error {
	if !r.Initialised {
		return errors.New("you need some observations before you can dump the data")
	}
	temp := r.Debug
	if data {
		r.Debug = true
		r.calcPredicted()
		r.Debug = temp
	} else {
		r.calcPredicted()
	}
	fmt.Println("\nN = ", len(r.Data))
	fmt.Println("Variance Observed = ", r.VarianceObserved)
	fmt.Println("Variance Predicted = ", r.VariancePredicted)
	fmt.Println("R2 = ", r.Rsquared)
	fmt.Println("Formula = ", r.Formula)
	fmt.Println("-----------------------------------------------------------------\n")
	return nil
}
