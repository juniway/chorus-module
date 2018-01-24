// Copyright 2015 The chorus Authors
// This file is part of the chorus library.
//
// The chorus library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The chorus library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the chorus library. If not, see <http://www.gnu.org/licenses/>.

package math

import (
	"fmt"
	"math/big"
	"testing"
)

type summer struct {
	numbers []*big.Int
}

func (s summer) Len() int { return len(s.numbers) }
func (s summer) Sum(i int) *big.Int {
	return s.numbers[i]
}

func TestSum(t *testing.T) {
	summer := summer{numbers: []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}}
	sum := Sum(summer)
	if sum.Cmp(big.NewInt(6)) != 0 {
		t.Errorf("got sum = %d, want 6", sum)
	}
}

func TestDist(t *testing.T) {
	var vectors = []Vector{
		{big.NewInt(1000), big.NewInt(1234)},
		{big.NewInt(500), big.NewInt(10023)},
		{big.NewInt(1034), big.NewInt(1987)},
		{big.NewInt(1034), big.NewInt(1987)},
		{big.NewInt(8983), big.NewInt(1977)},
		{big.NewInt(98382), big.NewInt(1887)},
		{big.NewInt(12398), big.NewInt(1287)},
		{big.NewInt(12398), big.NewInt(1487)},
		{big.NewInt(12398), big.NewInt(1987)},
		{big.NewInt(12398), big.NewInt(128)},
		{big.NewInt(12398), big.NewInt(1987)},
		{big.NewInt(1398), big.NewInt(187)},
		{big.NewInt(12328), big.NewInt(1927)},
		{big.NewInt(12398), big.NewInt(1987)},
		{big.NewInt(22398), big.NewInt(1287)},
		{big.NewInt(1370), big.NewInt(1981)},
		{big.NewInt(12398), big.NewInt(1957)},
		{big.NewInt(42198), big.NewInt(1987)},
	}

	VectorsBy(GasSort).Sort(vectors)
	fmt.Println(vectors)

	BP := big.NewInt(15)
	GL := big.NewInt(1000000)
	EP := big.NewInt(100)
	fmt.Println("BP", BP, "GL", GL, "EP", EP)
	GP := GasPrice(BP, GL, EP)
	fmt.Println("GP =", GP, "Wei per GU")

	S := len(vectors) / 4
	fmt.Println("L", len(vectors), "S", S)
	for i := 1; i <= S*4; i += S {
		fmt.Printf("T%d = %v\n", i, vectors[i])
	}

	g := VectorSum(GasSum).Sum(vectors)
	fmt.Printf("G = ∑g* (%v)\n", g)
}
