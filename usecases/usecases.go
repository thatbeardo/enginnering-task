package usecases

import (
	"fmt"
)

func (si SearchInteractor) Search(manufacturer, model string, year, price int) SearchResult {
	cars := si.CarRepository.GetAllCars()
	si.Logger.Log(fmt.Sprintf("Scanning through %d records", len(cars)))

	totalCars := 0
	makeModelMatch := 0
	pricingStatisticsCandidates := []string{}
	pricingDataMap := make(map[string][]pricingData)

	for _, car := range cars {

		var carKey = car.Make + car.Model
		if data, ok := pricingDataMap[car.Make+car.Model]; ok {
			data = append(data, pricingData{vehicleCount: car.VehicleCount, price: car.Price})
			pricingDataMap[carKey] = data
		} else {
			pricingDataMap[carKey] = []pricingData{{vehicleCount: car.VehicleCount, price: car.Price}}
		}

		if isCompleteMatch(manufacturer, model, year, price, car) {
			totalCars += car.VehicleCount
		}
		if isMakeModelMatch(manufacturer, model, car) {
			makeModelMatch += car.VehicleCount
		}
		if price == car.Price {
			pricingStatisticsCandidates = append(pricingStatisticsCandidates, car.Make+car.Model)
		}
	}

	pricingStatistics := computePricingStatistics(pricingStatisticsCandidates, pricingDataMap)

	return SearchResult{
		TotalCount:          totalCars,
		MakeModelMatchCount: makeModelMatch,
		PricingStatistics:   pricingStatistics,
	}
}
