package forever

import "iFei/model"

func AddArticle(article *model.Article) {
	//todo : 从 userid 中查出user, 更改user信息
	db.Create(article)
}
