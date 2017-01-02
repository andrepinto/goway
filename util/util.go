package util

import (
	"fmt"
	"github.com/andrepinto/goway/product"
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

func MergeInjectData(global []product.InjectData_v1, method []product.InjectData_v1) []product.InjectData_v1{
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

func FilterClientRoutesByAssets(cl *product.Client_v1, asset []string, f func(product.Routes_v1, []string) bool) []product.Routes_v1 {
	routes := make([]product.Routes_v1, 0)
	for _, v := range cl.Routes {
		if f(v, asset) {
			routes = append(routes, v)
		}
	}

	return routes
}


func FilterProductsRoutesByAssets(cl *product.Product_v1, asset []string, f func(product.Routes_v1, []string) bool) []product.Routes_v1 {
	routes := make([]product.Routes_v1, 0)
	for _, v := range cl.Routes {
		if f(v, asset) {
			routes = append(routes, v)
		}
	}

	return routes
}


func FilterArrClientRoutesByAssets(cl []product.Client_v1, asset []string, f func(product.Routes_v1, []string) bool) []product.Client_v1 {
	clients := make([]product.Client_v1, 0)
	for _, c := range cl {
		routes := make([]product.Routes_v1, 0)
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

func FilterArrProductRoutesByAssets(cl []product.Product_v1, asset []string, f func(product.Routes_v1, []string) bool) []product.Product_v1 {
	products := make([]product.Product_v1, 0)
	for _, c := range cl {
		routes := make([]product.Routes_v1, 0)
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

func FilterByAsset(route product.Routes_v1, assets []string) bool{
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