package task

import "net/http"

type API struct{}

// List godoc
//
// @summary List items
// @description List items
// @tags items
// @accept json
// @produce json
// @success 200 {array} DTO
// @failure 500 {object} err.Error
// @router /items [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {}

// Create godoc
//
// @summary Create item
// @description Create item
// @tags item
// @accept json
// @produce json
// @param body body Form true "Item form"
// @success 201
// @failure 400 {object} err.Error
// @failure 422 {object} err.Errors
// @failure 500 {object} err.Error
// @router /items [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

// Read godoc
//
// @summary Read item
// @description Read item
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @success 200 {object} DTO
// @failure 400 {object} err.Error
// @failure 404
// @failure 500 {object} err.Error
// @router /items/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {}

// Update godoc
//
// @summary Delete item
// @description Delete item
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @param body body Form true "Item form"
// @success 200
// @failure 400 {object} err.Error
// @failure 404
// @failure 422 {object} err.Errors
// @failure 500 {object} err.Error
// @router /items/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

// Delete godoc
//
// @summary Delete item
// @description
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @success 200
// @failure 400 {object} err.Error
// @failure 404
// @failure 500 {object} err.Error
// @router /items/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
