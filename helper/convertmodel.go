package helper

import (
	"net/url"

	"github.com/Alfeenn/article/model"
	"github.com/Alfeenn/article/model/web"
)

func ConvertModel(req model.User) web.CatResp {

	// categoryslice := GetCategoryObject(req)
	// slugname := DashString(req.Email)

	return web.CatResp{
		Id:        req.Id,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}

func GetCategoryObject(req model.User) []model.CategoryArticle {
	//get category object then slug category
	var sliceCategory []model.CategoryArticle
	categoryArticle := model.CategoryArticle{
		Name: req.Email,
		Url:  req.Password,
	}

	catName := LowerAndDash(categoryArticle.Name)

	for v := range catName {
		categoryArticle.Name = catName[v]
		categoryArticle.Slug = catName[v]
		url, _ := url.JoinPath("http://localhost:8000/api/categories", catName[v])
		categoryArticle.Url = url

		sliceCategory = append(sliceCategory, categoryArticle)
	}
	return sliceCategory
}
