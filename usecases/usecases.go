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
	suggestionMap := make(map[string]Suggestion)

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
		if float32(price) >= 0.9*float32(car.Price) && float32(price) <= 1.1*float32(car.Price) {
			si.Logger.Log(fmt.Sprintf("Matched %s in price window %f and %f", car.Make, 0.9*float32(car.Price), 1.1*float32(car.Price)))
			if _, present := suggestionMap[car.Make]; !present {
				suggestionMap[car.Make] = Suggestion{
					Make:  car.Make,
					Model: car.Model,
					Year:  car.Year,
					Price: car.Price,
				}
			}
		}
	}

	pricingStatistics := computePricingStatistics(pricingStatisticsCandidates, pricingDataMap)
	suggestions := computeSuggestions(suggestionMap)

	return SearchResult{
		TotalCount:          totalCars,
		MakeModelMatchCount: makeModelMatch,
		PricingStatistics:   pricingStatistics,
		Suggestions:         suggestions,
	}
}
