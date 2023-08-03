// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HotTuiJianReq struct {
	Ids []string `json:"ids,omitempty,default=[]"`
}

type PageReq struct {
	Page int `json:"page,omitempty,default=1"`
	Size int `json:"size,omitempty,default=5"`
}

type UpdataPostDataReq struct {
	Order int `json:"order,omitempty,default=1"`
	Page  int `json:"page,omitempty,default=1"`
	Size  int `json:"size,omitempty,default=5"`
}

type BookInfoReq struct {
	Id int `path:"id""`
}

type LastChapterReq struct {
	Id    int `json:"id""`
	Limit int `json:"limit""`
}

type HotBookTuiJianReq struct {
	Order int `json:"order,omitempty,default=1"`
	Page  int `json:"page,omitempty,default=1"`
	Size  int `json:"size,omitempty,default=5"`
}

type XiangGuanTuiJianReq struct {
	Id    int `json:"id""`
	Limit int `json:"limit""`
}

type ChapterListReq struct {
	Id int `json:"id""`
}

type ChapterContextReq struct {
	Id  int `json:"id"`
	Cid int `json:"cid"`
}

type AuthorAllReq struct {
	Name string `json:"name"`
}

type SortBookListReq struct {
	SortId int `json:"sortId"`
	Status int `json:"status"`
	Page   int `json:"page,omitempty,default=1"`
	Size   int `json:"size,omitempty,default=5"`
}

type SqlSerachBookReq struct {
	Key  string `json:"key"`
	Page int    `json:"page,omitempty,default=1"`
	Size int    `json:"size,omitempty,default=5"`
}

type UserRegReq struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Email    string `json:"email"`
}

type UserLoginReq struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

type UserBookCaseReq struct {
	Uid int `json:"uid"`
}

type AddBookCaseReq struct {
	Articleid   int    `json:"articleid"`
	Articlename string `json:"articlename"`
	Userid      int    `json:"userid"`
	Username    string `json:"username"`
	Chapterid   int    `json:"chapterid"`
	Chaptername string `json:"chaptername"`
}

type QueryBookCaseReq struct {
	Uid int `json:"uid"`
}

type DelBookCaseReq struct {
	CaseId int `json:"caseId"`
}
