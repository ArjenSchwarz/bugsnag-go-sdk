package api

import (
	"testing"
)

func TestParametersSortBy(t *testing.T) {
	params := Parameters{}
	sortBy, err := params.sortBy()
	if err != nil || sortBy != "created_at" {
		t.Error("Default value is incorrect")
	}
	params.Sort = "created_at"
	sortBy, err = params.sortBy()
	if err != nil || sortBy != "created_at" {
		t.Error("Setting sort to created_at doesn't work")
	}
	params.Sort = "updated_at"
	sortBy, err = params.sortBy()
	if err == nil || sortBy != "" {
		t.Error("It shouldn't be possible to sort by anything other than created_at")
	}
}

func TestParametersOrderDirection(t *testing.T) {
	params := Parameters{}
	order, err := params.direction()
	if err != nil || order != "desc" {
		t.Error("Default value is incorrect")
	}
	params.Direction = "desc"
	order, err = params.direction()
	if err != nil || order != "desc" {
		t.Error("Setting direction to desc doesn't work")
	}
	params.Direction = "asc"
	order, err = params.direction()
	if err != nil || order != "asc" {
		t.Error("Setting direction to asc doesn't work")
	}
	params.Direction = "updated_at"
	order, err = params.direction()
	if err == nil {
		t.Error("It shouldn't be possible to sort by anything other than asc or desc")
	}
}

func TestParametersPerPage(t *testing.T) {
	params := Parameters{}
	perPage, err := params.perPage()
	if err != nil || perPage != 30 {
		t.Error("Default value is incorrect")
	}
	params.PerPage = 10
	perPage, err = params.perPage()
	if err != nil || perPage != 10 {
		t.Error("Setting per page to 10 doesn't work")
	}
	params.PerPage = -1
	perPage, err = params.perPage()
	if err == nil || perPage != 0 {
		t.Error("It shouldn't be possible to request less than 0 items per page")
	}
}
