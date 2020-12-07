package article_service

import "gin-blog/models"

type Article struct {
	ID int
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID), nil
}

func (a *Article) Get() (*models.Article, error) {
	return models.GetArticle(a.ID)
}
