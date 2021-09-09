package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nix_education/model"
	"nix_education/model/repositories"
	"sync"
	"time"
)

func NewRestarauntsParser(urlRest string, urlItems string, restaurantRepositories *repositories.RestaurantsRepository, menuRepositories *repositories.MenuRepository) *RestaurantsParser {
	return &RestaurantsParser{
		restaurantsRepositories: restaurantRepositories,
		menuRepositories:        menuRepositories,
		urlRest:                 urlRest,
		urlItems:                urlItems,
	}
}

type RestaurantsParser struct {
	restaurantsRepositories *repositories.RestaurantsRepository
	menuRepositories        *repositories.MenuRepository
	urlRest                 string
	urlItems                string
}

type RestaurantsParserI interface {
	TimeFieldUpdate()
	SupplierParser(rest *model.Supliers)
	MenuParser(items model.Product, idRest int)
}

func (r RestaurantsParser) TimeFieldUpdate() {
	var restSup model.Supliers

	for {
		time.Sleep(time.Duration(60 * time.Second))
		rest := GetRestData(r.urlRest, restSup)
		var wg sync.WaitGroup
		for _, restaurant := range rest.Restaurants {
			wg.Add(1)
			go r.SupplierParser(&restaurant)
			wg.Done()
		}
		wg.Wait()
	}

}

func (r RestaurantsParser) SupplierParser(restaurant *model.Restaurant) {

	//var restSup model.Supliers
	var wg sync.WaitGroup
	//rest:= GetRestData(r.urlRest,restSup)
	//for _, restaurant := range rest.Restaurants {
	resultRest, err := r.restaurantsRepositories.GetSuppliersByID(restaurant.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resultRest.Id == 0 {
		r.restaurantsRepositories.CreateSuppliers(restaurant)
	} else {
		r.restaurantsRepositories.UpdateSuppliers(restaurant)
	}
	idRest := restaurant.Id
	var prItems model.RestarauntMenu
	product := GetMenuData(r.urlItems, prItems, idRest)
	for _, items := range product.Menu {
		wg.Add(1)
		go r.MenuParser(items, idRest)
		wg.Done()
	}
	wg.Wait()
	//}
}

func (r RestaurantsParser) MenuParser(items model.Product, idRest int) {
	resultItems, err := r.menuRepositories.GetMenuByRestID(items.ID, idRest)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resultItems.ID == 0 {
		r.menuRepositories.CreateMenu(idRest, &items)
	} else {
		r.menuRepositories.UpdateMenu(idRest, &items)
	}
}

func GetRestData(url string, restSup model.Supliers) *model.Supliers {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("We have some problem with parsing url. Please check it!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = json.Unmarshal(body, &restSup)
	if err != nil {
		fmt.Println("We have some problem with unmarshalling data. Please check it!")
	}
	return &restSup
}

func GetMenuData(urlItems string, prItems model.RestarauntMenu, idRest int) *model.RestarauntMenu {
	url := fmt.Sprintf(urlItems, idRest)
	respItem, err := http.Get(url)
	if err != nil {
		fmt.Println("We have some problem with parsing url. Please check it!")
	}
	body, err := ioutil.ReadAll(respItem.Body)
	defer respItem.Body.Close()
	err = json.Unmarshal(body, &prItems)
	if err != nil {
		fmt.Println("We have some problem with unmarshalling data. Please check it! тут проблема")
	}
	return &prItems
}
