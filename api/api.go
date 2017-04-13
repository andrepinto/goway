package api

import (
	"github.com/andrepinto/goway/action"
	"github.com/joelbraga/aztek"
)

type ApiOptions struct {
	Repository  aztek.Repository
	ActionEvent *action.ActionEvent
}

type ApiResource struct {
	Repository  aztek.Repository
	ActionEvent *action.ActionEvent
}

func NewApiResource(options *ApiOptions) *ApiResource {
	repo := options.Repository
	event := options.ActionEvent

	if repo == nil {
		panic("Repository is required")
	}

	return &ApiResource{
		repo,
		event,
	}
}
