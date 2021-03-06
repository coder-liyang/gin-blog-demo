package models

type Article struct {
	Model

	TagId         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_Url"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	//offset (pageNum-1)*pageSize
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? and deleted_on =?", id, 0).First(&article).Related(&article.Tag).Error
	return &article, err
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{})
	return true
}
