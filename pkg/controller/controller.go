package controller

import "github.com/efectn/fiber-boilerplate/pkg/service"

type Controller struct {
	Article *ArticleController
}

func NewController(articleService *service.ArticleService) *Controller {
	return &Controller{
		Article: &ArticleController{articleService: articleService},
	}
}