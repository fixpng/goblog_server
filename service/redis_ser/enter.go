package redis_ser

const (
	articleCommentCountPrefix = "article_comment_count"
	articleLookPrefix         = "article_look"
	articleDiggPrefix         = "article_digg"
	commentDiggPrefix         = "comment_digg"
)

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}

func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}

func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}

func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
