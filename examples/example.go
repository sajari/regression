package main

import "github.com/sajari/regression"

func main() {
	r := new(Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")
	r.Train(
		DataPoint(11.2, []float64{587000, 16.5, 6.2}),
		DataPoint(13.4, []float64{643000, 20.5, 6.4}),
		DataPoint(40.7, []float64{635000, 26.3, 9.3}),
		DataPoint(5.3, []float64{692000, 16.5, 5.3}),
		DataPoint(24.8, []float64{1248000, 19.2, 7.3}),
		DataPoint(12.7, []float64{643000, 16.5, 5.9}),
		DataPoint(20.9, []float64{1964000, 20.2, 6.4}),
		DataPoint(35.7, []float64{1531000, 21.3, 7.6}),
		DataPoint(8.7, []float64{713000, 17.2, 4.9}),
		DataPoint(9.6, []float64{749000, 14.3, 6.4}),
		DataPoint(14.5, []float64{7895000, 18.1, 6}),
		DataPoint(26.9, []float64{762000, 23.1, 7.4}),
		DataPoint(15.7, []float64{2793000, 19.1, 5.8}),
		DataPoint(36.2, []float64{741000, 24.7, 8.6}),
		DataPoint(18.1, []float64{625000, 18.6, 6.5}),
		DataPoint(28.9, []float64{854000, 24.9, 8.3}),
		DataPoint(14.9, []float64{716000, 17.9, 6.7}),
		DataPoint(25.8, []float64{921000, 22.4, 8.6}),
		DataPoint(21.7, []float64{595000, 20.2, 8.4}),
		DataPoint(25.7, []float64{3353000, 16.9, 6.7}),
	)
	r.Run()
}
