package usecases

import (
	"engineering-task/domain"
	"sort"
	"strings"
)

func isCompleteMatch(make, model string, year, price int, car domain.Car) bool {
	return isMakeModelMatch(make, model, car) && car.Year == year && car.Price == price
}

func isMakeModelMatch(make, model string, car domain.Car) bool {
	return strings.HasPrefix(strings.ToLower(car.Make), strings.ToLower(make)) &&
		strings.HasPrefix(strings.ToLower(car.Model), strings.ToLower(model))
}

func computePricingStatistics(pricingStatisticsCandidates []string, pricingDataMap map[string][]pricingData) []PricingStatistic {
	var pricingStatistics []PricingStatistic
	for _, carName := range pricingStatisticsCandidates {
		data := pricingDataMap[carName]
		sort.Slice(data, func(i, j int) bool {
			return data[i].price < data[j].price
		})
		medianPrice := computeMedianPrice(data)

		pricing := PricingStatistic{
			Vehicle:      carName,
			LowestPrice:  data[0].price,
			HighestPrice: data[len(data)-1].price,
			MedianPrice:  medianPrice,
		}
		pricingStatistics = append(pricingStatistics, pricing)
	}
	return pricingStatistics
}

func computeMedianPrice(data []pricingData) int {
	totalVehicles := 0
	medianPrice := 0
	for _, pricingData := range data {
		totalVehicles += pricingData.vehicleCount
	}

	if totalVehicles%2 == 0 {
		medianPrice = computeEvenMedian(data, totalVehicles)
	} else {
		medianPrice = computeOddMedian(data, totalVehicles)
	}
	return medianPrice
}

func computeEvenMedian(data []pricingData, totalVehicles int) int {
	firstMedianVehicleIndex := totalVehicles / 2
	secondMedianVehicleIndex := firstMedianVehicleIndex + 1

	firstVehiclePrice := 0
	secondVehiclePrice := 0
	for _, pricingData := range data {
		if firstMedianVehicleIndex-pricingData.vehicleCount > 0 {
			firstMedianVehicleIndex = firstMedianVehicleIndex - pricingData.vehicleCount
		} else {
			firstVehiclePrice = pricingData.price
			break
		}
	}
	for _, pricingData := range data {
		if secondMedianVehicleIndex-pricingData.vehicleCount > 0 {
			secondMedianVehicleIndex = secondMedianVehicleIndex - pricingData.vehicleCount
		} else {
			secondVehiclePrice = pricingData.price
			break
		}
	}

	return (firstVehiclePrice + secondVehiclePrice) / 2
}

func computeOddMedian(data []pricingData, totalVehicles int) int {
	medianPrice := 0
	medianVehicleIndex := (totalVehicles + 1) / 2
	for _, pricingData := range data {
		medianVehicleIndex = medianVehicleIndex - pricingData.vehicleCount
		if medianVehicleIndex <= 0 {
			medianPrice = pricingData.price
			break
		}
	}
	return medianPrice
}

func computeSuggestions(suggestionMap map[string]Suggestion) []Suggestion {
	suggestions := []Suggestion{}
	for key := range suggestionMap {
		suggestions = append(suggestions, suggestionMap[key])
	}

	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].Price <= suggestions[j].Price
	})

	return suggestions
}
