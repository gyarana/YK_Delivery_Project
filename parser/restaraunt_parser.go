package parser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"nix_education/model"
	"nix_education/model/repositories"
	"sync"
	"time"
)

func NewRestarauntsParser(urlRest string, urlItems string, logger *logrus.Logger, restaurantRepositories *repositories.RestaurantsRepository, menuRepositories *repositories.MenuRepository) *RestaurantsParser {
	return &RestaurantsParser{
		restaurantsRepositories: restaurantRepositories,
		menuRepositories:        menuRepositories,
		urlRest:                 urlRest,
		urlItems:                urlItems,
		logger:                  logger,
	}
}

type RestaurantsParser struct {
	restaurantsRepositories *repositories.RestaurantsRepository
	menuRepositories        *repositories.MenuRepository
	urlRest                 string
	urlItems                string
	logger                  *logrus.Logger
}

type RestaurantsParserI interface {
	TimeFieldUpdate()
	SupplierParser(restaurant *model.RestaurantParse)
	MenuParser(items model.ProductParse, idRest int)
	GetRestData(ctx context.Context, url string, restSup model.Suppliers) *model.Suppliers
	GetMenuData(ctx context.Context, urlItems string, prItems model.RestarauntMenu, idRest int) *model.RestarauntMenu
}

func (r RestaurantsParser) TimeFieldUpdate() {
	var restSup model.Suppliers
	for {
		ctx := context.Background()
		time.Sleep(time.Duration(1 * time.Second))
		rest := r.GetRestData(ctx, r.urlRest, restSup)
		var wg sync.WaitGroup
		for _, restaurant := range rest.Restaurants {
			wg.Add(1)
			r.SupplierParser(&restaurant)
			wg.Done()
		}
		wg.Wait()
	}
}

func (r RestaurantsParser) SupplierParser(restaurant *model.RestaurantParse) {

	var wg sync.WaitGroup
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
	ctx := context.Background()
	product := r.GetMenuData(ctx, r.urlItems, prItems, idRest)
	for _, items := range product.Menu {
		wg.Add(1)
		r.MenuParser(items, idRest)
		wg.Done()
	}
	wg.Wait()
}

func (r RestaurantsParser) MenuParser(items model.ProductParse, idRest int) {
	resultItems, err := r.menuRepositories.GetMenuByID(items.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resultItems.ID == 0 {
		r.menuRepositories.CreateMenu(&items)
	} else {
		r.menuRepositories.UpdateMenu(&items)
	}
}

func (r RestaurantsParser) GetRestData(ctx context.Context, url string, restSup model.Suppliers) *model.Suppliers {
	resp, err := http.Get(url)
	if err != nil {
		r.logger.Error("We have some problem with parsing url. Please check it!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = json.Unmarshal(body, &restSup)
	if err != nil {
		r.logger.Error("We have some problem with unmarshalling data. Please check it!")
	}
	select {
	case <-ctx.Done():
		return nil
	default:
		return &restSup
	}
}

func (r RestaurantsParser) GetMenuData(ctx context.Context, urlItems string, prItems model.RestarauntMenu, idRest int) *model.RestarauntMenu {
	url := fmt.Sprintf(urlItems, idRest)
	respItem, err := http.Get(url)
	if err != nil {
		r.logger.Error("We have some problem with parsing url. Please check it!")
	}
	body, err := ioutil.ReadAll(respItem.Body)
	defer respItem.Body.Close()
	err = json.Unmarshal(body, &prItems)
	if err != nil {
		r.logger.Error("We have some problem with unmarshalling data. Please check it!")
	}
	select {
	case <-ctx.Done():
		return nil
	default:
		return &prItems
	}

}
