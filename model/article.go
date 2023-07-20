package model

type ArticleModel struct {
	Id      int64  `json:"id"  msg:"ID不符合规范"`
	Title   string `json:"title" binding:"required" msg:"标题不能为空"`
	Content string `json:"content" binding:"required" msg:"内容不能为空"`
}

type UserModel struct {
	Name string `json:"name" binding:"sign,max=50,min=4" msg:"用户名不符合规范"`
	Age  int    `json:"age" binding:"gt=0,lt=100" msg:"用户年龄不符合规范"`
}
