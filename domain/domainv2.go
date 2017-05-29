package domain

import (
	"net/url"
	"fmt"
)

//TEMPORARY

type ServiceV1 struct {
	Code string
	Name string
	Schema string
	Endpoint string
	Port int
	BasePath string
}

func ServiceV1ToServiceTarget(svr []*ServiceV1) (map[string]*url.URL, error){
	result := make(map[string]*url.URL)
	for _ , k := range svr{
		url, err := url.Parse(fmt.Sprintf("%s://%s:%d/%s", k.Schema, k.Endpoint, k.Port, k.BasePath))
		if err!=nil{
			return nil, err
		}
		result[k.Code]=url
	}
	return result, nil
}