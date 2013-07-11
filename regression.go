package regression
 
import (
    "github.com/skelterjohn/go.matrix"
    "fmt"
    "strings"
    "strconv"
    "math"
)

type Regression struct {
    Names Describe
    Data []DataPoint
    RegCoeff map[int]float64
    Rsquared float64
    VarianceObserved float64
    VariancePredicted float64
    Debug bool
}

type DataPoint struct {
    Observed float64
    Variables []float64
    Predicted float64
    Error float64
}

type Describe struct {
    ObservedName string
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
        s := []string{"X", strconv.Itoa(i)};
        return strings.Join(s, "")
    }
    return x
}

func (r *Regression) AddDataPoint(d DataPoint) {
    r.Data = append(r.Data, d)
}

func (r *Regression) RunLinearRegression() {
    // First move the data points into the right format
    observations := len(r.Data)
    numOfVars := len(r.Data[0].Variables)

    if observations < (numOfVars+1) {
        fmt.Println("Error: Not enough observations to perform the regression.")
        return
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
        fmt.Println(err)
        return
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
    fmt.Println("\n-----------------------------------------------------------------")
    for i, val := range c {
        r.RegCoeff[i] = val
        if i == 0 {
            fmt.Print("Predicted = ", val)
        } else {
            fmt.Print(" + ", r.GetVarName(i-1), " * ", val)
        }

    }
    fmt.Println("\n-----------------------------------------------------------------\n")
    r.calcPredicted()
    r.calcVariance()
    r.calcRsquared()
}

func (r *Regression) GetRegCoeff(i int) float64 {
    if len(r.RegCoeff) == 0 {
        return 0
    }
    return r.RegCoeff[i]
}

func (r *Regression) calcPredicted() {
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
            fmt.Println(i, " Observed = ", r.Data[i].Observed, ", Predicted = ", predicted, ", Error = ", r.Data[i].Error)
        }
    }
}

func (r *Regression) calcVariance() {
    observations := len(r.Data)
    var obtotal, prtotal, obvar, prvar float64
    for i := 0; i < observations; i++ {
        obtotal += r.Data[i].Observed
        prtotal += r.Data[i].Predicted
    }
    obaverage := obtotal / float64(observations)
    praverage := prtotal / float64(observations)

    for i := 0; i < observations; i++ {
        obvar += math.Pow(r.Data[i].Observed - obaverage, 2)
        prvar += math.Pow(r.Data[i].Predicted - praverage, 2)
    }
    r.VarianceObserved = obvar / float64(observations)
    r.VariancePredicted = prvar / float64(observations)
    if r.Debug {
        fmt.Println("Variance Observed = ", r.VarianceObserved)
        fmt.Println("Variance Predicted = ", r.VariancePredicted)
    }
}

func (r *Regression) calcRsquared() {
    r.Rsquared = r.VariancePredicted / r.VarianceObserved
    if r.Debug {
        fmt.Println("R2 = ", r.Rsquared)
    }
}

func (r *Regression) Dump(data bool) {
    temp := r.Debug
    if data == true {
        r.Debug = true
        r.calcPredicted()
        r.Debug = temp
    } else {
        r.calcPredicted()
    }
    fmt.Println("-------------------------")
    fmt.Println("Variance Observed = ", r.VarianceObserved)
    fmt.Println("Variance Predicted = ", r.VariancePredicted)
    fmt.Println("R2 = ", r.Rsquared)
    fmt.Println("-------------------------\n")
}
