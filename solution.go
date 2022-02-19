package square

import "math"

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

type inCustomType int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum inCustomType) float64 {

	if sidesNum == SidesTriangle {
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	}
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
