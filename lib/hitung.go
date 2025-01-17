package lib

import "math"

type Hitung interface{
	Luas() float64
	Keliling() float64
}

type Lingkaran struct{
	Diameter float64
}

func (l Lingkaran) jariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}


type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) Keliling() float64 {
	return p.Sisi * 4
}	

