package util

import (
	"fmt"
	"github.com/andrepinto/goway/domain"
)

func ClientApiKey(path string, version string) string{
	return fmt.Sprintf("[%s-%s]", version, path)
}

func ClientApiHeaders(client string, product string, version string) string{
	return fmt.Sprintf("[%s-%s-%s]", version, client, product)
}


func ClientRouteCode(client string, product string) string{
	return fmt.Sprintf("%s%s", client, product)
}

func ProductCode(path string, version string) string{
	return fmt.Sprintf("[%s-%s]", version, path)
}

func MergeInjectData(global []*domain.InjectDataV1, method []*domain.InjectDataV1) []*domain.InjectDataV1{
	result := method

	if(len(global)==0){
		return method
	}

	for _, v := range global{
		for _, k := range method{
			if(v.Code==k.Code){
				break
			}
		}

		result = append(result, v)
	}

	return result
}

func FilterClientRoutesByAssets(cl *domain.ClientV1, asset []string, f func(*domain.RoutesV1, []string) bool) []*domain.RoutesV1 {
	routes := make([]*domain.RoutesV1, 0)
	for _, v := range cl.Routes {
		if f(v, asset) {
			routes = append(routes, v)
		}
	}

	return routes
}


func FilterProductsRoutesByAssets(cl *domain.ProductV1, asset []string, f func(*domain.RoutesV1, []string) bool) []*domain.RoutesV1 {
	routes := make([]*domain.RoutesV1, 0)
	for _, v := range cl.Routes {
		if f(v, asset) {
			routes = append(routes, v)
		}
	}

	return routes
}


func FilterArrClientRoutesByAssets(cl []*domain.ClientV1, asset []string, f func(*domain.RoutesV1, []string) bool) []*domain.ClientV1 {
	clients := make([]*domain.ClientV1, 0)
	for _, c := range cl {
		routes := make([]*domain.RoutesV1, 0)
		for _, v := range c.Routes {
			if f(v, asset) {
				routes = append(routes, v)
			}
		}

		c.Routes = routes
		clients = append(clients, c)
	}


	return clients
}

func FilterArrProductRoutesByAssets(cl []*domain.ProductV1, asset []string, f func(*domain.RoutesV1, []string) bool) []*domain.ProductV1 {
	products := make([]*domain.ProductV1, 0)
	for _, c := range cl {
		routes := make([]*domain.RoutesV1, 0)
		for _, v := range c.Routes {
			if f(v, asset) {
				routes = append(routes, v)
			}
		}

		c.Routes = routes
		products = append(products, c)
	}


	return products
}

func FilterByAsset(route *domain.RoutesV1, assets []string) bool{
	return Filter(assets, route.Asset)
}


func Filter(vs []string, value string) bool {
	for _, v := range vs {
		if v==value {
			return true
		}
	}
	return false
}

func PrepareUrl(url string) string{
	if  url[len(url)-1] != '/' {
		return fmt.Sprintf("%s/",url)
	}
	return  url
}