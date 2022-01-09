package usecases

import (
	"engineering-task/domain"
	"fmt"
	"sort"
	"strings"
)

type Suggestion struct {
	Make  string
	Model string
	Year  string
	Price string
}

type pricingData struct {
	vehicleCount int
	price        int
}

type PricingStatistic struct {
	Vehicle      string
	LowestPrice  int
	MedianPrice  int
	HighestPrice int
}

type SearchResult struct {
	TotalCount          int
	MakeModelMatchCount int
	PricingStatistics   []PricingStatistic
	Suggestions         []Suggestion
}

type SearchRepository interface {
	Search(make, model, year string, budget int) []SearchResult
}

type Logger interface {
	Log(args ...interface{})
}

type Car struct {
	Make  string
	Model string
	Year  string
	Price int
}

type SearchInteractor struct {
	CarRepository domain.CarRepository
	Logger        Logger
}

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
