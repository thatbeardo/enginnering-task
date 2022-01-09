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

		totalVehicles := 0
		for _, pricingData := range data {
			totalVehicles += pricingData.vehicleCount
		}
		medianVehicleIndex := totalVehicles / 2
		medianPrice := 0
		for _, pricingData := range data {
			medianVehicleIndex = medianVehicleIndex - pricingData.vehicleCount
			if medianVehicleIndex <= 0 {
				medianPrice = pricingData.price
				break
			}
		}

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
