package main

import (
	"fmt"
)

func main() {

	cantonPopulations := make(map[string]int)
	cantonPopulations = map[string]int{
		"Zurich":                 1538848,
		"Bern":                   1039101,
		"Waadt":                  804861,
		"Aargau":                 685642,
		"St.Gallen":              510670,
		"Genf":                   504031,
		"Luzern":                 413048,
		"Tessin":                 351471,
		"Wallis":                 345394,
		"Fribourg":               321717,
		"Basel-Landschaft":       289404,
		"Thurgau":                279493,
		"Solothurn":              275177,
		"Graubuenden":            198988,
		"Basel-Stadt":            195783,
		"Neuenburg":              176467,
		"Schwyz":                 160457,
		"Zug":                    127612,
		"Schaffhausen":           82337,
		"Jura":                   73563,
		"Appenzell Ausserrhoden": 55432,
		"Nidwalden":              43076,
		"Glarus":                 40852,
		"Obwalden":               37924,
		"Uri":                    36694,
		"Appenzell Innerhoden":   16127,
	}

	fmt.Println(calcPercentages(cantonPopulations))
}

// Calculate the percentages of each canton
func calcPercentages(input map[string]int) map[string]float32 {
	// Caluculate the totals for switzerland
	total := 0
	result := make(map[string]float32)
	for _, v := range input { // Using only the values, keys are discarded (_)
		total += v
	}
	// Calculating the percentages
	for k, v := range input {
		percentage := (float32(v) / float32(total)) * 100
		result[k] = percentage
	}
	return result
}
