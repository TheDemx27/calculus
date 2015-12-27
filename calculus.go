package calculus // import "github.com/TheDemx27/calculus"

import (
	"math"
	"fmt"
)

func SetPrec(defaultPrec int, usrPrec []int) int {
	var prec int
	if len(usrPrec) > 0 {
		prec = usrPrec[0]
	} else {
		prec = defaultPrec
	}
	return prec
}

func (fnc Function) SymDiff() string {
	 fmt.Println(fnc.ParseToks())
	 return "compiles..."
	// exp := GroupTerms(fnc.GetToksAbstract, fnc.GetToksLit())
	//
	// switch (exp) {
	// 	case "C*V":
	// 	case "C*V^C":
	// 	case "C*V^V":
	// 	case "ln(V)":
	// 	case "sin(V)":
	// 	case "cos(V)":
	// 	case "-sin(V)":
	// 	case "-cos(V)":
	// 	case "1/V":
	// 	case "V*V":
	// 	case "V/V":
	// 	case "":
	// }
	// return exp
}

func (fnc Function) Diff(point float64, usrPrec ...int) float64 {
	prec := SetPrec(10000000, usrPrec)
	h := 1 / float64(prec)
	return (fnc.Eval(point+h) - fnc.Eval(point)) / h
}

func (fnc Function) AntiDiff(lower float64, upper float64, usrPrec ...int) float64 {
	prec := SetPrec(1000, usrPrec)
	parts := make([]float64, 2*prec+1)
	r := upper - lower
	nf := float64(prec)
	dx0 := r / nf
	parts[0] = fnc.Eval(lower) * dx0
	parts[1] = fnc.Eval(lower+dx0*.5) * dx0 * 4
	x0 := lower + dx0
	for i := 1; i < prec; i++ {
		x1 := lower + float64(i+1)*r/nf
		xmid := (x0 + x1) * .5
		dx := x1 - x0
		parts[2*i] = fnc.Eval(x0) * dx * 2
		parts[2*i+1] = fnc.Eval(xmid) * dx * 4
		x0 = x1
	}
	parts[2*prec] = fnc.Eval(upper) * dx0
	return sum(parts) / 6
}

func sum(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	var parts []float64
	for _, x := range v {
		var i int
		for _, p := range parts {
			sum := p + x
			var err float64
			if math.Abs(x) < math.Abs(p) {
				err = x - (sum - p)
			} else {
				err = p - (sum - x)
			}
			if err != 0 {
				parts[i] = err
				i++
			}
			x = sum
		}
		parts = append(parts[:i], x)
	}
	var sum float64
	for _, x := range parts {
		sum += x
	}
	return sum
}
