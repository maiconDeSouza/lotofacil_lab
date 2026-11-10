package random

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/maiconDeSouza/lotofacil_lab/internal/model"
)

func GenerateLotofacilNumbers15() []int {
	s := rand.Perm(25)

	s = s[:15]

	for i := range s {
		s[i]++
	}

	slices.Sort(s)

	return s
}

func GenerateFrequent(numbers *[]*model.NumberFrequency) (model.Numbers, error) {
	if numbers == nil || len(*numbers) != 25 {
		return nil, fmt.Errorf("Erro gera números frequentes")
	}

	n := *numbers

	range1 := n[0:5]   //5
	range2 := n[5:10]  //4
	range3 := n[10:15] //3
	range4 := n[15:20] //2
	range5 := n[20:25] //1

	var tier1 []int
	var tier2 []int
	var tier3 []int
	var tier4 []int
	var tier5 []int

	for _, v := range range1 {
		tier1 = append(tier1, v.Number)
	}

	for _, v := range range2 {
		tier2 = append(tier2, v.Number)
	}

	for _, v := range range3 {
		tier3 = append(tier3, v.Number)
	}

	for _, v := range range4 {
		tier4 = append(tier4, v.Number)
	}

	for _, v := range range5 {
		tier5 = append(tier5, v.Number)
	}

	var result []int
	r1, err := DrawNumbers(tier1, 5)
	if err != nil {
		return nil, fmt.Errorf("Erro gerar números")
	}
	r2, err := DrawNumbers(tier2, 4)
	if err != nil {
		return nil, fmt.Errorf("Erro gerar números")
	}
	r3, err := DrawNumbers(tier3, 3)
	if err != nil {
		return nil, fmt.Errorf("Erro gerar números")
	}
	r4, err := DrawNumbers(tier4, 2)
	if err != nil {
		return nil, fmt.Errorf("Erro gerar números")
	}
	r5, err := DrawNumbers(tier5, 1)
	if err != nil {
		return nil, fmt.Errorf("Erro gerar números")
	}

	result = append(result, r1...)
	result = append(result, r2...)
	result = append(result, r3...)
	result = append(result, r4...)
	result = append(result, r5...)

	if len(result) != 15 {
		return nil, fmt.Errorf("Erro gerar números")
	}

	return result, nil
}

func DrawNumbers(source []int, quantity int) ([]int, error) {
	if quantity > len(source) {
		return nil, fmt.Errorf("quantidade solicitada (%d) é maior que o número de elementos disponíveis (%d)", quantity, len(source))
	}
	if quantity <= 0 {
		return nil, fmt.Errorf("a quantidade a sortear deve ser maior que zero")
	}

	copiedNumbers := make([]int, len(source))
	copy(copiedNumbers, source)

	rand.Shuffle(len(copiedNumbers), func(i, j int) {
		copiedNumbers[i], copiedNumbers[j] = copiedNumbers[j], copiedNumbers[i]
	})

	return copiedNumbers[:quantity], nil
}
