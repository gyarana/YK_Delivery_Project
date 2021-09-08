
package parser

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"nix_education/model"
	"nix_education/model/repositories"
)

func NewRestarauntsParser(urlRest string,urlItems string,restaurantRepositories *repositories.RestaurantsRepository, menuRepositories *repositories.MenuRepository) *RestaurantsParser {
	return &RestaurantsParser{
		restaurantsRepositories: restaurantRepositories,
		menuRepositories: menuRepositories,
		urlRest:            urlRest,
		urlItems: urlItems,
	}
}

type RestaurantsParser struct {
	restaurantsRepositories *repositories.RestaurantsRepository
	menuRepositories * repositories.MenuRepository
	urlRest            string
	urlItems            string
}

type RestaurantsParserI interface {
	RestarauntsAndMenuParser()
	MenuParser(items model.Product,idRest int)
}

func (r RestaurantsParser) RestarauntsAndMenuParser() {

	var restSup model.Supliers
	rest:=GetAndUnmarshalRestData(r.urlRest,restSup)
	for _, restaurant := range rest.Restaurants {
		resultRest, err := r.restaurantsRepositories.GetSuppliersByID(restaurant.Id)
		if err != nil {
			fmt.Println(err.Error())
		}
		if restaurant.Id == resultRest.Id {
			r.restaurantsRepositories.UpdateSuppliers(&restaurant)
		} else {
			r.restaurantsRepositories.CreateSuppliers(&restaurant)
		}
		idRest := restaurant.Id
		var prItems model.RestarauntMenu
		product:= GetAndUnmarshalMenuData(r.urlItems,prItems,idRest)
		for _, items := range product.Menu {
			go r.MenuParser(items, idRest)
		}
	}
}

func (r RestaurantsParser) MenuParser(items model.Product,idRest int) {
	resultItems, err := r.menuRepositories.GetMenuByRestID(items.ID, idRest)
	if err != nil {
		fmt.Println(err.Error())
	}
	if items.ID == resultItems.ID {
		r.menuRepositories.UpdateMenu(idRest, &items)
	} else {
		r.menuRepositories.CreateMenu(idRest,&items)
	}
}

func GetAndUnmarshalRestData(url string,restSup model.Supliers) *model.Supliers{
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

func GetAndUnmarshalMenuData(urlItems string,prItems model.RestarauntMenu,idRest int) *model.RestarauntMenu{
	url := fmt.Sprintf(urlItems, idRest)
	respItem, err := http.Get(url)
	if err != nil {
		fmt.Println("We have some problem with parsing url. Please check it!")
	}
	body, err := ioutil.ReadAll(respItem.Body)
	defer respItem.Body.Close()
	err = json.Unmarshal(body, &prItems)
	if err != nil {
		fmt.Println("We have some problem with unmarshalling data. Please check it!")
	}
	return &prItems
}

